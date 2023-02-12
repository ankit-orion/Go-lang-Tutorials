package main

import "fmt"

func main() {
	fmt.Println("learning if else control statement in go lang")
	loginCount := 25
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else {
		result = "bas kar bhai"
	}
	fmt.Println(result)

	if 10%2 == 0 {
		fmt.Println("it is even")
	} else {
		fmt.Println("it is not even")
	}

	// initializing and checking at the same time
	if num := 10; num < 5 {
		fmt.Println("the number is ", num)
	} else {
		fmt.Println("number is graeter than 5")
	}
}
