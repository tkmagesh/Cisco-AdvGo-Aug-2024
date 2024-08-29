/*
modify the program to adopt "share memory by communicating"
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primesCh := findPrimes(1000, 20000)
	for no := range primesCh {
		fmt.Println("Prime :", no)
	}
	fmt.Println("Done")
}

func findPrimes(start, end int) <-chan int {
	primesCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if isPrime(no) {
					primesCh <- no
				}
			}()
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
