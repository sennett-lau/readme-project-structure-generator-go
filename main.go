package main

import (
	"fmt"
	"sennett-lau/rpsg/utils"
)

func main() {
	// get the structure of the current directory
	dir, err := utils.GetProjectStructure(".")
	if err != nil {
		fmt.Println("Error getting project structure:", err)
		return
	}

	// print the structure of the directory
	structure := utils.ConstructStructure(dir, 0, false)

	fmt.Println(structure)
}