package features

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Issue(issueBrands, issueIssuers, issueBrand, issueIssuer string, tail []string) {
	fmt.Println(" brands file name:", issueBrands)
	fmt.Println(" issuers file name:", issueIssuers)
	fmt.Println(" brand:", issueBrand)
	fmt.Println(" issuer:", issueIssuer)
	fmt.Println(" tail:", tail)
	fmt.Println()

	checkBrandAndIssuerInput(issueBrands, issueIssuers)
	if len(tail) != 0 {
		fmt.Fprintln(os.Stderr, "There should not be any args when using the subcommand 'issue'")
		os.Exit(1)
	}
	checkEmptyInput(issueBrand, issueIssuer)

	fileContentBrands, err1 := os.ReadFile("./text/" + issueBrands)
	check(err1)
	fileContentIssuers, err2 := os.ReadFile("./text/" + issueIssuers)
	check(err2)

	cardBrands := strings.Split(string(fileContentBrands), "\n")
	cardIssuers := strings.Split(string(fileContentIssuers), "\n")
	issuerNumber := findCardIssuerNumber(cardIssuers, issueIssuer) // firstly I need to pass the issuer number to check for compatability
	brandNumber := findCardBrandNumber(cardBrands, issueBrand, issuerNumber)

	fmt.Println(brandNumber)
	fmt.Println(issuerNumber)

	if brandNumber != issuerNumber[:len(brandNumber)] {
		fmt.Fprintln(os.Stderr, "Number prefixes do not match")
		os.Exit(1)
	}

	if strings.ToUpper(issueBrand) == "AMEX" {
		fillerRndNum := rand.Intn(int(math.Pow10(13 - len(issuerNumber))))
		cardStartingDigits := issuerNumber + strconv.Itoa(fillerRndNum) + "**"
		generateNums(true, cardStartingDigits)
	} else {
		fillerRndNum := rand.Intn(int(math.Pow10(14 - len(issuerNumber))))
		cardStartingDigits := issuerNumber + strconv.Itoa(fillerRndNum) + "**"
		generateNums(true, cardStartingDigits)
	}
}

func findCardBrandNumber(cardBrands []string, issueBrand, issuerNumber string) string {
	for _, v := range cardBrands {
		if strings.Count(v, ":") > 1 || strings.Count(v, ":") == 0 {
			fmt.Fprintln(os.Stderr, "Incorrect format - too many : (two-dots) symbols or zero separators; want only 1")
			os.Exit(1)
		}

		splittedCardBrand := strings.Split(v, ":")
		if splittedCardBrand[0] == "" || splittedCardBrand[1] == "" {
			fmt.Fprintln(os.Stderr, "Incorrect input format of a card type: empty name or number prefix")
			os.Exit(1)
		}

		if splittedCardBrand[1] == issuerNumber[:len(splittedCardBrand[1])] {
			return splittedCardBrand[1]
		}
	}
	return ""
}

func findCardIssuerNumber(cardIssuers []string, issueIssuer string) string {
	for _, v := range cardIssuers {
		if strings.Count(v, ":") != 1 {
			fmt.Fprintln(os.Stderr, "Incorrect format - want only 1 colon in each line")
			os.Exit(1)
		}

		splittedCardIssuer := strings.Split(v, ":")
		if splittedCardIssuer[0] == "" || splittedCardIssuer[1] == "" {
			fmt.Fprintln(os.Stderr, "Incorrect input format of a card type: empty name and/or number prefix")
			os.Exit(1)
		}

		if splittedCardIssuer[0] == issueIssuer {
			return splittedCardIssuer[1]
		}
	}
	return "0"
}
