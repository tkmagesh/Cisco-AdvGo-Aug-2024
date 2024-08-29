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
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	timeoutCtx, cancel := context.WithTimeout(valCtx, 10*time.Second)

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
	fmt.Println("[printRandom] root-key :", ctx.Value("root-key"))
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
