package board

import (
	"piece"
)

type Castling struct {
    WhiteKingSide bool
    WhiteQueenSide bool
    BlackKingSide bool
    BlackQueenSide bool
}
type Board struct {
    Pieces []piece.Piece
    turn piece.Color
    Castling Castling
}
func (b *Board) fillDefaults() {
    b.Castling = Castling{true, true, true, true}
}
func (b *Board) GetTurn() piece.Color {
    return b.turn
}
func (b *Board) ChangeTurn() {
    b.turn = b.turn.Opposite()
}
func (b *Board) PieceAt(x uint, y uint) (*piece.Piece, int) {
    for i, p := range b.Pieces {
        if p.Position.X == x && p.Position.Y == y {
            return &b.Pieces[i], i
        }
    }
    return nil, -1
}
func (b *Board) MovePiece(from piece.Position, to piece.Position) {
    if from == to {
        panic("Can't move to the same position")
    }

    fromPiece, _ := b.PieceAt(from.X, from.Y)
    if fromPiece == nil {
        panic("No piece at position")
    }

    if fromPiece.Color != b.turn {
        panic("Not your turn")
    }

    toPiece, toIndex := b.PieceAt(to.X, to.Y)
    if toPiece != nil {
        if toPiece.Color == b.turn {
            panic("Can't capture your own piece")
        }

        b.removePiece(uint(toIndex))
    }


    fromPiece.Move(to)
}
func (b *Board) removePiece(i uint) {
    b.Pieces[i] = b.Pieces[len(b.Pieces)-1]
    b.Pieces = b.Pieces[:len(b.Pieces)-1]
}

func InitialPositionBoard() Board {
    initialFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
    return BoardFromFen(initialFen)
}
