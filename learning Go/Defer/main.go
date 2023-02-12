package main

import "fmt"

func main() {
	fmt.Println("learning defer")
	// defer -> when a function executes it executes line by line
	// even though it is compiled and as soon as you put defer keyword
	//whatever is the next line you have written after defer will execute at the very end of the
	// functon itself

	/************************************************************************/

	defer fmt.Println("this line will be printed at the end of the program as we've used defer")
	fmt.Println(10 + 20)

	// if you have several defer then it works as LIFO - last in first out
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello my name is ankit mishra")
	myDefer()
}

// printing in reverse order using defer
func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
