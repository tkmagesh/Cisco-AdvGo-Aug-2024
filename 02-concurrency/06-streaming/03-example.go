package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Channel closed!")

}

// producer
func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		fmt.Println("count :", count)
		fmt.Scanln()
		for i := range count {
			ch <- (i + 1) * 10
		}
		close(ch)
	}()
	return ch
}
