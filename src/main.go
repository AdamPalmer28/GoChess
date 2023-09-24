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


	fen := "8/1p1b2pk/5p1p/p1p5/2B2P1K/2P4P/PP2R1P1/3r4 b - - 0 1"
	gs := chess_engine.CreateGameFen(fen)

	// start the game
	//gs := chess_engine.StartGame()
	
	gs.Board.Print()

	print(gs.Enpass_ind)
	

	cli := cli_engine.MakeConfig(gs)
	for {
		if !gs.White_to_move {
			// AI move
			println("AI move")

			chess_bot.Best_Move(gs, 6)	
			gs.Board.Print()
		}
		// score := chess_bot.Evaluate(gs)
		// println("Score: ", score)

		result := cli.Run()

		// condition to break the loop
		if !result {
			break
		}

	}
}