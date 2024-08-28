package main

import (
	"fmt"
)

// share memory by communcating

// consumer
func main() {
	ch := add(100, 200)
	result := <-ch
	fmt.Println("Result :", result)
}

// producer
func add(x, y int) chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result
	}()
	return ch
}
