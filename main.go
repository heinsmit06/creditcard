package main

import (
	"creditcard/features"
	"fmt"
	"os"
)

func main() {
	feature := os.Args[1]
	sliceOfArgs := os.Args[1:]
	// to check whether there is only one feature among the args
	featuresCounter := 0
	for _, args := range sliceOfArgs {
		if args == "validate" || args == "generate" || args == "information" || args == "issue" {
			featuresCounter++
		}
	}
	if featuresCounter > 1 {
		fmt.Println("Please enter only one feature out of 4")
		os.Exit(1)
	}

	// to check if an argument for a feature is the second argument
	if (feature != "validate") && (feature != "generate") && (feature != "information") && (feature != "issue") {
		fmt.Println("The argument for the feature is not in the correct place, it should be right after the path to the bin; or all letters of the feature must be lowercase.")
		os.Exit(1)
	}

	// after we are sure that the argument is in the correct place and there is only one argument
	switch feature {
	case "validate":
		features.Validate(sliceOfArgs[1:]...)
		// case "generate":
		//	features.Generate
		// case "information":
		//	features.Information
		// case "issue":
		//	features.Issue
	}
}
