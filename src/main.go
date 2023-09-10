package main

import (
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

	fen := "8/8/8/8/8/8/8/R3K3 w Q - 0 1"
	gs := chess_engine.CreateGameFen(fen)


	// start the game
	//gs := chess_engine.StartGame()

	gs.Board.Print()

	cli := cli_engine.MakeConfig(gs)

	for {
		result := cli.Run()

		// condition to break the loop
		if !result {
			break
		}

	}
}