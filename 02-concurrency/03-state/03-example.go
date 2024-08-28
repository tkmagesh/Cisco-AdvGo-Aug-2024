/*
using sync/atomic apis for atomic operations
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go Increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&count, 1)
}
