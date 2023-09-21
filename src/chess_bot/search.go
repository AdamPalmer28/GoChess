package chess_bot

import (
	"chess/chess_engine/gamestate"
	"fmt"
)

type Search struct { 
	gs *gamestate.GameState 
	best_move uint
	best_eval float64
	
	MaxDepth uint
	QuieDepth uint
	
	TT TranspositionTable
	
	total_nodes uint
	pruned_nodes uint
}

type TranspositionTable struct {

	chesshash uint
	depth uint // depth of the position
	score float64 // score of the position
}

// Search for the best move
func Best_Move(gs *gamestate.GameState, depth uint) {

	// loop through the moves

	cur_search := Search{gs: gs, 
		MaxDepth: depth, 
		QuieDepth: 0, 

		total_nodes: 0,
		pruned_nodes: 0,
	}


	AlphaBeta(&cur_search, -100000, 100000, 0)

	fmt.Printf("Total nodes: %d Pruned nodes: %d\n", cur_search.total_nodes, cur_search.pruned_nodes)

	best_move := cur_search.best_move
	best_score := cur_search.best_eval


	fmt.Printf("Best move: %b Score: %f\n", best_move, best_score)
	gs.Make_move(best_move)
}


// ============================================================================
// AlphaBeta search
// ============================================================================

// AlphaBeta search for the best move
func AlphaBeta(cur_search *Search, alpha float64, beta float64, 
						cur_depth uint) float64{
	var score float64 = 0
	
	if cur_depth == cur_search.MaxDepth {
		// quiescence search
		return Quiescence(cur_search, alpha, beta, 0)
	}
	gs := cur_search.gs

	// game over
	if gs.GameOver {
		if gs.InCheck {
			// checkmate
			score = (10_000 - float64(cur_depth) * 100)

		} else {
			// stalemate
			score = 0
		}
		if score > alpha{
			alpha = score
		}
		return alpha
	}

	for _, move := range gs.MoveList {
		cur_search.total_nodes += 1


		//cb := cur_search.gs.Board.Copy()
		gs.Make_move(move)

		// beta alpha are negated because we are looking at the opponent's
		score = -AlphaBeta(cur_search, -beta, -alpha, cur_depth + 1)

		// if move == 0b011111_111011 {
		// 	println(score, alpha, beta)
		// }
		gs.Undo()

		// if !cb.Identical(gs.Board) {
		// 	fmt.Printf("Move: %b\n", move)
			
		// 	println("Old board: ")
		// 	cb.Print()	
		// 	println("New board: ")
		// 	gs.Board.Print()

		// 	panic("board not equal")
		// }
		if score >= beta {
			// beta cutoff - prune move
			cur_search.pruned_nodes += 1
			return beta
		}
		if score > alpha {
			// new best move
			alpha = score
			if cur_depth == 0 {
				cur_search.best_move = move
				cur_search.best_eval = float64(cur_search.gs.PlayerBoard.Fwd / 8) * score
			}
		}
	}
	return alpha
}


// Quiescence search for the best move
func Quiescence(cur_search *Search, alpha float64, beta float64, cur_quie_depth uint) float64 {
	
	gs := cur_search.gs
	turn_scalar := float64(gs.PlayerBoard.Fwd / 8) 

	// game over
	if gs.GameOver {
		if gs.InCheck {
			// checkmate
			return turn_scalar * 100_000
		} else {
			// stalemate
			return 0
		}
	}

	eval := turn_scalar * (Evaluate(gs))
	//eval := Evaluate(cur_search.gs)
	// assumes not in zugzwang
	if eval >= beta {
		cur_search.pruned_nodes += 1
		return beta
	}
	if eval > alpha {
		alpha = eval
	}

	if cur_quie_depth == cur_search.QuieDepth {
		return eval
	}


	important_moves := gs.MoveList.ImportantMoves()
	for _, move := range important_moves {
		cur_search.total_nodes += 1


		gs.Make_move(move)
		
		score := -Quiescence(cur_search, -beta, -alpha, cur_quie_depth + 1)
		
		gs.Undo()
		
		if score >= beta {
			// beta cutoff - prune move 
			cur_search.pruned_nodes += 1
			return beta
		}
		if score > alpha {
			// new best move
			alpha = score
		}
	}
	return alpha
}
