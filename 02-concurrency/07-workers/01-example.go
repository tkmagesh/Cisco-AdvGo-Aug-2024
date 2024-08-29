/*
modify the program spin up finite # of workers (5 or 10 or 15)
the workers have to share the load of processing the range of numbers
*/
package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	var workerCount, start, end int
	flag.IntVar(&workerCount, "workercount", 1, "Number or workers to employ!")
	flag.IntVar(&start, "start", 0, "starting range")
	flag.IntVar(&end, "end", 0, "ending range")
	flag.Parse()
	primesCh := findPrimes(start, end, workerCount)
	for no := range primesCh {
		fmt.Println("Prime :", no)
	}
	fmt.Println("Done")
}

func findPrimes(start, end int, workerCount int) <-chan int {
	primesCh := make(chan int)
	dataCh := dataProducer(start, end)
	go func() {
		wg := &sync.WaitGroup{}
		for id := range workerCount {
			wg.Add(1)
			go primeWorker(id, wg, dataCh, primesCh)
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func dataProducer(start, end int) <-chan int {
	dataCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
		// fmt.Println("Finished producing the data...!")
	}()
	return dataCh
}

func primeWorker(workerId int, wg *sync.WaitGroup, dataCh <-chan int, primesCh chan<- int) {
	fmt.Println("Starting worker #", workerId)
	defer wg.Done()
	for no := range dataCh {
		// fmt.Printf("worker #%d, processing %d\n", workerId, no)
		if isPrime(no) {
			primesCh <- no
		}
	}
	// fmt.Println("Shutting down worker #", workerId)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
