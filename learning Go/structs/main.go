package main

import "fmt"

func main() {
	fmt.Println("learning struct in go lang")
	fmt.Println("we don't have class in go so we have struct in go lang")

	user1 := User{"ankit", "ankit.12010958@gmail.com", true, 20}
	fmt.Println(user1)
	// if you want to print the strcut along with its key you can do something like ths
	fmt.Printf("user 1 details are %+v\n", user1)
	fmt.Printf("\nName of the person is %v and email address of the person is %v", user1.Name, user1.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
