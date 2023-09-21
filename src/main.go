package main

import (
	"chess/chess_bot"
	"chess/chess_engine"
	"chess/chess_engine/move_gen/magic"
	"chess/cli_engine"
)

func main() {

	//time to generate attack rays

	if false {
		println("Generating diag magics... \n")
		magic.Gen_all_magics(true) // generate diagonal magics
		println("\n\nGenerating straight magics... \n")
		magic.Gen_all_magics(false) // generate straight magics
	}

	//fen := "4k1r1/8/1q6/3NQ3/N6R/2B5/3PPK2/8 w - - 0 1"
	//gs := chess_engine.CreateGameFen(fen)


	// start the game
	gs := chess_engine.StartGame()
	
	gs.Board.Print()
	

	cli := cli_engine.MakeConfig(gs)
	for {
		if !gs.White_to_move {
			// AI move
			println("AI move")

			chess_bot.Best_Move(gs, 4)	
			gs.Board.Print()
		}

		result := cli.Run()

		// condition to break the loop
		if !result {
			break
		}

	}
}