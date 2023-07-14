package tests

import (
	"testing"
	"sennett-lau/rpsg/utils"
)

func TestContains(t *testing.T) {
	ignoreList := []string{"*.txt", "folder1", "file2.txt"}

	// Test case 1: name matches an ignore entry that is a file type
	if !utils.Contains(ignoreList, "file1.txt") {
		t.Errorf("contains returned false for file1.txt, expected true")
	}

	// Test case 2: name matches an ignore entry that is a directory name
	if !utils.Contains(ignoreList, "folder1") {
		t.Errorf("contains returned false for folder1, expected true")
	}

	// Test case 3: name matches an ignore entry that is a file name
	if !utils.Contains(ignoreList, "file2.txt") {
		t.Errorf("contains returned false for file2.txt, expected true")
	}

	// Test case 4: name does not match any ignore entry
	if utils.Contains(ignoreList, "file3.js") {
		t.Errorf("contains returned true for file3.js, expected false")
	}
}