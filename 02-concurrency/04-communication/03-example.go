package main

import "fmt"

func main() {
	// ver 1.0
	/*
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	// ver 2.0
	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	// ver 3.0
	/*
		ch := make(chan int)
		go func() {
			ch <- 100 //(2.NB)
		}()
		data := <-ch //(1.B)(3.UB)
		fmt.Println(data)
	*/

	// ver 4.0
	ch := make(chan int)
	ch <- 100
	// receiving and printing has to happen in a goroutine
	data := <-ch
	fmt.Println(data)

}
