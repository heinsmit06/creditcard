package main

import (
	"creditcard/features"
	"flag"
	"fmt"
	"os"
)

func main() {
	// defining new 4 subcommands
	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	validateStdin := validateCmd.Bool("stdin", false, "flag to pass number from stdin")

	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generatePick := generateCmd.Bool("pick", false, "flag to randomly pick a single entry from possible variants")

	//informationCmd := flag.NewFlagSet("information", flag.ExitOnError)
	//informationStdin := informationCmd.Bool("stdin", false, "flag to pass number from stdin")
	//informationBrands := informationCmd.String("brands", "brands.txt", "flag to store a .txt file with brands' names")
	//informationIssuers := informationCmd.String("issuers", "issuers.txt", "flag to store a .txt file with issuers' names")
	//
	//issueCmd := flag.NewFlagSet("information", flag.ExitOnError)
	//issueBrands := issueCmd.String("brands", "brands.txt", "flag to store a .txt file with brands' names")
	//issueIssuers := issueCmd.String("issuers", "issuers.txt", "flag to store a .txt file with issuers' names")
	//issueBrand := issueCmd.String("brand", "", "flag to choose a specific brand")
	//issueIssuer := issueCmd.String("issuer", "", "flag to choose a specific issuer")

	// to prevent the "out of range" error when accessing os.Args
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Expected 'validate' or 'generate' or 'information' or 'issue' subcommands")
		os.Exit(1)
	}

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
	//if (feature != "validate") && (feature != "generate") && (feature != "information") && (feature != "issue") {
	//	fmt.Println("The argument for the feature is not in the correct place, it should be right after the path to the bin; or all letters of the feature must be lowercase.")
	//	os.Exit(1)
	//}

	// after we are sure that the argument is in the correct place and there is only one argument
	switch feature {
	case "validate":
		validateCmd.Parse(os.Args[2:])
		features.Validate(*validateStdin, sliceOfArgs[1:]...)
	case "generate":
		generateCmd.Parse(os.Args[2:])
		features.Generate(*generatePick, sliceOfArgs[1:]...)
		// case "information":
		//	features.Information
		// case "issue":
		//	features.Issue
	}
}
