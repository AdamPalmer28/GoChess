package cli_engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func (cfg *Config) Run() bool {
	
	scanner := bufio.NewScanner(os.Stdin)
	var user_cmd string

    // Prompt the user for input
    fmt.Print("Enter a line of text: ")


    // Read the next line from standard input
    if scanner.Scan() {
		// Get the text that the user entered
		user_cmd = scanner.Text()

        // Print the input line
        fmt.Println("\nChess Engine:", user_cmd)
    } else {
        // If an error occurred while reading input
        fmt.Println("Error reading input:", scanner.Err())
    }

	// lowercase the input
	user_cmd = strings.ToLower(user_cmd)
	inputs := strings.Split(user_cmd, " ")
	
	cmd := inputs[0]

	gs := cfg.gs

	if cmd == "print" {

		gs.Board.Print()
	}
 
	// move a1a2
	if cmd == "move" {
		
		move := inputs[1]

		start := move[0:2]
		end := move[2:] // string + 2 for promotion
		end_sq := end[0:2] 

		s := move_to_index(start)
		e := move_to_index(end_sq)

		fmt.Printf("Square: start: %s, end: %s\n", start, end_sq)
		fmt.Printf("Index: start: %d, end: %d\n", s, e)

		//gs.Board.Print()

	}

	if cmd == "quit" {
		return false
	}

	return true

}