package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createPyramid(peak int) string {
	var pyramid string
	for i := 1; i < peak; i++ {
		pyramid += strings.Repeat("*", i) + "\n"
	}

	for i := peak; i > 0; i-- {
		pyramid += strings.Repeat("*", i) + "\n"
	}

	return pyramid
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./pyramids {integer}")
		os.Exit(1)
	}

	peak, err := strconv.Atoi(os.Args[1])
	if err != nil || peak < 0 {
		fmt.Println("Must provide a positive integer as an argument!")
		os.Exit(1)
	}

	fmt.Println(createPyramid(peak))
}
