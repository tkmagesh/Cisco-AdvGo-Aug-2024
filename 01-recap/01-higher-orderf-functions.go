/*
HOF
	- Functions like data
	- Assign functions as values to variables
	- Pass functions as arguments to other functions
	- Return functions as return values
*/

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// ver 2.0
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply Result :", x*y)
		}, 100, 200)
	*/

	// ver 3.0
	/*
		logAdd := logWrapper(add)
		logAdd(100, 200)

		logSubtract := logWrapper(subtract)
		logSubtract(100, 200)

		logMultiply := logWrapper(func(x, y int) {
			fmt.Println("Multiply Result :", x*y)
		})
		logMultiply(100, 200)
	*/

	// ver 4.0
	/*
		logAdd := logWrapper(add)
		add := profileWrapper(logAdd)

		logSubtract := logWrapper(subtract)
		subtract := profileWrapper(logSubtract)
	*/

	add := profileWrapper(logWrapper(add))
	subtract := profileWrapper(logWrapper(subtract))

	add(100, 200)
	subtract(100, 200)
	profileWrapper(logWrapper(func(i1, i2 int) {
		fmt.Println("Multiply Result :", i1*i2)
	}))(100, 200)
}

/* ver 4.0 */
func profileWrapper(op Operation) Operation {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		log.Println("[profile] Elapsed :", elapsed)
	}
}

/* ver 3.0 */
/*
func logWrapper(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("[log] Operation started")
		op(x, y)
		log.Println("[log] Operation completed")
	}
}
*/
type Operation func(int, int)

func logWrapper(op Operation) Operation {
	return func(x, y int) {
		log.Println("[log] Operation started")
		op(x, y)
		log.Println("[log] Operation completed")
	}
}

/* ver 2.0 */
/*
func logAdd(x, y int) {
	log.Println("[log] Operation started")
	add(x, y)
	log.Println("[log] Operation completed")
}

func logSubtract(x, y int) {
	log.Println("[log] Operation started")
	subtract(x, y)
	log.Println("[log] Operation completed")
}
*/

// Applying "commanility variability" for the above
func logOperation(op func(int, int), x, y int) {
	log.Println("[log] Operation started")
	op(x, y)
	log.Println("[log] Operation completed")
}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
