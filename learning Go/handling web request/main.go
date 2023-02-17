package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://twitter.com/"

func main() {
	fmt.Println("learning handling web reqest")

	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("type of response %T\n", response)
	// it is our response to close the request
	defer response.Body.Close()
	dataByte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("response receive from the web request", string(dataByte))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
