package piece

func (piece *Piece) rookMoves(board [8][8]*Piece) []Position {
	directions := []direction{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	return piece.slideMoves(board, directions)
}
