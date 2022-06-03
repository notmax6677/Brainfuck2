package src

import (
	"fmt"
	"strconv"
)

// tries to run the given string as an integer, which will be outputted as binary code
func runInteger(integer string) {
	// try to convert the integer to a string
	number, err := strconv.ParseFloat(integer, 32)

	// if error is anything other than nil
	if err != nil {
		fmt.Print("\nYou fucking idiot.\n\n") // tell the user they're a fucking idiot

		panic(err) // panic the error
	}

	// turn the number into an integer
	intNumber := int(number)

	// print the integer in binary form
	fmt.Print(strconv.FormatInt(int64(intNumber), 2))
	//fmt.Println(intNumber)
}
