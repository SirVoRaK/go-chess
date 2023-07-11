package piece

func (piece *Piece) kingMoves(board [8][8]*Piece) []Position {
	directions := []direction{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	return piece.fixedMoves(board, directions)
}
