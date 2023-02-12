package main

import "fmt"

func main() {
	fmt.Println("arrays in go lang")

	// decaring an array on go

	var fruitList [4]string
	fruitList[0] = "banana"
	fruitList[1] = "litchi"
	fruitList[2] = "hello"
	fruitList[3] = "apple"
	fmt.Println("printing the first value of the array", fruitList[0])
	fmt.Println("printing the first value of the array", fruitList[1])
	fmt.Println("printing the first value of the array", fruitList[2])
	fmt.Println("printing the first value of the array", fruitList[3])

	// if you want to print all the values of the array simple print the name of the array
	fmt.Println("printing the array just by writing its name", fruitList)

	// printing the values of array
	fmt.Println("length of the array is ", len(fruitList))

	vegetables := [4]string{"potato", "karela", "gobhi", "greeen potato"}
	fmt.Println("priting the vegetables", vegetables)

}
