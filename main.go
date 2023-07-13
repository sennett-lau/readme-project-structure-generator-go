package main

import (
	"fmt"
	"sennett-lau/rpsg/utils"
)

func main() {
	// get the structure of the current directory

	ignoreList := utils.GetDefaultIgnoreList()

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
