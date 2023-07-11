package piece

func (piece *Piece) bishopMoves(board [8][8]*Piece) []Position {
	directions := []direction{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	return piece.slideMoves(board, directions)
}
