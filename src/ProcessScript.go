package src

import (
	"fmt"
	"strings"
	"time"
)

// process the script
func processScript(script string) {
	// remove all newlines
	script = strings.Replace(script, "\n", "", -1)

	// this next if statement is almost completely useless but a great way to be annoying lol

	// if last character of script is new command character
	if string(script[len(script)-1]) == syntax["new_command"] {
		script = strings.TrimSuffix(script, syntax["new_command"]) // remove the character and move on

		// get commands by splitting the script with the new command character
		commands := strings.Split(script, syntax["new_command"])

		// iterate thru commands
		for i := 0; i < len(commands); i++ {
			processCommand(commands[i])
		}
	} else { // otherwise throw error
		fmt.Println(errors["last_char_of_script"])
	}
}

// process a single command and take action accordingly
func processCommand(command string) {
	// switch statement for first character of command
	switch string(command[0]) {

	// if first char is the begin character char
	case syntax["character"]:
		runCharacter(strings.TrimPrefix(command, syntax["character"])) // run the character, input the argument as the command is, except remove the first char

	// if first char is the begin integer char
	case syntax["integer"]:
		runInteger(strings.TrimPrefix(command, syntax["integer"])) // run the character, input the argument as the command is, except remove the first char

	// if first char is the begin add char
	case syntax["addition"]:
		add(strings.TrimPrefix(command, syntax["addition"]))

	// if first char is the begin subtract char
	case syntax["subtraction"]:
		subtract(strings.TrimPrefix(command, syntax["subtraction"]))

	default:
		fmt.Println(errors["first_char_of_command"])
	}

	// sleep 0.5 seconds for every character that needs to be printed (just for extra brainfuckery)
	time.Sleep(time.Millisecond * 500)
}
