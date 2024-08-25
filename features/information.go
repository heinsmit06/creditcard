package features

import (
	"fmt"
	"os"
	"strings"
)

func Information(stdin bool, brandFileName, issuerFileName string, cardNumbers []string) {
	fmt.Println(" stdin:", stdin)
	fmt.Println(" brand:", brandFileName)
	fmt.Println(" issuer", issuerFileName)
	fmt.Println(" tail:", cardNumbers)
	fmt.Println()

	checkBrandAndIssuerInput(brandFileName, issuerFileName)
	checkArgsLength(cardNumbers)

	for _, cardNumber := range cardNumbers {
		if checkCardValidity(cardNumber) && (len(cardNumber) >= 13 && len(cardNumber) <= 16) && (luhnsAlgorithm(strToNum(cardNumber))) {
			// read from files in a given dir; check for nil pointer
			fileContentBrands, err1 := os.ReadFile("./text/" + brandFileName)
			check(err1)
			fileContentIssuers, err2 := os.ReadFile("./text/" + issuerFileName)
			check(err2)

			// conversion of byte file contents to string type; dividing the big string into smaller ones and storing in a slice
			cardBrands := strings.Split(string(fileContentBrands), "\n")
			cardIssuers := strings.Split(string(fileContentIssuers), "\n")
			// fmt.Println(cardBrands)
			// fmt.Println(cardIssuers)

			brand := findCardBrand(cardBrands, cardNumber)
			issuer := findCardIssuer(cardIssuers, cardNumber)

			fmt.Fprintln(os.Stdout, cardNumber)
			fmt.Fprintln(os.Stdout, "Correct: yes")
			fmt.Fprintln(os.Stdout, "Card Brand:", brand)
			fmt.Fprintln(os.Stdout, "Card Issuer:", issuer)
		} else {
			fmt.Fprintln(os.Stderr, cardNumber)
			fmt.Fprintln(os.Stderr, "Correct: no")
			fmt.Fprintln(os.Stderr, "Card Brand: -")
			fmt.Fprintln(os.Stderr, "Card Issuer: -")
			fmt.Println()
		}
	}
}

func findCardBrand(cardBrands []string, cardNumber string) string {
	for _, v := range cardBrands {
		if strings.Count(v, ":") > 1 || strings.Count(v, ":") == 0 {
			fmt.Fprintln(os.Stderr, "Incorrect format - too many : (two-dots) symbols or zero separators; want only 1")
			fmt.Println("cardBrands")
			os.Exit(1)
		}

		splittedCardBrand := strings.Split(v, ":")
		if splittedCardBrand[0] == "" || splittedCardBrand[1] == "" {
			fmt.Fprintln(os.Stderr, "Incorrect input format of a card type: empty name or number prefix")
			os.Exit(1)
		}

		brandName := splittedCardBrand[0]
		startingDigits := splittedCardBrand[1]

		if startingDigits == cardNumber[:len(startingDigits)] {
			return brandName
		}
	}
	return "-"
}

func findCardIssuer(cardIssuers []string, cardNumber string) string {
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

		issuerName := splittedCardIssuer[0]
		startingDigits := splittedCardIssuer[1]

		if startingDigits == cardNumber[:len(startingDigits)] {
			return issuerName
		}

	}
	return "-"
}
