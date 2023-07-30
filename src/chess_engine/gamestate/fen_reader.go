package gamestate

import (
	"ChessEngine/chess_engine/board"
	"strings"
	"unicode"
	"strconv"
)

Piece_to_index := func(piece rune) int {



/* FEN format (seperated by spaces)
1: piece placement data (starting with rank 8 and ending with rank 1)
2: active color (i.e. whose turn it is)
3: castling availability - max: KQkq
4: en passant target square
5: halfmove clock: number of halfmoves since last capture or pawn advance
6: move number
*/

func FEN_to_gs(fen string) GameState {



	// split fen into sections
	fen_sections := strings.Split(fen, " ")

	piece_data := fen_sections[0]





}

func FENdata_to_board(fen_data string) *board.ChessBoard{

	// split board section into rows
	rows := strings.Split(fen_data, "/")


	// WP, WN, WB, WR, WQ, WK, BP, BN, BB, BR, BQ, BK
	var bbs [12]board.Bitboard
	
	for i := 0; i < 12; i++ {
		bbs[i] = 0
	}

	// loop through rows
	for row_ind, row := range rows {

		// loop through row
		for _, char := range row {

			// check if char is a number
			if unicode.IsNumber(char) {

				// convert char to int
				spaces, _ := strconv.Atoi(string(char))

				// add spaces to row
				for i := 0; i < spaces; i++ {

					// convert row and col to index
					ind := row_ind*8 + i

					// add empty square to board
					bbs[6] |= 1 << ind
				}

			} else {

				// convert row and col to index
				ind := row_ind*8 + int(char-'a')

				// add piece to board
				bbs[Piece_to_index(char)] |= 1 << ind
			}
		}
	}
	

}
