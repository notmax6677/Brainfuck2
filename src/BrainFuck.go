package src

import "fmt"

func Brainfuck(script string) {
	// if script is a brainfuck2 file
	if verifyExtension(script) {
		// load the content of the script
		scriptContent := loadScript(script)

		processScript(scriptContent)
	} else { // otherwise
		fmt.Println(errors["wrong_file_extension"]) // print wrong file extension error
	}
}
