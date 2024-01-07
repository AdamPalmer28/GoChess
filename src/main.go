// package main

// import (
// 	"chess/src/chess_bot"
// 	"chess/src/chess_engine"
// 	"chess/src/chess_engine/gamestate"
// 	"chess/src/chess_engine/move_gen/magic"
// 	"chess/src/cli_engine"
// 	"fmt"
// )

// func main() {

// 	//time to generate attack rays

// 	if false {
// 		println("Generating diag magics... \n")
// 		magic.Gen_all_magics(true) // generate diagonal magics
// 		println("\n\nGenerating straight magics... \n")
// 		magic.Gen_all_magics(false) // generate straight magics
// 	}

// 	//fen := "2qk4/8/8/8/8/6K1/8/3q4 b - - 0 1" // black mate in 2
// 	//fen := "4k3/8/Q7/8/8/8/1R6/4K3 w - - 0 1" // white mate in 2
// 	//gs := chess_engine.CreateGameFen(fen)

// 	gamestate.InitZobrist()

// 	// start the game
// 	gs := chess_engine.StartGame()
	
// 	gs.Board.Print()
	

// 	cli := cli_engine.MakeConfig(gs)
// 	for {
		
// 		if !gs.White_to_move {
// 			// AI move
// 			println("\nAI move\n")

// 			chess_bot.FindBestMove(gs, 7, true)	
// 			gs.Board.Print()
// 		}
// 		score := chess_bot.Evaluate(gs)
// 		fmt.Printf("\nCurrent Eval: %.3f", score)
// 		println("\n")

// 		result := cli.Run()

// 		// condition to break the loop
// 		if !result {
// 			break
// 		}

// 	}
// }