package main

import (
	"chess/chess_engine"
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
	type bitboard uint64

	var b bitboard = 01
	b = b << 1
	fmt.Println(b)
	b = b << 5
	fmt.Println(b)
}