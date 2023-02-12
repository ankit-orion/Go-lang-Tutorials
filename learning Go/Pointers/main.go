package main

import "fmt"

func main() {
	fmt.Println("welcome to the class of pointers")
	/***************************************************
		// this will give default value to nil
		// var ptr *int // this pointer will store values of type integer

		// fmt.Println("value of pointer is ", ptr)
	*******************************************/

	mynumber := 23
	var ptr = &mynumber                       // here & will store the address of mynumber in ptr
	fmt.Println("value of pointer is ", ptr)  // to print the address of ptr
	fmt.Println("value of pointer is ", *ptr) // to print the value stored in ptr

	*ptr = *ptr + 10 // this will access the value stored in ptr
	// and then we are updating the value to *ptr = 23 + 10
	// so new value will be 23 + 10 = 33
	fmt.Println("adress after updating the value in pointer will remian the same ", ptr)
	fmt.Println("updated value stoed in pointer is ", *ptr)

}
