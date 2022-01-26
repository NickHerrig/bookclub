package main

import (
    "flag"
    "fmt"
)

func main() {
	size_of_pyramid_arg := flag.Int("size", 5, "size of pyramid")
    flag.Parse()

    size_of_pyramid := *size_of_pyramid_arg
	
	for i := 0; i < size_of_pyramid*2+1; i++ {

		line_len := 0

		if i <= size_of_pyramid {
				line_len = i
		}
	    
		if i > size_of_pyramid {
				line_len = size_of_pyramid*2-i
		}

		print_string := ""

		for j := 0; j <= line_len; j++ {
            print_string += "*"
		}

		fmt.Print(print_string, "\n")
	}
}