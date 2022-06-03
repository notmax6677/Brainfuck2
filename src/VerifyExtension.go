package src

// verifies the file extension and returns true or false based on whether or not the file extension is a brainfuck2 file
// true = yes its a brainfuck file
// false = no its not
func verifyExtension(script string) bool {
	// if last x characters are equivalent to the given file extension
	if script[len(script)-len(syntax["file_extension"]):] == syntax["file_extension"] {
		return true
	} else { // else
		return false
	}
}
