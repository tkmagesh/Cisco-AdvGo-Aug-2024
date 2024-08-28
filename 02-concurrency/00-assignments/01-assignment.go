/*
Write a printNos() function that will execute concurrently

PrintNos will execute PrintOddNos() and PrintEvenNos() concurrently

PrintOddNos() -> print the first 20 odd numbers with a time delay of 2 seconds between each number

PrintEvenNos() -> print the first 20 even numbers with a time delay of 3 seconds between each number

Make sure the main function exits only after all the concurrent operations are completed

main()
	-> printNos()
		-> printOddNos()
		-> printEvenNos()

DO NOT create a waitgroup at package scope
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
	go PrintNos(wg)
	wg.Wait()
	fmt.Println("Done")
}

func PrintNos(wg *sync.WaitGroup) {
	defer wg.Done()
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go PrintEvenNos(wg1)

	wg1.Add(1)
	go PrintOddNos(wg1)
	wg1.Wait()
}

func PrintEvenNos(wg *sync.WaitGroup) {
	defer wg.Done()
	for i, count := 0, 20; count > 0; count-- {
		i += 2
		fmt.Println("Even No :", i)
		time.Sleep(3 * time.Second)
	}
}

func PrintOddNos(wg *sync.WaitGroup) {
	defer wg.Done()
	for i, count := 1, 20; count > 0; count-- {
		i += 2
		fmt.Println("Odd No :", i)
		time.Sleep(2 * time.Second)
	}
}
