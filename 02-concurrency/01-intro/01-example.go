package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling through the scheduler to execute
	f2()
	// poor man's synchronization techniques. DO NOT USE THEM
	// block the execution so that the scheduler can look for other goroutines scheduled and execute them
	time.Sleep(3 * time.Second)
	// fmt.Scanln()

}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
