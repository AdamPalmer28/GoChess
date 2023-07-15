package main

import (
	"chess/chess_engine"
	"fmt"
)

func main() {

	println("Hello World")

	board := chess_engine.StartGame()

	board.Print()

	var input string

	fmt.Print("Type your input: ")
	fmt.Scan(&input)
	fmt.Println("Your input is:", input)
}
