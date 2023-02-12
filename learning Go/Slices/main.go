package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcone to the class of slices")
	var fruits = []string{"apples", "tomato", "peach"}
	fmt.Println("values in the fruit", fruits)
	// if you want to add values in the slices you can use append keyword
	fmt.Println("adding new values to the fruit slice")
	fruits = append(fruits, "banana", "mango")
	fmt.Println("after append the value in fruit are ", fruits)
	fruits = append(fruits[2:3])
	fmt.Println(fruits)

	// to [_:3] -> this will start from 0 to 3-1
	// make will take two arguments here first is what kind of
	// data type you want to create and second is the size of data type
	// so here we are creating a slice of size 4
	highScore := make([]int, 4)
	highScore[0] = 834
	highScore[1] = 345
	highScore[2] = 45
	highScore[3] = 956
	/*
		highScore[4] = 554 -> this will throw an error
		"index out of range [4] with length 4"
	*/
	fmt.Println(highScore)
	// if we will use append keyword to add values in the slice
	// then go lang behaves differently here
	highScore = append(highScore, 4334, 4333, 13242, 25242) // this append will add these values in the slice
	// even though the size is assigned to 4 it will atomatically changes its size
	fmt.Println("printing the values of highscore after append", highScore)

	// we can use sort to sort any slice
	fmt.Println("before sort", highScore)
	sort.Ints(highScore)
	fmt.Println("after sort ", highScore)

}
