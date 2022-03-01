package main

/*
Bookclub Example Program #1:

Did you just fall into an Escher painting? Are you in Egypt? Things have definitely gotten sideways and pointy. There is clearly a pyramid off in the distance, but it appears to be flipped! Also, its size appears to vary based on how hard you squint (don't wanna get too much sand in your eyes).

Write a program that accepts a single integer argument (must be greater than zero). Based on this value, print out a sideways pyramid of asterisks that peaks at the provided size.

Here are a few examples:
```bash
go run pyramids.go 1
*

go run pyramids.go 2
*
**
*

go run pyramids.go 5
*
**
***
****
*****
****
***
**
*
```
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	if size <= 0 {
		fmt.Println("size too low")
		os.Exit(1)
	}

	for i := 1; i < size; i++ {
		line := ""
		for j := 0; j < i; j++ {
			line += "*"
		}
		fmt.Println(line)
	}

	for i := size; i > 0; i-- {
		line := ""
		for j := 0; j < i; j++ {
			line += "*"
		}
		fmt.Println(line)
	}
}
