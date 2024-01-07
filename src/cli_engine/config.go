package cli_engine

import (
	gamestate "chess/src/chess_engine/gamestate"
)

type Config struct {
	gs *gamestate.GameState
	cmds []string
}


func MakeConfig(gs *gamestate.GameState) *Config {

	cfg := &Config{
		gs: gs,
		cmds: []string{},
	}

	return cfg
}


