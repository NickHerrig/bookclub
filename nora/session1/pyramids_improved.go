package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Provided input was not an int!")
		return;
	}

	//Iterate up the pyramid
	for i := 1; i <= input; i++ {
		printPyramidRow(i)
	}

	//Iterate down the pyramid
	for i := input - 1; i > 0; i-- {
		printPyramidRow(i)
	}
}

//Prints a row of stars, given the desired number of stars
func printPyramidRow(numberOfStars int) {
	for currentStar := 0; currentStar < numberOfStars; currentStar++ {
		fmt.Print("*")
	}
	fmt.Println("")
}
