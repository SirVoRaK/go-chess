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
    Turn piece.Color
    Castling Castling
}
func (b *Board) fillDefaults() {
    b.Castling = Castling{true, true, true, true}
}
func (b *Board) PieceAt(x uint, y uint) *piece.Piece {
    for _, p := range b.Pieces {
        if p.Position.X == x && p.Position.Y == y {
            return &p
        }
    }
    return nil
}
func (b *Board) MovePiece(from piece.Position, to piece.Position) {
    for i, p := range b.Pieces {
        if p.Position.X == from.X && p.Position.Y == from.Y {
            b.Pieces[i].MoveTo(to)
            return
        }
    }
    panic("No piece at position")
}

func InitialPositionBoard() Board {
    initialFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
    return BoardFromFen(initialFen)
}
