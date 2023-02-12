package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files in go lang")
	content := "Creating files using go lang file handling: "
	// here we are creating file
	// we need to handle the error also
	// here we are receiving the created file in
	// file and just in case if there's any error while craeting the file we are also handling that error
	file, err := os.Create("./mynotes.txt")
	if err != nil {
		// panic -> panic will shut down the execution of
		//  program and will show the error
		panic(err)
	}
	// here we are using the io.WriteString function this will
	// read the content of the file and it will print the content
	// syntax - io.WriteString(File_Name, File_Content)

	// here we are checking if the file is created or not
	// as the file is created it will show the length of the file
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	// used to close the file as the file has many methods called
	// this will close the file here the name of the file is "file"
	defer file.Close() //it is recommended to use defer as we might
	// write some other code se we don't want our file to get closed
	// as defer will close the file at the end of the program
	readFile("./mynotes.txt")
}

func readFile(filename string) {
	// creation is the os operation
	// ioutil -> it has lot of operations here we are using the ReadFile operation
	// this will read the file content

	// whenever you read data it is not rad as it is
	// the data is in byte format so you need to convert the data
	// in the required format you want to read
	// here we want to read the data in string format
	// alos if you'll read data from internet you will get data in
	// byte format
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// here we are printing the data without converting it into the string
	// this will print the data in bytes


	fmt.Println("printing the file data in bytes")
	fmt.Println("text data inside the file is \n", databyte)
	fmt.Println("printing the file data after converting it to string ")
	fmt.Println("text data inside the file is \n", string(databyte))
	
}
