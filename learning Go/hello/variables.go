package main

import "fmt"

// you cannot use gloval ariables like this
// total := 5000
// to use this outside you can initialize like this
var total int = 4520

// you can declare constant too
const loginToken string = "fsjscshcshfw" // public variable

//here first letter is capital so it a public variable it is
// similar to writing public: in cpp
func main() {
	var username string = "ankit"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)
	var boolval bool = false
	fmt.Println(boolval)
	fmt.Printf("variable is of type : %T \n", boolval)

	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("variable is of type : %T \n", anothervar)

	var anotherstr string
	fmt.Println(anotherstr)
	fmt.Printf("variable is of type : %T \n", anotherstr)

	// here we aven't intitalized the type pf the variable
	// so the lexar will automatically decide the varaible type itself

	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("variable is of type : %T \n", website)

	// no var type
	numberofuser := 500
	fmt.Println(numberofuser)
	// printing gloabl variable here
	fmt.Println(total)

	fmt.Println(loginToken)

}
