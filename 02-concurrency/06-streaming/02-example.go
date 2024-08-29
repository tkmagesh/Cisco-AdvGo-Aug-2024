package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		fmt.Println("Channel closed!")
		break
	}

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
