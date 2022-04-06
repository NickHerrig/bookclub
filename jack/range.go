package main

import (
  "fmt"
)

func sendData(ch chan int, value int) {
  for i:=0; i<value; i++ {
    ch <- i
  }
  close(ch)
}

func xrange(n int) chan int {
  // IMPLEMENT
  ch := make(chan int)
  go sendData(ch, n)
  return ch
}

func main() {
  for i := range xrange(5) {
    fmt.Println(i)
  }
}
