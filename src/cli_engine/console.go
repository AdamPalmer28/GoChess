package cli_engine

import (
	"bufio"
	"chess/cli_engine/cli_debug"
	"fmt"
	"os"
	"strings"
)


func (cfg *Config) Run() bool {
	
	scanner := bufio.NewScanner(os.Stdin)
	var user_cmd string

    // Prompt the user for input
	fmt.Printf("\nMove: %x ", cfg.gs.Moveno)
	if cfg.gs.White_to_move {
		fmt.Print("(White to move)")
	} else {
		fmt.Print("(Black to move)")
	}

    fmt.Println("\nEnter a command: ")


    // Read the next line from standard input
    if scanner.Scan() {
		// Get the text that the user entered
		user_cmd = scanner.Text()

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
	} else if cmd == "quit" {
		return false

	} else if cmd == "bb" {
		cli_debug.Bitboard_cli(*gs, inputs)

	// } else if cmd == "fen" {

	// 	fen_str := strings.Join(inputs[1:7], " ")


		
	} else {
		// assume move

		uci_move := cmd
		result := cfg.move_input(uci_move)

		if result == false {
			fmt.Println("Invalid move")
		}
		
	}


	return true

}