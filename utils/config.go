package utils

import ()

func GetDefaultIgnoreList() []string {
	return []string{
		".git",
		".gitignore",
		".idea",
		".vscode",
		"node_modules",
		"vendor",
		"*.exe",
		"*.dll",
		"*.so",
		"*.dylib",
		"*.min.js",
		"*.min.css",
		"*.min.html",
		"*.min.json",
		"*.min.xml",
		"venv",
		"__pycache__",
		"*.pyc",
	}
}
