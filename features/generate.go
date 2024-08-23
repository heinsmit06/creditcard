package features

import (
	"fmt"
	"os"
)

func Generate(pick bool, sliceOfArgs []string) {
	fmt.Println(" pick:", pick)
	fmt.Println(" tail:", sliceOfArgs)

	checkTailLength(sliceOfArgs)
	checkInputValidity(sliceOfArgs)
	checkAsterisksQuantity(sliceOfArgs)
	checkAsterisksPlacement(sliceOfArgs)

	generateNums(sliceOfArgs[0])
}

func generateNums(tail string) {
	asterisksCount := countAsterisksQuantity(tail)
	switch asterisksCount {
	case 1:
		intSlice := strToNum(tail[:(len(tail) - 1)])
		for i := 0; i < 10; i++ {
			clone := sliceGetClone(intSlice)
			clone = append(clone, i)
			if luhnsAlgorithm(clone) {
				for _, v := range clone {
					fmt.Fprintln(os.Stdout, v)
				}
			}
			fmt.Println("reached")
		}
		// case 2:

		// case 3:

		// case 4:
	}
}

func sliceGetClone(src []int) []int {
	slc := make([]int, 0)
	slc = src
	return slc
}
