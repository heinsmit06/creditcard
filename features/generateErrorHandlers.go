package features

import (
	"fmt"
	"os"
)

// to ensure that there is only one argument
func checkTailLength(tail []string) {
	if len(tail) != 1 {
		fmt.Fprintln(os.Stderr, "There has to be only 1 argument")
		os.Exit(1)
	}
}

// to ensure that there are only digits or asterisks
func checkInputValidity(tail []string) {
	for _, v := range tail[0] {
		if (v != '*') && (v < 48 || v > 57) {
			fmt.Fprintln(os.Stderr, "Input must consist of only digits or asterisks")
			os.Exit(1)
		}
	}
}

// to ensure that there are max 4 asterisks
func checkAsterisksQuantity(tail []string) {
	asterisksCount := countAsterisksQuantity(tail[0])
	if asterisksCount > 4 {
		fmt.Fprintln(os.Stderr, "Max. allowed number of asterisks is 4")
		os.Exit(1)
	} else if asterisksCount == 0 {
		fmt.Fprintln(os.Stderr, "Please enter at least 1 asterisk to generate a card number")
		os.Exit(1)
	}
}

// to ensure that asterisks were placed correctly
func checkAsterisksPlacement(tail []string) {
	asterisksCount := countAsterisksQuantity(tail[0])
	const errorMessage = "Incorrect placement of an asterisk"
	switch asterisksCount {
	case 1:
		for i := 0; i < len(tail[0])-1; i++ {
			if tail[0][i] == '*' {
				fmt.Fprintln(os.Stderr, errorMessage)
				os.Exit(1)
			}
		}
	case 2:
		for i := 0; i < len(tail[0])-2; i++ {
			if tail[0][i] == '*' {
				fmt.Fprintln(os.Stderr, errorMessage)
				os.Exit(1)
			}
		}
	case 3:
		for i := 0; i < len(tail[0])-3; i++ {
			if tail[0][i] == '*' {
				fmt.Fprintln(os.Stderr, errorMessage)
				os.Exit(1)
			}
		}
	case 4:
		for i := 0; i < len(tail[0])-4; i++ {
			if tail[0][i] == '*' {
				fmt.Fprintln(os.Stderr, errorMessage)
				os.Exit(1)
			}
		}
	}
}

// count asterisks quantity
func countAsterisksQuantity(cardNumber string) int {
	asterisksCount := 0
	for _, v := range cardNumber {
		if v == '*' {
			asterisksCount++
		}
	}
	return asterisksCount
}
