package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in go lang")
	greeter()

	// function to create two numbers using function
	result := sum(20, 30)
	fmt.Println(result)
	fmt.Println("adding so many numbers")
	proRes := proAdder(3, 23, 141, 4, 424, 2, 53, 53, 3)
	fmt.Println(proRes)
}

// youj cannot create a function inside a function
// you can call a function inside a function
func greeter() {
	fmt.Println("namastey from India")
}
func sum(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
