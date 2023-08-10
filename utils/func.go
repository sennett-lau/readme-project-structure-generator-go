package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sennett-lau/rpsg/types"
	"strings"
)

func GetProjectStructure(path string, ignoreList []string) (types.Directory, error) {
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
			// check if the file or directory should be ignored
			if Contains(ignoreList, file.Name()) {
				continue
			}

			// recursively get the structure of the child file or directory
			childPath := filepath.Join(path, file.Name())
			childDir, err := GetProjectStructure(childPath, ignoreList)
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

func ConstructStructure(dir types.Directory, indent int, isLast bool, depth int, maxDepth int, isExtendDepth []bool) string {
	var output string
	numOfFiles := len(dir.Files)
	numOfSubDirectories := len(dir.SubDirectories)
	
	if depth == maxDepth {
		for i := 0; i < indent; i++ {
			if i == indent-1 {
				output += "└── ...\n"
			} else if isExtendDepth[i] {
				output += "│   "
			} else {
				output += "    "
			}
		}
	} else {
		// print the name of the directory with the appropriate indentation
		for i := 0; i < indent; i++ {
			if i == indent-1 {
				if isLast {
					output += "└── "
				} else {
					output += "├── "
				}
			} else if isExtendDepth[i] {
				output += "│   "
			} else {
				output += "    "
			}
		}
		output += dir.Name + "\n"
	

		// recursively print the structure of each child directory
		dirIndex := 0
		if depth < maxDepth + 1 {
			for _, child := range dir.SubDirectories {
				if dirIndex == numOfSubDirectories-1 && numOfFiles == 0 {
					isExtendDepth = append(isExtendDepth, false)
				} else {
					isExtendDepth = append(isExtendDepth, true)
				}
				output += ConstructStructure(child, indent+1, dirIndex == numOfSubDirectories-1 && numOfFiles == 0, depth + 1, maxDepth, isExtendDepth)
				dirIndex++
			}
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
				} else if isExtendDepth[i] {
					output += "│   "
				} else {
					output += "    "
				}
			}
			fileIndex++
			output += file + "\n"
		}
	}

	return output
}

func SaveStructureToFile(structure string, path string) error {
	structure = "```\n" + structure + "```"

	// create the file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// write the structure to the file
	_, err = file.WriteString(structure)
	if err != nil {
		return err
	}

	return nil
}

func Contains(ignoreList []string, name string) bool {
	for _, ignore := range ignoreList {
		if strings.HasPrefix(ignore, "*.") {
			// if the ignore entry is a file type, check if the file extension matches
			if filepath.Ext(name) == ignore[1:] {
				return true
			}
		} else if ignore == name {
			// if the ignore entry is a file or directory name, check if the name matches
			return true
		}
	}
	return false
}

func GetDotRpsgIgnoreList(file string) ([]string, error) {
	// check if the .rpsgignore file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return []string{}, nil
	}

	// read the .rpsgignore file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// split the file content by newlines
	lines := strings.Split(string(content), "\n")

	// remove empty lines and comments
	var ignoreList []string
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "#") {
			ignoreList = append(ignoreList, line)
		}
	}

	return ignoreList, nil
}
