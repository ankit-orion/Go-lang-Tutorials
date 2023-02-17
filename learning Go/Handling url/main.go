package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/search?q=123&oq=123&aqs=chrome..69i57j69i60.1056j0j1&sourceid=chrome&ie=UTF-8"

func main() {
	fmt.Println("learning to handle url")
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	//here we aer creating a varoble to store the query from the url

	qparams := result.Query()
	//checking the type of query received
	fmt.Printf("The type of qparams are %T\n", qparams)
	fmt.Println(qparams["search_query"])
	fmt.Println(result.RawQuery)

}
