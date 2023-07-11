package piece

func (piece *Piece) isSameColor(otherPiece *Piece) bool {
	return piece.Color == otherPiece.Color
}

func (piece *Piece) slideMoves(board [8][8]*Piece, directions []direction) []Position {
	legalMoves := []Position{}
	for _, direction := range directions {
		position := piece.Position
		for true {
			position.X += uint(direction.x)
			position.Y += uint(direction.y)
			if position.X > 7 || position.Y > 7 || position.X < 0 || position.Y < 0 {
				break
			}

			if boardHasPiece(board, position.X, position.Y) {
				if !piece.isSameColor(board[position.Y][position.X]) {
					legalMoves = append(legalMoves, position)
				}
				break
			}
			legalMoves = append(legalMoves, position)
		}
	}
	return legalMoves
}
func (piece *Piece) fixedMoves(board [8][8]*Piece, directions []direction) []Position {
	legalMoves := []Position{}
	for _, direction := range directions {
		position := piece.Position
		position.X += uint(direction.x)
		position.Y += uint(direction.y)
		if position.X > 7 || position.Y > 7 || position.X < 0 || position.Y < 0 {
			continue
		}

		if boardHasPiece(board, position.X, position.Y) {
			if !piece.isSameColor(board[position.Y][position.X]) {
				legalMoves = append(legalMoves, position)
			}
			continue
		}
		legalMoves = append(legalMoves, position)
	}
	return legalMoves
}
