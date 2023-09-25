package chess_bot

import (
	"chess/chess_engine/gamestate"
	"chess/cli_engine"
	"fmt"
	"time"
)

type Search struct { 
	gs *gamestate.GameState 
	best_move uint
	best_eval float64
	
	MaxDepth uint
	QuieDepth uint
	
	TT map[uint64]TT
	
	total_nodes uint
	pruned_nodes uint
	TT_hits uint
}

type TT struct {
	depth uint // depth of the position
	score float64 // score of the position
}

// Search for the best move
func Best_Move(gs *gamestate.GameState, depth uint) {

	// loop through the moves

	cur_search := Search{gs: gs, 
		MaxDepth: depth, 
		QuieDepth: 4, 
		total_nodes: 0,
		pruned_nodes: 0,
		TT_hits: 0,
		TT: make(map[uint64]TT),
	}
	start := time.Now()

	AlphaBeta(&cur_search, -100000, 100000, 0)

	elapsed := time.Since(start)
	nodes_per_sec := float64(cur_search.total_nodes) / elapsed.Seconds()

	fmt.Printf("\nTime elapsed: %.2f - Nodes per second: %.0f\n", elapsed.Seconds(), nodes_per_sec)
	
	fmt.Printf("Nodes: %d Pruned: %d (TT: hits %d)\n", cur_search.total_nodes, cur_search.pruned_nodes, cur_search.TT_hits)
	fmt.Printf("Depth: %d, QuieDepth %d\n", cur_search.MaxDepth, cur_search.QuieDepth)

	best_move := cur_search.best_move
	best_score := cur_search.best_eval


	fmt.Printf("\nBest move:\n")
	cli_engine.GetMoves([]uint{best_move})
	fmt.Printf("Score: %f\n", best_score)
	gs.Make_move(best_move)
}


// ============================================================================
// AlphaBeta search
// ============================================================================

// AlphaBeta search for the best move
func AlphaBeta(cur_search *Search, alpha float64, beta float64, 
						cur_depth uint) float64{
	var score float64 = 0

	// check transposition table
	gs_hash := cur_search.gs.Hash
	tt_entry, ok := cur_search.TT[gs_hash]; 

	if (tt_entry.depth >= cur_depth) && ok {
		// TT hit
		cur_search.TT_hits += 1
		return tt_entry.score
	}
	
	if cur_depth == cur_search.MaxDepth { // max depth reached
		// quiescence search
		return Quiescence(cur_search, alpha, beta, 0)
	}
	gs := cur_search.gs

	// game over ----------------------------------------
	if gs.GameOver {
		if gs.InCheck {// checkmate
			return float64(gs.PlayerBoard.Fwd / 8) * (10_000 - float64(cur_depth-1) * 100)
		} else {// stalemate
			return 0
		}
	}

	// search moves --------------------------------------
	for _, move := range gs.MoveList {
		cur_search.total_nodes += 1

		gs.Make_move(move)

		// beta alpha are negated because we are looking at the opponent's
		score = -AlphaBeta(cur_search, -beta, -alpha, cur_depth + 1)

		// add to transposition table
		cur_search.TT[gs_hash] = TT{depth: cur_depth, score: score}

		gs.Undo()

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
	var score float64 = 0

	// check transposition table
	gs_hash := gs.Hash
	tt_entry, ok := cur_search.TT[gs_hash]
	if (tt_entry.depth >= cur_search.MaxDepth + cur_quie_depth) && ok {
		// TT hit
		cur_search.TT_hits += 1
		return tt_entry.score
	}

	// game over
	if gs.GameOver {
		if gs.InCheck {// checkmate
			depth := cur_search.MaxDepth + cur_quie_depth
			return turn_scalar * (10_000 - float64(depth-1) * 100)
		} else {// stalemate
			return 0
		}
	}

	eval := turn_scalar * (Evaluate(gs))

	// assumes not in zugzwang
	if eval >= beta {
		cur_search.pruned_nodes += 1
		return beta
	}
	if eval > alpha {
		alpha = eval
	}

	if (cur_quie_depth >= cur_search.QuieDepth) && (!gs.InCheck) {
		// if not in check and quiescence depth reached
		return eval
	}


	important_moves := gs.MoveList.ImportantMoves()
	for _, move := range important_moves {
		cur_search.total_nodes += 1


		gs.Make_move(move)
		
		score = -Quiescence(cur_search, -beta, -alpha, cur_quie_depth + 1)

		// add to transposition table
		cur_search.TT[gs_hash] = TT{depth: cur_search.MaxDepth + cur_quie_depth, score: score}

		
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
