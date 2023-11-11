package main

import (
	"fmt"
	"os"
	"sennett-lau/rpsg/utils"
)

func main() {
	var err error

	// ====================================================================================================

	ignoreList := utils.GetDefaultIgnoreList()

	// the list of extensions to ignore from .rpsgignore
	var dotRpsgIgnoreList []string

	// get the list of extensions to ignore from .rpsgignore
	dotRpsgIgnoreList, err = utils.GetDotRpsgIgnoreList(".rpsgignore")

	if err != nil {
		fmt.Println("Error getting .rpsgignore list:", err)
		return
	}

	combinedList := append(ignoreList, dotRpsgIgnoreList...)

	// ====================================================================================================

	flagOutputs, err := utils.ArgsFlagCheck(os.Args[1:])

	if err != nil {
		fmt.Println(err)
		return
	}

	combinedList = append(combinedList, flagOutputs.IgnoreListExtends...)

	if flagOutputs.ShowIgnoreList {
		fmt.Println("Ignore list:")
		for _, item := range ignoreList {
			fmt.Println(item)
		}
		fmt.Println()
	}

	if flagOutputs.ShowHelp {
		fmt.Println("Usage: rpsg [flags]")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("  -h, --help                        show help")
		fmt.Println("  -s, --show-ignore-list            show the list of extensions to ignore")
		fmt.Println("  -i, --ignore-list-extends         add extensions to ignore")
		fmt.Println("  -d, --max-depth                   set the maximum depth to traverse")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  rpsg -s")
		return
	}

	if !flagOutputs.RunRpsg {
		return
	}

	dir, err := utils.GetProjectStructure(".", combinedList)
	if err != nil {
		fmt.Println("Error getting project structure:", err)
		return
	}

	// print the structure of the directory
	// maxDepth + 1 because the root directory is not counted as a level
	structure := utils.ConstructStructure(dir, 0, false, 0, flagOutputs.MaxDepth + 1, []bool{})

	err = utils.SaveStructureToFile(structure, "structure.md")

	if err != nil {
		fmt.Println("Error saving structure to file:", err)
		return
	}

	fmt.Println("Successfully saved structure to file")
}
