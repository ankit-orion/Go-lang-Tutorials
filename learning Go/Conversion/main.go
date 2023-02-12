package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString(('\n'))
	// this will throw en error because the moment we will hit enter 4 comes with \n
	//so to avoid this we use
	fmt.Println("thanks for the rating ", input)

	// now what if the user wants to add one more rating to the rating
	// which he has given in that we cannot jsut simply do like
	// numRating = input + 1
	// as input is of tpe string so in that case we need to convert it

	// now we have to remove the \n

	// so we will se strins.TrimSpace(input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating ", numRating+1)
	}

	// other functions of strings that we can use are
	var learning string = "AnkIt MishRA"
	fmt.Println(strings.ToLower(learning))
}
