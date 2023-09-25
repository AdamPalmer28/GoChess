package board

// update piece locations for a given piece type
func (cb *ChessBoard) UpdatePieceLocations (i uint) {

	switch i {
		case 0: // white pawn
			cb.PieceLocations[0] = cb.WhitePawns.Index()
		case 1: // white knight
			cb.PieceLocations[1] = cb.WhiteKnights.Index()
		case 2: // white bishop
			cb.PieceLocations[2] = cb.WhiteBishops.Index()
		case 3: // white rook
			cb.PieceLocations[3] = cb.WhiteRooks.Index()
		case 4: // white queen
			cb.PieceLocations[4] = cb.WhiteQueens.Index()
		case 5: // white king
			cb.PieceLocations[5] = cb.WhiteKing.Index()
		case 6: // black pawn
			cb.PieceLocations[6] = cb.BlackPawns.Index()
		case 7: // black knight
			cb.PieceLocations[7] = cb.BlackKnights.Index()
		case 8: // black bishop
			cb.PieceLocations[8] = cb.BlackBishops.Index()
		case 9: // black rook
			cb.PieceLocations[9] = cb.BlackRooks.Index()
		case 10: // black queen
			cb.PieceLocations[10] = cb.BlackQueens.Index()
		case 11: // black king
			cb.PieceLocations[11] = cb.BlackKing.Index()	
	}
}
