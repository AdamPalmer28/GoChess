package main

import (
	"chess/chess_engine"
	"chess/chess_engine/move_gen"
	"chess/cli_engine"
)

func main() {

	// time to generate attack rays
	if true {

		all_occ := move_gen.AllOccupancy(13, true)
		inner_bb := move_gen.InnerOccupancy(13)
		inner_bb.Print()

		println("All occupancy: ", len(all_occ))

		for i, occ := range all_occ {
			println("Index: ", i)
			occ.Print()
			if i == 100 {
				break
			}
		}
		return
	}


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