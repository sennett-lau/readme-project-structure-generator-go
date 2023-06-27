package types

import ()

type Directory struct {
	Name           string
	Files          []string
	SubDirectories []Directory
}
