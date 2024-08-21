package features

import (
	"flag"
	"fmt"
)

func Generate(pick bool, sliceOfArgs ...string) {
	pickPtr := flag.Bool("pick", false, "randomly picks a single entry from possible variants")
	flag.Parse()

	fmt.Println("pick:", *pickPtr)
	fmt.Println("tail:", flag.Args())
}
