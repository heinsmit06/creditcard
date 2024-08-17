package features

import (
	"fmt"
	"os"
)

func Validate(cardNumber string) {
	if len(cardNumber) >= 13 {
		// Luhn's Algorithm to validate the card number
		if luhnsAlgorithm(cardNumber) {
			fmt.Fprintln(os.Stdout, "OK")
		} else {
			fmt.Println("\nlen is correct")
			fmt.Fprintln(os.Stderr, "INCORRECT")
		}
	} else {
		fmt.Fprintln(os.Stderr, "INCORRECT")
	}
}

func luhnsAlgorithm(cardNumber string) bool {
	slc := []int{}
	for _, digit := range cardNumber {
		slc = append(slc, int(digit-48)) // writing digits in card number to a slice
	}

	sliceLength := len(slc) // taking constant length, because append will change the slice's length during execution
	// doubling the first and every other digit
	for i := 0; i < sliceLength; i = i + 2 {

		slc[i] = slc[i] * 2
		// checking if the doubling resulted in double digit number; adding them separately, if positive
		if slc[i] >= 10 {
			firstDigit, secondDigit := slc[i]/10, slc[i]%10
			slc[i] = firstDigit
			slc = append(slc, secondDigit)
		}

	}

	sum := 0
	for i := 0; i < len(slc); i++ {
		sum += slc[i]
	}

	// if the final sum is divisible by 10, then the credit card is valid
	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}
