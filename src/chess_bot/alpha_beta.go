package chess_bot

import "chess/src/chess_engine/gamestate"

// ============================================================================
// AlphaBeta search
// ============================================================================

// AlphaBeta search for the best move
func AlphaBeta(cur_search *Search, alpha float64, beta float64, cur_depth uint) float64 {
	var score float64 = 0
	gs := cur_search.gs

	if gs.GameOver {
		// game over
		score = game_over_score(gs, cur_depth)

		if score >= beta {
			cur_search.pruned_nodes += 1
			return beta
		}
		if score > alpha { // new best move
			alpha = score
		}
	}


	if cur_depth == cur_search.MaxDepth { // max depth reached
		// quiescence search
		return Quiescence(cur_search, alpha, beta, 0)
	}

	// search moves --------------------------------------
	for _, move := range gs.MoveList {
		cur_search.total_nodes += 1

		gs.Make_move(move)

		found, tt_score := checkTT(cur_search, cur_depth)

		if found == 2 { // TT hit
			score = tt_score
		} else { // no TT hit
			// beta alpha are negated because we are looking at the opponent's
			score = -AlphaBeta(cur_search, -beta, -alpha, cur_depth+1)
			//cur_search.TT[gs.Hash] = TT{depth: cur_depth, score: score}
		}

		gs.Undo() // undo move before exiting loop

		// --------------------------------------
		// alpha beta pruning
		if score >= beta { // beta cutoff - prune move
			cur_search.pruned_nodes += 1
			return beta
		}
		if score > alpha { // new best move
			alpha = score
			if cur_depth == 0 {
				cur_search.best_move = move
				cur_search.best_eval = float64(cur_search.gs.PlayerBoard.Fwd/8) * score
			}
		}
	}
	return alpha
}

// Quiescence search for the best move
func Quiescence(cur_search *Search, alpha float64, beta float64, cur_quie_depth uint) float64 {

	gs := cur_search.gs
	turn_scalar := float64(gs.PlayerBoard.Fwd/8)
	var score float64 = 0

	// get eval of position
	if gs.GameOver {
		// game over
		score = game_over_score(gs, cur_search.MaxDepth + cur_quie_depth)
	} else {
		// evaluate the board
		score = turn_scalar * (Evaluate(gs))
	}

	if score >= beta {
		cur_search.pruned_nodes += 1
		return beta
	}
	if score > alpha { // new best move
		alpha = score
	}


	if (cur_quie_depth >= cur_search.QuieDepth) && (!gs.InCheck) {
		// if not in check and quiescence depth reached
		return alpha
	}

	// search moves --------------------------------------
	important_moves := gs.ScoreMoveList.ImportantMoves()
	for _, move := range important_moves {
		cur_search.total_nodes += 1

		gs.Make_move(move)

		found, tt_score := checkTT(cur_search, cur_search.MaxDepth+cur_quie_depth)

		if found == 2 { // TT hit
			score = tt_score
		} else { // no TT hit
			score = -Quiescence(cur_search, -beta, -alpha, cur_quie_depth+1)
			//cur_search.TT[gs.Hash] = TT{depth: cur_search.MaxDepth + cur_quie_depth, score: score}
		}

		gs.Undo() // undo move before exiting loop


		// --------------------------------------
		// alpha beta pruning
		if score >= beta { // beta cutoff - prune move 
			cur_search.pruned_nodes += 1
			return beta
		}
		if score > alpha { // new best move
			alpha = score
		}
	}
	return alpha
}


func game_over_score(gs *gamestate.GameState, depth uint) float64 {

	var score float64 = 0
	turn_scalar := float64(gs.PlayerBoard.Fwd/8)
	if gs.InCheck { // checkmate
		score = turn_scalar * (10_000 - float64(depth-1)*100)
	} else { // stalemate
		score = 0
	}
	return score
}