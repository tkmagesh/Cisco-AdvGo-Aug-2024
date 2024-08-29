/*
modify the program as below
 1. genNos() will genarate the number but print the generated numbers in the main()
 2. keep generating the numbers until kill signal is received
 3. to send kill signal
    kill -s INT <pid>
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println("process id :", os.Getpid())
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	nosCh := genNos(stopCh)
	for no := range nosCh {
		fmt.Println("no :", no)
	}
	fmt.Println("Done!")
}

func genNos(stopCh chan os.Signal) <-chan int {
	nosCh := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				signal.Stop(stopCh)
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
