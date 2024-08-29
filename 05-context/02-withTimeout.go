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
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel()
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	go printRandom(timeoutCtx, wg)
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
			if ctx.Err() == context.Canceled {
				fmt.Println("programmatic cancellation occurred")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("cancelled due to timeout")
			}
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			no := rand.Intn(100)
			fmt.Println("[printRandom] no :", no)
		}

	}
}
