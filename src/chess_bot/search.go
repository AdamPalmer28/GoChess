package chess_bot

import (
	"chess/chess_engine/gamestate"
	"fmt"
)

type Search struct { //! modify search to work with this struct 
	gs gamestate.GameState 
	best_move uint
	best_eval float64

	Depth uint
	Quie_depth uint

	Transposition_table map[uint]float64 //! change to its own struct

	total_nodes uint
	pruned_nodes uint
}


// Search for the best move
func Best_Move(gs *gamestate.GameState, depth uint) {

	// loop through the moves

	var best_move uint = 0
	var best_score float64 
	if gs.White_to_move {
		best_score = -100000
	} else {
		best_score = 100000
	}
	// start alpha beta search
	// alpha = -100000 as this is the worst possible score
	// beta = 100000 as this is the best possible score
	for _, move := range gs.MoveList {

		gs.Make_move(move)

		score := AlphaBeta(gs, -100000, 100000, depth - 1)

		gs.Undo()

		//fmt.Printf("Move: %b Score: %f\n", move, score)

		if gs.White_to_move {
			if score > best_score {
				best_score = score
				best_move = move
			}
		} else {
			if score < best_score {
				best_score = score
				best_move = move
			}
		}
	}

	fmt.Printf("Best move: %b Score: %f\n", best_move, best_score)
	gs.Make_move(best_move)
}


var quie_depth uint = 4

// AlphaBeta search for the best move
func AlphaBeta(gs *gamestate.GameState, alpha float64, beta float64, 
						depth_left uint) float64{

	if depth_left == 0 {
		// quiescence search
		return Quiescence(gs, alpha, beta, quie_depth)
	}

	var score float64 = 0
	for _, move := range gs.MoveList {
		gs.Make_move(move)

		// beta alpha are negated because we are looking at the opponent's
		score = -AlphaBeta(gs, -beta, -alpha, depth_left - 1)

		gs.Undo()

		if score >= beta {
			// beta cutoff - prune move
			return beta
		}
		if score > alpha {
			// new best move
			alpha = score
		}
	}
	return score
}

// Quiescence search for the best move
func Quiescence(gs *gamestate.GameState, alpha float64, beta float64, quie_depth uint) float64 {
	
	eval := float64(gs.PlayerBoard.Fwd / 8) * (Evaluate(gs))

	// assumes not in zugzwang
	if eval >= beta {
		return beta
	}
	if eval > alpha {
		alpha = eval
	}

	if quie_depth == 0 {
		// exit
		return eval
	}


	important_moves := gs.MoveList.ImportantMoves()
	for _, move := range important_moves {
		gs.Make_move(move)
		score := -Quiescence(gs, -beta, -alpha, quie_depth - 1)

		gs.Undo()

		if score >= beta {
			// beta cutoff - prune move 
			return beta
		}
		if score > alpha {
			// new best move
			alpha = score
		}
	}
	return alpha
}
