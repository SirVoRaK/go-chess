package piece

import (
	"strings"
)

type PieceType string
type Color string

const (
	PAWN   PieceType = "pawn"
	KNIGHT PieceType = "knight"
	BISHOP PieceType = "bishop"
	ROOK   PieceType = "rook"
	QUEEN  PieceType = "queen"
	KING   PieceType = "king"
)
const (
	WHITE Color = "white"
	BLACK Color = "black"
)

func (c *Color) Opposite() Color {
	if *c == WHITE {
		return BLACK
	}
	return WHITE
}

type Position struct {
	X uint
	Y uint
}
type Piece struct {
	PieceType PieceType
	Color     Color
	Position  Position
}

func (p *Piece) Move(position Position) {
	if position.X > 7 || position.Y > 7 {
		panic("Invalid position")
	}
	p.Position.X = position.X
	p.Position.Y = position.Y
}
func (p *Piece) Letter() string {
	letter := ""
	if p.PieceType == KNIGHT {
		letter = "n"
	} else {
		letter = string(p.PieceType[0])
	}
	if p.Color == WHITE {
		return strings.ToUpper(letter)
	}
	return letter
}

var lettersToPieces = map[string]PieceType{
	"p": PAWN, "n": KNIGHT, "b": BISHOP, "r": ROOK, "q": QUEEN, "k": KING,
}

func LetterToPiece(letter string, position Position) Piece {
	lowerLetter := strings.ToLower(letter)
	pieceType, ok := lettersToPieces[lowerLetter]
	if !ok {
		panic("Invalid piece letter")
	}
	color := WHITE
	if letter == lowerLetter {
		color = BLACK
	}
	return Piece{pieceType, color, position}
}

func LetterToColumn(letter string) uint {
	letters := "abcdefgh"
	for i := 0; i < len(letters); i++ {
		if letter == string(letters[i]) {
			return uint(i)
		}
	}
	panic("Invalid letter")
}
func NumberToRow(number uint) uint {
	return 8 - number
}
