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
	}

	//Iterate up the pyramid
	for i := 1; i <= input; i++ {
		for currentStar := 0; currentStar < i; currentStar++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//Iterate down the pyramid
	for i := input - 1; i > 0; i-- {
		for currentStar := 0; currentStar < i; currentStar++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
