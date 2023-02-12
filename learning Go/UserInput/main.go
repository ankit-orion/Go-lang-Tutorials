package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	// bufio package :- this is a buffer using this you can read from
	// input output and you can take input from keyboard

	// declaring a varible reader here we are creater newReader and we want to
	// read from os.Stdin
	reader := bufio.NewReader(os.Stdin)
	//since this is an println so it will automatically inject /n at the end of the code
	// this is going to be used un case of eroor handling

	fmt.Println("enter the rating for our pizza")
	// now somebody is listeing or reading from bufio
	// now we have to store that in a varaible
	// so we will use this thing ->

	// comma ok syntax || error ok
	//in go lang if something goes worng we need to store
	// the error in a varible
	// so basically we store problems or erros in a variable

	// storing
	// here we are storing err and we will continue reading the file until we hit a \n or line change
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)
	fmt.Printf("type of the rating is %T", input)

}
