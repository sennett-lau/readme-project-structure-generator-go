package main

import (
	"fmt"
	"os"
	"sennett-lau/rpsg/utils"
)

func main() {
	// get the structure of the current directory

	showIgnoreList := false
	runRpsg := true

	for _, arg := range os.Args {
		if arg == "--show-ignore-list" {
			showIgnoreList = true
			runRpsg = false
			break
		}
	}

	ignoreList := utils.GetDefaultIgnoreList()

	if showIgnoreList {
		fmt.Println("Ignore list:")
		for _, item := range ignoreList {
			fmt.Println(item)
		}
		fmt.Println()
	}

	if !runRpsg {
		return
	}

	dir, err := utils.GetProjectStructure(".", ignoreList)
	if err != nil {
		fmt.Println("Error getting project structure:", err)
		return
	}

	// print the structure of the directory
	structure := utils.ConstructStructure(dir, 0, false)

	err = utils.SaveStructureToFile(structure, "structure.md")

	if err != nil {
		fmt.Println("Error saving structure to file:", err)
		return
	}

	fmt.Println("Successfully saved structure to file")
}
