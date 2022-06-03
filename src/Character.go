package src

import (
	"fmt"
	"strconv"
	"strings"
)

// try to run the given binary form of a character, and output it as a string
func runCharacter(character string) {
	// if the character is a space character
	if character == syntax["space_character"] {
		fmt.Print(" ") // print a space
	} else { // else
		// boolean equation to see if the character is lowercase
		lowercase := string((character[len(character)-1])) == "l"

		// character that will be used for conversion
		actualCharacter := character

		// if lowercase is true
		if lowercase {
			// remove the last character from the actualCharacter string
			actualCharacter = strings.TrimSuffix(actualCharacter, "l")
		}

		// check if length of character string is 8
		properLength := len(actualCharacter) == 8

		// amount of characters that are either 1 or 0
		var properCharacters int

		// iterate through actualCharacter
		for i := 0; i < len(actualCharacter); i++ {
			// if the indexed character is either a 0 or a 1
			if string(actualCharacter[i]) == "0" || string(actualCharacter[i]) == "1" {
				properCharacters++ // add 1 to properCharacters
			}
		}

		// if there are 8 characters in actualCharacter and they're all 1s and 0s
		if properLength && properCharacters == 8 {
			// convert binary string to actual binary then a letter
			b := make([]byte, 0)
			for _, s := range strings.Fields(actualCharacter) {
				n, _ := strconv.ParseUint(s, 2, 8)
				b = append(b, byte(n))
			}

			// if lowercase is true
			if lowercase {
				fmt.Print(strings.ToLower(string(b))) // print lowercase version of it
			} else { // otherwise
				fmt.Print(string(b)) // print uppercase (default) version
			}

		} else { // otherwise
			fmt.Println(errors["wrong_binary_input"]) // show wrong binary input error message
		}
	}
}

// found solution for converting binary to characters on this reddit post lol
// https://www.reddit.com/r/golang/comments/7pwlh6/comment/dskjctu/?utm_source=share&utm_medium=web2x&context=3
