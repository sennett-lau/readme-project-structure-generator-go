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

	param := listStrings[1]
	rePattern := regexp.MustCompile(`^([\w]+.[\w]+|[\w]+)/?(,([\w]+.[\w]+|[\w]+)/?)*$`)

	return rePattern.MatchString(param)
}
