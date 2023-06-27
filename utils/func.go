package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sennett-lau/rpsg/types"
)

func GetProjectStructure(path string) (types.Directory, error) {
	// get information about the file or directory at the specified path
	info, err := os.Stat(path)
	if err != nil {
		return types.Directory{}, err
	}

	// create a new directory with the name of the file or directory
	dir := types.Directory{Name: info.Name()}

	// if the path is a directory, recursively get the structure of its children
	if info.IsDir() {
		// get a list of all files and directories in the directory
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return types.Directory{}, err
		}

		// loop through each file or directory in the directory
		for _, file := range files {
			// recursively get the structure of the child file or directory
			childPath := filepath.Join(path, file.Name())
			childDir, err := GetProjectStructure(childPath)
			if err != nil {
				return types.Directory{}, err
			}

			// add the child directory to the parent directory's subdirectories
			if file.IsDir() {
				dir.SubDirectories = append(dir.SubDirectories, childDir)
			} else {
				dir.Files = append(dir.Files, childDir.Name)
			}
		}
	}

	return dir, nil
}

func ConstructStructure(dir types.Directory, indent int, isLast bool) string {
	var output string
	numOfFiles := len(dir.Files)
	numOfSubDirectories := len(dir.SubDirectories)

	// print the name of the directory with the appropriate indentation
	for i := 0; i < indent; i++ {
		if i == indent-1 {
			if isLast {
				output += "└── "
			} else {
				output += "├── "
			}
		} else {
			output += "│   "
		}
	}
	output += dir.Name + "\n"

	// recursively print the structure of each child directory
	dirIndex := 0
	for _, child := range dir.SubDirectories {
		output += ConstructStructure(child, indent+1, dirIndex == numOfSubDirectories-1 && numOfFiles == 0)
		dirIndex++
	}

	// print the files in the directory with the appropriate indentation
	fileIndex := 0
	for _, file := range dir.Files {
		for i := 0; i < indent+1; i++ {
			if i == indent {
				if fileIndex == len(dir.Files)-1 {
					output += "└── "
				} else {
					output += "├── "
				}
			} else {
				if isLast && i == indent-1 {
					output += "    "
				} else {
					output += "│   "
				}
			}
		}
		fileIndex++
		output += file + "\n"
	}

	return output
}
