package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of go lang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	// creating time

	createdDate := time.Date(2023, time.March, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println("created date is ", createdDate)
	// to make date look pretty
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
