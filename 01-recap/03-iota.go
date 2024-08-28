package main

import "fmt"

func main() {
	/*
		const Red int = 0
		const Green int = 1
		const Blue int = 2
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	// auto generating constant values
	/*
		const (
			Red   int = iota
			Green int = iota
			Blue  int = iota
		)
	*/

	/*
		const (
			Red int = iota
			Green
			Blue
		)
	*/

	/*
		const (
			Red = iota
			Green
			Blue
		)
	*/

	/*
		const (
			Red   = iota + 2 // 0 + 2
			Green            // 1 + 2
			Blue             // 2 + 2
		)
	*/

	const (
		Red   = iota + 2 // 0 + 2
		Green            // 1 + 2
		_
		Blue
	)
	fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)

	// Mimicking Unix file permissions

	/*
		const (
			X = iota << 1
			W
			R
		)

		fmt.Println(X, W, R)
		fmt.Printf("%b %b %b\n", X, W, R)
	*/

	const (
		X = 1 << iota
		W
		R
		XW  = X | W
		WR  = W | R
		XWR = X | W | R
	)
	fmt.Printf("%b %b %b %b %b %b\n", X, W, XW, R, WR, XWR)

}
