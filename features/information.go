package features

import (
	"fmt"
	"os"
)

func Information(stdin bool, brand, issuer string, cardNumbers []string) {
	fmt.Println(" stdin:", stdin)
	fmt.Println(" brand:", brand)
	fmt.Println(" issuer", issuer)
	fmt.Println(" tail:", cardNumbers)

	checkBrandAndIssuerInput(brand, issuer)
	checkArgsLength(cardNumbers)

	for _, cardNumber := range cardNumbers {
		if checkCardValidity(cardNumber) && (len(cardNumber) >= 13 && len(cardNumber) <= 16) && (luhnsAlgorithm(strToNum(cardNumber))) {
			// cardBrand
			// CardIssuer

			fmt.Fprintln(os.Stdout, cardNumber)
			fmt.Fprintln(os.Stdout, "Correct: yes")
			fmt.Fprintln(os.Stdout, "Card Brand:")
			fmt.Fprintln(os.Stdout, "Card Issuer:")
		} else {
			fmt.Fprintln(os.Stderr, cardNumber)
			fmt.Fprintln(os.Stderr, "Correct: no")
			fmt.Fprintln(os.Stderr, "Card Brand: -")
			fmt.Fprintln(os.Stderr, "Card Issuer: -")
		}
	}
}
