package main

import (
	"fmt"
	"sync"
)

// share memory by communcating

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
