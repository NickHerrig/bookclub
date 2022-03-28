package main

import (
	"fmt"
)

func main() {
	for i := range xrange(5) {
		fmt.Println(i)
	}
}

func xrange(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c<-i
		}
		close(c)
	}()

	return c
}
