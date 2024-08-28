/*
modify the program to take advantage of go concurrency
1. execute isPrime() concurrently
2. Print the prime numbers in the main()
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := findPrimes(1000, 2000)
	for _, no := range primes {
		fmt.Println("Prime :", no)
	}
	fmt.Println("# of primes :", len(primes))
	fmt.Println("Done")
}

func findPrimes(start, end int) []int {
	var primes []int
	var mutex sync.Mutex
	wg := &sync.WaitGroup{}

	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				mutex.Lock()
				{
					primes = append(primes, no)
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
