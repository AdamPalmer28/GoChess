package cli_engine

import (
	"chess/chess_engine/make_game"
)

type Config struct {
	gs *make_game.GameState
	cmds []string
}


func MakeConfig(gs *make_game.GameState) *Config {

	cfg := &Config{
		gs: gs,
		cmds: []string{},
	}

	return cfg
}


