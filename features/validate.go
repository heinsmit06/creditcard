package features

import (
	"fmt"
	"os"
)

func Validate(stdin bool, cardNumbers ...string) {
	fmt.Println(" stdin:", stdin)
	fmt.Println(" tail:", sliceOfArgs)

	for _, cardNumber := range cardNumbers {
		if checkCardValidity(cardNumber) {
			if len(cardNumber) >= 13 && len(cardNumber) <= 16 {
				// Luhn's Algorithm to validate the card number
				if luhnsAlgorithm(strToNum(cardNumber)) {
					fmt.Fprintln(os.Stdout, "OK")
				} else {
					fmt.Fprintln(os.Stderr, "INCORRECT")
				}
			} else {
				fmt.Fprintln(os.Stderr, "Incorrect input")
			}
		} else {
			fmt.Fprintln(os.Stderr, "Incorrect input")
		}
	}
}

func checkCardValidity(cardNumber string) bool {
	// to check for cases when input contains something other than numbers
	for _, unit := range cardNumber {
		if unit < 48 || unit > 57 {
			return false
		}
	}
	return true
}

func strToNum(cardNumber string) []int {
	slc := []int{}
	for _, digit := range cardNumber {
		slc = append(slc, int(digit-48)) // writing digits in card number to a slice
	}
	return slc
}

func luhnsAlgorithm(cardNumber []int) bool {
	sum := 0
	for i := len(cardNumber) - 1; i >= 0; i-- {
		n := cardNumber[i]
		if (len(cardNumber)-1-i)%2 == 1 {
			n = n * 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return sum%10 == 0
}
