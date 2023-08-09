package main

import (
	"fmt"
	"os"
	"sennett-lau/rpsg/utils"
	"strings"
	"strconv"
)

func main() {
	// ====================================================================================================
	// The following flags can be changed to modify the behavior of the program
	// ====================================================================================================
	// run the rpsg program to generate the structure.md file
	runRpsg := true

	// show the ignore list and exit without running the rpsg program with the --show-ignore-list flag
	showIgnoreList := false

	// set the max depth (default is 6) of the structure with the --max-depth flag
	maxDepth := 6

	// the list of extensions to ignore with the --extend-ignore-list flag
	var ignoreListExtends []string

	// the list of extensions to ignore from .rpsgignore
	var dotRpsgIgnoreList []string

	// ====================================================================================================

	var err error

	// get the list of extensions to ignore from .rpsgignore
	dotRpsgIgnoreList, err = utils.GetDotRpsgIgnoreList(".rpsgignore")

	if err != nil {
		fmt.Println("Error getting .rpsgignore list:", err)
		return
	}

	for _, arg := range os.Args {
		if arg == "--show-ignore-list" {
			showIgnoreList = true
			runRpsg = false
			break
		} else if strings.HasPrefix(arg, "--extend-ignore-list=") {
			if utils.ArgIsValidExtendIgnoreList(arg) == false {
				fmt.Println("Format Error")
				return
			}

			ignoreListExtends = strings.Split(strings.Split(arg, "=")[1], ",")
		} else if strings.HasPrefix(arg, "--max-depth=") {
			if utils.ArgIsValidMaxDepth(arg) == false {
				fmt.Println("Format Error")
				return
			}

			maxDepth, err = strconv.Atoi(strings.Split(arg, "=")[1])
		}
	}

	ignoreList := utils.GetDefaultIgnoreList()

	combinedList := append(ignoreList, dotRpsgIgnoreList...)

	combinedList = append(combinedList, ignoreListExtends...)

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

	dir, err := utils.GetProjectStructure(".", combinedList)
	if err != nil {
		fmt.Println("Error getting project structure:", err)
		return
	}

	// print the structure of the directory
	structure := utils.ConstructStructure(dir, 0, false, 0, maxDepth, []bool{})

	err = utils.SaveStructureToFile(structure, "structure.md")

	if err != nil {
		fmt.Println("Error saving structure to file:", err)
		return
	}

	fmt.Println("Successfully saved structure to file")
}
