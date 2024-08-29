package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go printNos(wg)
	wg.Wait()
}

/*
func printNos(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("no :", i)
		elapsed := time.Since(start)
		if elapsed >= time.Duration(10*time.Second) {
			fmt.Println("Elapsed")
			break
		}
	}
}
*/

func printNos(wg *sync.WaitGroup) {
	defer wg.Done()
	// timeoutCh := timeout(10 * time.Second)
	timeoutCh := time.After(10 * time.Second)
LOOP:
	for i := 0; ; i++ {
		select {
		case <-timeoutCh:
			fmt.Println("Elapsed")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("no :", i)
		}
	}
}
