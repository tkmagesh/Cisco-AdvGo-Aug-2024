package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	go printRandom(cancelCtx, wg)
	wg.Wait()
	fmt.Println("Done!")
}

func printRandom(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancellation signal received")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			no := rand.Intn(100)
			fmt.Println("[printRandom] no :", no)
		}

	}
}
