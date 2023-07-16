package main

import (
	"chess/chess_engine"
	"fmt"
)

func main() {

	println("Hello World")

	board := chess_engine.StartGame()

	board.Print()

	while true; {
		fmt.println("new cmd")
		chess_console(board)

	}
}
