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
func (b *Board) PieceAt(x uint, y uint) *piece.Piece {
    for i, p := range b.Pieces {
        if p.Position.X == x && p.Position.Y == y {
            return &b.Pieces[i]
        }
    }
    return nil
}
func (b *Board) MovePiece(from piece.Position, to piece.Position) {
    fromPiece := b.PieceAt(from.X, from.Y)
    if fromPiece == nil {
        panic("No piece at position")
    }

    if fromPiece.Color != b.turn {
        panic("Not your turn")
    }

    fromPiece.Move(to)
}

func InitialPositionBoard() Board {
    initialFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
    return BoardFromFen(initialFen)
}
