package chess_bot

// Transposition table entry
type TT struct {
	depth uint // depth of the position
	score float64 // score of the position
}

// modify cur_search, and return the score

// 0 - no TT hit
// 1 - TT hit, but depth is too shallow
// 2 - TT hit, and successful

func checkTT(cur_search *Search, depth uint) (uint, float64) {

	hash := cur_search.gs.Hash

	tt_entry, ok := cur_search.TT[hash]
	if ok {
		// TT hit
		cur_search.TT_hits += 1
		
		if (tt_entry.depth <= depth)  {
			// TT successful
			cur_search.TT_success += 1
			return 2, tt_entry.score
		}
		// TT hit, but depth is too shallow
		return 1, tt_entry.score
	}
	return 0, 0
}