package main

import (
	"fmt"
)

func main() {
	fmt.Println("learning maps in go lang")
	// map[values]key

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["cpp"] = "c plus plus lang"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	// here it prints all the complete map
	fmt.Println("list of all languages ", languages)
	// if you want to print particular map and its value
	fmt.Println("map and its value ", languages["js"])

	// you can delete any map just by accessing its key
	delete(languages, "js")
	fmt.Println(languages)

	// printing the map using for loop
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

}
