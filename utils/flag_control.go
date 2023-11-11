package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type flag struct {
	Name        string
	ShortName   string
	Description string
	IsValid     func(string) bool
}

type flagOutput struct {
	// run the rpsg program to generate the structure.md file
	RunRpsg bool

	// the list of extensions to ignore with the --extend-ignore-list / -e flag
	IgnoreListExtends []string

	// set the max depth (default is 6) of the structure with the --max-depth / -d flag
	MaxDepth int

	// show the ignore list and exit without running the rpsg program with the --show-ignore-list / -s flag
	ShowIgnoreList bool

	// show the help and exit without running the rpsg program with the --help / -h flag
	ShowHelp bool
}

var flags = []flag{
	{Name: "--extend-ignore-list", ShortName: "-e", Description: "Extend the default ignore list with the given list of files and directories", IsValid: ArgIsValidExtendIgnoreList},
	{Name: "--max-depth", ShortName: "-d", Description: "Set the maximum depth of the search", IsValid: ArgIsValidMaxDepth},
	{Name: "--show-ignore-list", ShortName: "-s", Description: "Show the default ignore list", IsValid: (func(string) bool)(nil)},
	{Name: "--help", ShortName: "-h", Description: "Show the help", IsValid: (func(string) bool)(nil)},
}

var flagOutputs = flagOutput{
	RunRpsg:           true,
	IgnoreListExtends: []string{},
	MaxDepth:          6,
	ShowIgnoreList:    false,
	ShowHelp:          false,
}

func ArgsFlagCheck(args []string) (flagOutput, error) {
	skipArg := false
	var err error

	for index, arg := range args {
		// skip the arg after the short form of the flag
		if skipArg {
			skipArg = false
			continue
		}

		if !strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			return flagOutputs, fmt.Errorf("invalid argument: %s", arg)
		}

		isShortForm := strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--")

		if arg == "--show-ignore-list" || arg == "-s" {
			flagOutputs.ShowIgnoreList = true
			flagOutputs.RunRpsg = false

			return flagOutputs, nil
		}

		if arg == "--help" || arg == "-h" {
			flagOutputs.ShowHelp = true
			flagOutputs.RunRpsg = false

			return flagOutputs, nil
		}

		var targetFlag *flag

		for _, flag := range flags {
			if isShortForm && flag.ShortName == arg {
				targetFlag = &flag
				skipArg = true
				break
			} else if !isShortForm && flag.Name == strings.Split(arg, "=")[0] {
				targetFlag = &flag
				break
			}
		}

		if targetFlag == nil {
			return flagOutputs, fmt.Errorf("invalid flag: %s", arg)
		}

		var displayArg string

		if isShortForm {
			displayArg = arg
		} else {
			displayArg = strings.Split(arg, "=")[0]
		}

		if len(args) == index+1 && isShortForm {
			return flagOutputs, fmt.Errorf("missing argument after %s", displayArg)
		}

		// combine the flag with its argument by adding a "=" between them
		var flagWithArg string

		if isShortForm {
			flagWithArg = arg + "=" + args[index+1]
		} else {
			flagWithArg = arg
		}

		if !targetFlag.IsValid(flagWithArg) {
			return flagOutputs, fmt.Errorf("invalid argument after %s", displayArg)
		}

		// flag update
		switch targetFlag.Name {
		case "--extend-ignore-list":
			flagOutputs.IgnoreListExtends = strings.Split(strings.Split(flagWithArg, "=")[1], ",")
		case "--max-depth":
			flagOutputs.MaxDepth, err = strconv.Atoi(strings.Split(flagWithArg, "=")[1])

			if err != nil {
				return flagOutputs, err
			}
		}
	}

	return flagOutputs, nil
}
