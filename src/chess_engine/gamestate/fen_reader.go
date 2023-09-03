package gamestate

import (
	"chess/chess_engine/board"
	"strconv"
	"strings"
	"unicode"
)

/* FEN format (seperated by spaces):
1: piece placement data (starting with rank 8 and ending with rank 1)
2: active color (i.e. whose turn it is)
3: castling availability - max: KQkq
4: en passant target square
5: halfmove clock: number of halfmoves since last capture or pawn advance
6: move number
*/

func FEN_to_gs(fen string) *GameState {

	// split fen into sections
	fen_sections := strings.Split(fen, " ")

	piece_data := fen_sections[0]
	active_color := fen_sections[1]
	castling_availability := fen_sections[2]
	en_passant_sq := fen_sections[3]
	halfmove_clock, _ := strconv.Atoi(fen_sections[4])
	move_number, _ := strconv.Atoi(fen_sections[5])

	// board data
	board := FENdata_to_board(piece_data)

	// castle rights
	var white_castle uint
	var black_castle uint

	if strings.Contains(castling_availability, "Q") {
		white_castle |= 0b01
	}
	if strings.Contains(castling_availability, "K") {
		white_castle |= 0b10
	}
	if strings.Contains(castling_availability, "q") {
		black_castle |= 0b01
	}
	if strings.Contains(castling_availability, "k") {
		black_castle |= 0b10
	}

	// en passant target
	var en_passant_squ_ind uint = 64
	if en_passant_sq != "-" {
		en_passant_squ_ind = uint(en_passant_sq[0] - 'a') + 
								uint(en_passant_sq[1] - '1')*8
	}

	// create gamestate
	gs := &GameState{
		Board: *board,
		WhiteCastle: white_castle,
		BlackCastle: black_castle,
		White_to_move: active_color == "w",

		Enpass_ind: en_passant_squ_ind,

		MoveList: []uint{},
		Moveno: uint(move_number),
		HalfMoveNo: uint(halfmove_clock),
	}

	return gs
}

func FENdata_to_board(fen_data string) *board.ChessBoard{

	// split board section into rows
	rows := strings.Split(fen_data, "/")

	// WP, WN, WB, WR, WQ, WK, BP, BN, BB, BR, BQ, BK
	var bbs [12]board.Bitboard
	
	for i := 0; i < 12; i++ {
		bbs[i] = 0
	}


	var ind int = 64
	// loop through rows
	for _, row := range rows {
		ind -= 8
		// loop through row 
		for _, char := range row {

			// check if char is a number
			if unicode.IsNumber(char) {

				// convert char to int
				num, _ := strconv.Atoi(string(char))

				ind += num


			} else {

				// convert row and col to index
				piece := string(char)

				bb_ind := PieceBBind[piece]
				
				bbs[bb_ind] |= board.Bitboard(1) << uint(ind)
				ind ++
				
			}
		}
		ind -= 8
	}

	// create board
	board := &board.ChessBoard{
		WhitePawns: &bbs[0],
		WhiteKnights: &bbs[1],
		WhiteBishops: &bbs[2],
		WhiteRooks: &bbs[3],
		WhiteQueens: &bbs[4],
		WhiteKing: &bbs[5],
		BlackPawns: &bbs[6],
		BlackKnights: &bbs[7],
		BlackBishops: &bbs[8],
		BlackRooks: &bbs[9],
		BlackQueens: &bbs[10],
		BlackKing: &bbs[11],
	}

	return board
}
