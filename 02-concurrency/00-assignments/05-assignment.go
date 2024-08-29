/*
modify the program as below
1. genNos() will genarate the number but print the generated numbers in the main()
2. keep generating the numbers until the user hits ENTER key
*/
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

func genNos(wg *sync.WaitGroup) {
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
