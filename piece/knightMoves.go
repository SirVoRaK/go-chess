package piece

func (piece *Piece) knightMoves(board [8][8]*Piece) []Position {
	directions := []direction{
		{1, 2},
		{2, 1},
		{2, -1},
		{1, -2},
		{-1, -2},
		{-2, -1},
		{-2, 1},
		{-1, 2},
	}
	return piece.fixedMoves(board, directions)
}
