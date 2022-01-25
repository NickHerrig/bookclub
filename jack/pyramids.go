package main

import (
    "os"
    "fmt"
    "strconv"
)


func main() {
    height, _ := strconv.Atoi(os.Args[1])
    i := 1
    for i <= height {
        print_star(i)
        fmt.Printf("\n")
        i = i+1

    }
    i = height-1
    for i > 0 {
        print_star(i)
        fmt.Printf("\n")
        i = i-1
    }
}

func print_star(num int) {
    i := 0
    for i < num {
      fmt.Printf("*")
      i = i+1
    }
}
