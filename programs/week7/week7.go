package main

import (
	"fmt"
)

func xrange(n int) chan int {
	channel := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

func main() {
	for i := range xrange(5) {
		fmt.Println(i)
	}
}
