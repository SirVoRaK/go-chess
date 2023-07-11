package piece

func (piece *Piece) pawnMoves(board [8][8]*Piece) []Position {
	if piece.Color == WHITE {
		return whitePawnMoves(piece, board)
	}
	return blackPawnMoves(piece, board)
}

func whitePawnMoves(piece *Piece, board [8][8]*Piece) []Position {
	return pawnMoves(piece, board, 6, -1)
}

func blackPawnMoves(piece *Piece, board [8][8]*Piece) []Position {
	return pawnMoves(piece, board, 1, 1)
}

func pawnMoves(piece *Piece, board [8][8]*Piece, initialRow uint, forwardFactor int) []Position {
	legalMoves := []Position{}
	forwardRow := piece.Position.Y + uint(forwardFactor)
	if forwardRow < 0 || forwardRow > 7 {
		return legalMoves
	}

	if !boardHasPiece(board, piece.Position.X, forwardRow) {
		legalMoves = append(legalMoves, Position{piece.Position.X, forwardRow})

		doubleForwardRow := piece.Position.Y + uint(forwardFactor*2)
		if piece.Position.Y == initialRow && !boardHasPiece(board, piece.Position.X, doubleForwardRow) {
			legalMoves = append(legalMoves, Position{piece.Position.X, doubleForwardRow})
		}
	}

	captureLeft := int(piece.Position.X) - 1
	if captureLeft >= 0 {
		captureLeftPiece := board[forwardRow][captureLeft]
		if captureLeftPiece != nil && captureLeftPiece.Color != piece.Color {
			legalMoves = append(legalMoves, Position{uint(captureLeft), forwardRow})
		}
	}

	captureRight := int(piece.Position.X) + 1
	if captureRight <= 7 {
		captureRightPiece := board[forwardRow][captureRight]
		if captureRightPiece != nil && captureRightPiece.Color != piece.Color {
			legalMoves = append(legalMoves, Position{uint(captureRight), forwardRow})
		}
	}

	return legalMoves
}
