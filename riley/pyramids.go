/*
Bookclub Example Program #1:

Did you just fall into an Escher painting? Are you in Egypt?
Things have definitely gotten sideways and pointy.
There is clearly a pyramid off in the distance, but it appears to be flipped!
Also, its size appears to vary based on how hard you squint (don't wanna get too much sand in your eyes).

Write a program that accepts a single integer argument (must be greater than zero).
Based on this value, print out a sideways pyramid of asterisks that peaks at the provided size.
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	num := flag.Int("n", 0, "# of rows")
	char := flag.String("c", "*", "character")
	flag.Parse()

	for i := 1; i <= *num; i++ {
		fmt.Printf("%s\n", strings.Repeat(*char, i))
	}
	for j := *num - 1; j > 0; j-- {
		fmt.Printf("%s\n", strings.Repeat(*char, j))
	}
}
