package chess_bot

import (
	"chess/src/chess_engine/gamestate"
	"chess/src/cli_engine"
	"fmt"
	"math"
	"strconv"
)

// Search struct
type Search struct { 

	gs *gamestate.GameState 

	// parameters
	MaxDepth uint
	QuieDepth uint

	// results data
	best_move uint
	best_eval float64

	MoveTree []MoveScoreTree
	
	// Internal data
	TT map[uint64]TT
	
	// stats
	total_nodes uint
	pruned_nodes uint
	TT_hits uint
	TT_success uint
}

// move eval lists
type MoveScoreTree struct {
		MoveNo uint
		Score float64
		NextMove *[]MoveScoreTree
	}

// ChessBot Main function
func FindBestMove(gs *gamestate.GameState, depth uint, make_move bool) Search {

	// loop through the moves

	cur_search := Search{
		gs: gs, 
		MaxDepth: depth, 
		QuieDepth: 4, 
		TT: make(map[uint64]TT),
	}
	//start := time.Now()

	MoveEvalTree := []MoveScoreTree{}

	AlphaBeta(&cur_search, -100000, 100000, 0, &MoveEvalTree)

	//elapsed := time.Since(start)

	// print results
	//cur_search.Print(elapsed.Seconds())


	if make_move {
		// make the best move
		best_move := cur_search.best_move
		gs.Make_move(best_move)
	}

	return cur_search
}





// ============================================================================
// Display search results
// ============================================================================


// print stats of the search
func (cur_search *Search) Print(time float64) {

	fmt.Printf("\nSearch stats\n===============================================\n")

	// best move
	best_move := cur_search.best_move
	best_score := cur_search.best_eval
	displayScore(best_score)
	cli_engine.GetMoves([]uint{best_move})

	// time
	computed_nodes := cur_search.total_nodes - cur_search.TT_success
	nodes_per_sec := float64(cur_search.total_nodes) / time
	computed_nodes_per_sec := float64(computed_nodes) / time

	fmt.Printf("\nTime elapsed: %.1f", time)
	fmt.Printf("\nNPS: %.0f\nCNPS: %.0f)\n", nodes_per_sec, computed_nodes_per_sec)

	// depth
	fmt.Printf("\nDepth\n--------------------------------------\n")
	fmt.Printf("Depth: %d\n", cur_search.MaxDepth)
	fmt.Printf("QuieDepth: %d\n", cur_search.QuieDepth)

	// nodes
	fmt.Printf("\nNodes\n--------------------------------------\n")
	fmt.Printf("Nodes: %s\n", formatWithCommas(cur_search.total_nodes))
	fmt.Printf("Pruned: %s\n", formatWithCommas(cur_search.pruned_nodes))
	fmt.Printf("\nTT:\n")
	fmt.Printf("   successes: %s\n", formatWithCommas(cur_search.TT_success))
	fmt.Printf("   hits: %s\n", formatWithCommas(cur_search.TT_hits))

	// create gap
	fmt.Printf("\n\n\n")
}

// display evaluation score (+ mate in x)
func displayScore(score float64) {
	
	if (score > 5000) || (score < -5000) {
		// mate in x
		diff := 10_000 - math.Abs(score) 

		depth := int(diff / 100) 
		mate_in := (depth + 1) / 2 

		if score < 0 { 
			// black has mate in x
			fmt.Printf("\nMate in %d\n", -mate_in)
		} else {
			// white has mate in x
			fmt.Printf("\nMate in %d\n", mate_in)
		}

	} else {
		fmt.Printf("\nScore: %.3f\n", score)
	}
}

func formatWithCommas(number uint) string {
    // Convert the integer to a string
    str := strconv.Itoa(int(number))

    // Determine the length of the string
    length := len(str)

    // Calculate the number of commas needed
    numCommas := (length - 1) / 3

    // Create a buffer to hold the formatted string
    formatted := make([]byte, length+numCommas)

    // Loop through the original string and insert commas
    j := 0
    for i := 0; i < length; i++ {
        if i > 0 && (length-i)%3 == 0 {
            formatted[j] = ','
            j++
        }
        formatted[j] = str[i]
        j++
    }

    // Convert the byte slice back to a string
    return string(formatted)
}