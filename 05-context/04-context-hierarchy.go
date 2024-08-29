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

	wg1 := &sync.WaitGroup{}

	// overriding root-key
	newRootValCtx := context.WithValue(ctx, "root-key", "new-root-value")
	randValCtx := context.WithValue(newRootValCtx, "random-key", "random-value")

	wg1.Add(1)
	evenCtx, cancel := context.WithTimeout(randValCtx, 5*time.Second)
	defer cancel()
	go printEven(evenCtx, wg1)

	wg1.Add(1)
	oddCtx, cancel := context.WithTimeout(randValCtx, 7*time.Second)
	defer cancel()
	go printOdd(oddCtx, wg1)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[printRandom] cancellation signal received")
			if ctx.Err() == context.Canceled {
				fmt.Println("[printRandom] programmatic cancellation occurred")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[printRandom] cancelled due to timeout")
			}
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			no := rand.Intn(100)
			fmt.Println("[printRandom] no :", no)
		}
	}
	wg1.Wait()
}

func printEven(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printEven] root-key :", ctx.Value("root-key"))
	fmt.Println("[printEven] random-key :", ctx.Value("random-key"))
LOOP:
	for i := 0; ; i += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[printEven] cancellation signal received")
			if ctx.Err() == context.Canceled {
				fmt.Println("[printEven] programmatic cancellation occurred")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[printEven] cancelled due to timeout")
			}
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("[printEven] even :", i)
		}
	}
}

func printOdd(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printOdd] root-key :", ctx.Value("root-key"))
	fmt.Println("[printOdd] random-key :", ctx.Value("random-key"))
LOOP:
	for i := 1; ; i += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[printOdd] cancellation signal received")
			if ctx.Err() == context.Canceled {
				fmt.Println("[printOdd] programmatic cancellation occurred")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[printOdd] cancelled due to timeout")
			}
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("[printOdd] even :", i)
		}
	}
}
