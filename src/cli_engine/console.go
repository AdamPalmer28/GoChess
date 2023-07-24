package cli_engine

import (
	"fmt"
	"strings"
)

func (cfg *Config) Run() bool {
	
	var user_cmd string

	fmt.Scan(&user_cmd)

	fmt.Println("Your input is:", user_cmd)

	inputs := strings.Split(user_cmd, " ")
	fmt.Println("Inputs", inputs)
	cmd := inputs[0]

	gs := cfg.gs

	if cmd == "print" {

	cfg.gs.Board.Print()
	}
 
	// move a1a2
	if cmd == "move" {
		
		move := inputs[0]

		start := move[0:2]
		end := move[2:] // string + 2 for promotion
		end_sq := end[0:2] 

		s := move_to_index(start)
		e := move_to_index(end_sq)

		fmt.Printf("Square: start: %s, end: %s\n", start, end_sq)
		fmt.Printf("Index: start: %d, end: %d\n", s, e)

		gs.Board.Print()

	}

	return true

}