package main

import (
	"chess/chess_engine"
	//"chess/chess_cli"
	"fmt"
)

func main() {

	println("Hello World")

	board := chess_engine.StartGame()

	board.Print()

	for {

		fmt.Println("new cmd")
		//chess_cli.chess_console(board)
		

		// condition to break the loop
		if true {
			break
		}
	}
}