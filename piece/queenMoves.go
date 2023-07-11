package piece

func (piece *Piece) queenMoves(board [8][8]*Piece) []Position {
	rookMoves := piece.rookMoves(board)
	bishopMoves := piece.bishopMoves(board)
	return append(rookMoves, bishopMoves...)
}
