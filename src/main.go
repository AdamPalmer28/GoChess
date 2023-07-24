package main

import (
	"chess/chess_engine"
	"chess/cli_engine"
	"fmt"
)

func main() {

	println("Hello World")

	gs := chess_engine.StartGame()

	gs.Board.Print()

	cli := cli_engine.MakeConfig(gs)

	for {

		fmt.Println("new cmd")

		result := cli.Run()
		

		// condition to break the loop
		if result == false {
			break
		}

	}
}