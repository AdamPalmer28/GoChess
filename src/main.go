package main

import (
	"chess/chess_engine"
	"chess/cli_engine"
)

func main() {

	//time to generate attack rays



	// if true {
	// 	magic.Gen_all_magics(true)
	// 	magic.Gen_all_magics(false)
	// }



	// start the game
	gs := chess_engine.StartGame()

	gs.Board.Print()

	cli := cli_engine.MakeConfig(gs)

	for {
		result := cli.Run()

		// condition to break the loop
		if result == false {
			break
		}

	}
}