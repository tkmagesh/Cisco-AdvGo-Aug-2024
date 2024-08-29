/*
modify the program as below
1. genNos() will genarate the number but print the generated numbers in the main()
2. keep generating the numbers until the user hits ENTER key
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	nosCh := genNos(stopCh)
	for no := range nosCh {
		fmt.Println("no :", no)
	}
	fmt.Println("Done!")
}

func genNos(stopCh <-chan struct{}) <-chan int {
	nosCh := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				fmt.Println("stop signal received")
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("no :", i)
			}
		}
		close(nosCh)
	}()
	return nosCh
}
