package utils

import (
	"regexp"
	"strings"
)

func ArgIsValidExtendIgnoreList(arg string) bool {
	listStrings := strings.Split(arg, "=")

	if len(listStrings) != 2 {
		return false
	}

	if listStrings[0] != "--extend-ignore-list" && listStrings[0] != "-e" {
		return false
	}

	param := listStrings[1]
	rePattern := regexp.MustCompile(`^([\w]+.[\w]+|[\w]+)/?(,([\w]+.[\w]+|[\w]+)/?)*$`)

	return rePattern.MatchString(param)
}

func ArgIsValidMaxDepth(arg string) bool {
	listStrings := strings.Split(arg, "=")

	if len(listStrings) != 2 {
		return false
	}

	if listStrings[0] != "--max-depth" && listStrings[0] != "-d" {
		return false
	}

	param := listStrings[1]
	rePattern := regexp.MustCompile(`^[1-9]$|^10$`)

	return rePattern.MatchString(param)
}
