/*
Create custom type that encapsulates concurrent safe state manipulate
*/
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Add(delta int) {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go Increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.count)
}

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add(1)
}
