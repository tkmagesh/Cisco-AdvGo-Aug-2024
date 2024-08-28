/*
Write a printNos() function that will execute concurrently

PrintNos will execute PrintOddNos() and PrintEvenNos() concurrently

PrintOddNos() -> print the first 20 odd numbers with a time delay of 2 seconds between each number

PrintEvenNos() -> print the first 20 even numbers with a time delay of 3 seconds between each number

Make sure the main function exits only after all the concurrent operations are completed

main()
	-> printNos()
		-> printOddNos()
		-> printEvenNos()

DO NOT create a waitgroup at package scope
*/