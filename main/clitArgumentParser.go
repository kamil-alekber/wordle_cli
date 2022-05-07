package main

import (
	"fmt"
	"os"
)

func isAllowedArgument(arg string) bool {
	appArgs := []string{"-h", "--help", "-v", "--version", "-f", "--file", "-o", "--output"}

	for _, v := range appArgs {
		if v == arg {
			return true
		}
	}
	return false
}

func cliArgumentParser() map[string]string {
	initialArgs := os.Args[1:]
	fmt.Printf("argument list: %v\n", initialArgs)

	parsedArguments := make(map[string]string)

	for i, arg := range initialArgs {
		// check if argument is value
		if arg[0] != '-' {
			continue
		}

		if !isAllowedArgument(arg) {
			fmt.Println("Argument " + arg + " is not allowed")
			os.Exit(1)
		}

		//build map of arguments
		if i == len(initialArgs)-1 {
			parsedArguments[arg] = ""
			break
		} else if initialArgs[i+1][0] != '-' {
			parsedArguments[arg] = initialArgs[i+1]
		} else {
			parsedArguments[arg] = ""
		}

	}

	return parsedArguments

}
