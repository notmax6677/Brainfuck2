package src

import (
	"fmt"
	"strconv"
	"strings"
)

// add two numbers
func add(command string) {
	// if number splitter exists in equation/command
	if strings.Contains(command, syntax["number_splitter"]) {
		// split the command in two
		twoNumbers := strings.Split(command, syntax["number_splitter"])

		// try to convert the first number from string to int
		num1, err1 := strconv.ParseFloat(twoNumbers[0], 32)

		// try to convert the second number from string to int
		num2, err2 := strconv.ParseFloat(twoNumbers[1], 32)

		// if error 1 isn't nil
		if err1 != nil {
			fmt.Print("\nYou fucking idiot.\n\n") // call the user a fucking idiot

			panic(err1) // panic the message
		}

		// if error 2 isn't nil
		if err2 != nil {
			fmt.Print("\nYou fucking idiot.\n\n") // call the user a fucking idiot

			panic(err2) // panic the message
		}

		// print the two values added, while also turning them into integers
		fmt.Print(int(num1) + int(num2))

	} else { // else
		fmt.Println(errors["no_number_splitter"]) // print no number splitter error message
	}
}

// subtract two numbers
func subtract(command string) {
	// if number splitter exists in equation/command
	if strings.Contains(command, syntax["number_splitter"]) {
		// split the command in two
		twoNumbers := strings.Split(command, syntax["number_splitter"])

		// try to convert the first number from string to int
		num1, err1 := strconv.ParseFloat(twoNumbers[0], 32)

		// try to convert the second number from string to int
		num2, err2 := strconv.ParseFloat(twoNumbers[1], 32)

		// if error 1 isn't nil
		if err1 != nil {
			fmt.Print("\nYou fucking idiot.\n\n") // call the user a fucking idiot

			panic(err1) // panic the message
		}

		// if error 2 isn't nil
		if err2 != nil {
			fmt.Print("\nYou fucking idiot.\n\n") // call the user a fucking idiot

			panic(err2) // panic the message
		}

		// print the two values subtracted, while also turning them into integers
		fmt.Print(int(num1) - int(num2))

	} else { // else
		fmt.Println(errors["no_number_splitter"]) // print no number splitter error message
	}
}
