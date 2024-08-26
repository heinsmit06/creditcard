package features

import (
	"fmt"
	"os"
)

func checkBrandAndIssuerInput(brand, issuer string) {
	if brand == "" || issuer == "" {
		fmt.Fprintln(os.Stderr, "Please enter file name/s")
		os.Exit(1)
	} else if brand != "brands.txt" || issuer != "issuers.txt" {
		fmt.Fprintln(os.Stderr, "Incorrect file name/s; must be 'brands.txt' and 'issuers.txt'")
		os.Exit(1)
	}
}

func checkArgsLength(cardNumbers []string) {
	if len(cardNumbers) == 0 {
		fmt.Fprintln(os.Stderr, "There are no args, enter a card number")
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e, ": Incorrect path")
		os.Exit(1)
	}
}

func checkEmptyInput(brand, issuer string) {
	if brand == "" || issuer == "" {
		fmt.Fprintln(os.Stderr, "Please enter brand and/or issuer names")
		os.Exit(1)
	}
}
