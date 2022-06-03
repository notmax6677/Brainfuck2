package main

import (
	"bf2/src"
	"os"
)

func main() {
	// get second argument, which is the string
	script := os.Args[1]

	// if the first argument is "help"
	if script == "help" {
		src.Help()
	} else {
		// brainfuck it
		src.Brainfuck(script)
	}
}
