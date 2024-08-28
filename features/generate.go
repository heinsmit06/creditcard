package features

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
)

func Generate(pick bool, sliceOfArgs []string) {
	// fmt.Println(" pick:", pick)
	// fmt.Println(" tail:", sliceOfArgs)

	checkTailLength(sliceOfArgs)
	checkInputValidity(sliceOfArgs)
	checkAsterisksQuantity(sliceOfArgs)
	checkAsterisksPlacement(sliceOfArgs)
	generateCheckInputLength(sliceOfArgs)
	generateNums(pick, sliceOfArgs[0])
}

func generateNums(pick bool, tail string) {
	asterisksCount := countAsterisksQuantity(tail)
	storage := []string{}

	switch asterisksCount {
	case 1:
		intSlice := strToNum(tail[:(len(tail) - 1)])
		for i := 0; i < 10; i++ {
			clone := slices.Clone(intSlice)
			clone = append(clone, i)
			if luhnsAlgorithm(clone) {
				storage = append(storage, intSliceToStr(clone))
			}
		}
	case 2:
		intSlice := strToNum(tail[:(len(tail) - 2)])
		for i := 0; i < 100; i++ {
			clone := slices.Clone(intSlice)
			clone = append(clone, i/10)
			clone = append(clone, i%10)
			if luhnsAlgorithm(clone) {
				storage = append(storage, intSliceToStr(clone))
			}
		}
	case 3:
		intSlice := strToNum(tail[:(len(tail) - 3)])
		for i := 0; i < 1000; i++ {
			clone := slices.Clone(intSlice)
			clone = append(clone, i/100)
			clone = append(clone, (i/10)%10)
			clone = append(clone, i%10)
			if luhnsAlgorithm(clone) {
				storage = append(storage, intSliceToStr(clone))
			}
		}
	case 4:
		intSlice := strToNum(tail[:(len(tail) - 4)])
		for i := 0; i < 10000; i++ {
			clone := slices.Clone(intSlice)
			clone = append(clone, i/1000)
			clone = append(clone, (i/100)%10)
			clone = append(clone, (i/10)%10)
			clone = append(clone, i%10)
			if luhnsAlgorithm(clone) {
				storage = append(storage, intSliceToStr(clone))
			}
		}
	}

	// if pick is false, print every card number in our storage slice
	if !pick {
		for _, v := range storage {
			fmt.Fprintln(os.Stdout, v)
		}
	} else {
		idx := rand.Intn(len(storage))
		fmt.Fprintln(os.Stdout, storage[idx])
	}
}

func intSliceToStr(intSlice []int) string {
	valuesText := []string{}

	// Create a string slice using itoa
	// and append strings to it
	for _, v := range intSlice {
		intString := itoa(v)
		valuesText = append(valuesText, intString)
	}

	// Join the string slice
	result := strings.Join(valuesText, "")
	return result
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	sign := ""

	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+48)) + result
		n /= 10
	}

	return sign + result
}
