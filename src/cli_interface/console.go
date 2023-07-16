package chess_cli

import (
	"chess/chess_engine/make_game"
	"fmt"
)

func chess_console(gs *make_game.GameState) {
	
	var user_cmd string

	fmt.Scan(&user_cmd)

	fmt.Println("Your input is:", user_cmd)
}