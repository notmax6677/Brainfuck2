package src

import (
	"fmt"
	"io/ioutil"
)

// loads the script and returns a string containing the contents of the script
func loadScript(script string) string {
	// read the file
	content, err := ioutil.ReadFile("./" + script)

	// if error is equal to something other than nil
	if err != nil {
		fmt.Print("\nYou fucking idiot.\n\n") // call the user a fucking idiot
		panic(err)                            // panic and print the error
	}

	// return stringified version of content
	return string(content)
}
