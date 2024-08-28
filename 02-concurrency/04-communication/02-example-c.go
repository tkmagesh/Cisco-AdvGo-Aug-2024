package main

import (
	"fmt"
)

// share memory by communcating

func main() {

	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	result := <-ch
	fmt.Println("Result :", result)
}

// 3rd party api (Cannot modify the code)
func add(x, y int) int {
	result := x + y
	return result
}
