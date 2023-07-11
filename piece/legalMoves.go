package piece

import (
	"fmt"
)

type direction struct {
	x int
	y int
}

func (piece *Piece) LegalMoves(board [8][8]*Piece) []Position {
	legalMoves := []Position{}
	if piece.PieceType == PAWN {
		legalMoves = piece.pawnMoves(board)
	} else if piece.PieceType == BISHOP {
		legalMoves = piece.bishopMoves(board)
	} else if piece.PieceType == ROOK {
		legalMoves = piece.rookMoves(board)
	} else if piece.PieceType == QUEEN {
		legalMoves = piece.queenMoves(board)
	} else if piece.PieceType == KNIGHT {
		legalMoves = piece.knightMoves(board)
	} else if piece.PieceType == KING {
		legalMoves = piece.kingMoves(board)
	} else {
		panic("Unknown piece type")
	}

	printPositions(legalMoves)

	return legalMoves
}

func printPositions(positions []Position) {
	for _, position := range positions {
		fmt.Printf("%s ", IndexToNotation(position.X, position.Y))
	}
	fmt.Println()
}

func boardHasPiece(board [8][8]*Piece, x uint, y uint) bool {
	return board[y][x] != nil
}

func IndexToNotation(x uint, y uint) string {
	if x > 7 || y > 7 {
		return ""
	}
	letters := "abcdefgh"
	return string(letters[x]) + fmt.Sprint(8-y)
}
