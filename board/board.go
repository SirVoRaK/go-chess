package board

import (
    "fmt"
    "piece"
)

type Board struct {
    Pieces []piece.Piece
}
func (b *Board) Print() {
    lineLength := uint(9 * 4 + 1)
    rowNumber := uint(8)
    PrintLine(lineLength)
    for row := uint(0); row < 8; row++ {
        for col := uint(0); col < 8; col++ {
            piece := b.PieceAt(col, row)
            letter := " "
            if piece != nil {
                letter = piece.Letter()
            }
            fmt.Printf("| %s ", letter)

            if col == 7 {
                fmt.Printf("| %d ", rowNumber)
                rowNumber--
            }
        }
        fmt.Println("|")
        PrintLine(lineLength)
    }
    letters := "abcdefgh "
    for i := 0; i < len(letters); i++ {
        fmt.Printf("| %c ", letters[i])
    }
    fmt.Println("|")
    PrintLine(lineLength)
}
func (b *Board) PieceAt(x uint, y uint) *piece.Piece {
    for _, p := range b.Pieces {
        if p.Position.X == x && p.Position.Y == y {
            return &p
        }
    }
    return nil
}

func PrintLine(size uint) {
    fmt.Print("+")
    for i := uint(0); i < size - 2; i++ {
        fmt.Print("-")
    }
    fmt.Println("+")
}

var initialPosition = [][]string {
    {"r", "n", "b", "q", "k", "b", "n", "r"},
    {"p", "p", "p", "p", "p", "p", "p", "p"},
    {" ", " ", " ", " ", " ", " ", " ", " "},
    {" ", " ", " ", " ", " ", " ", " ", " "},
    {" ", " ", " ", " ", " ", " ", " ", " "},
    {" ", " ", " ", " ", " ", " ", " ", " "},
    {"P", "P", "P", "P", "P", "P", "P", "P"},
    {"R", "N", "B", "Q", "K", "B", "N", "R"},
}
func InitialPositionBoard() Board {
    board := Board{}
    for row := uint(0); row < 8; row++ {
        for col := uint(0); col < 8; col++ {
            letter := initialPosition[row][col]
            if letter != " " {
                board.Pieces = append(board.Pieces, piece.LetterToPiece(letter, piece.Position{X: col, Y: row}))
            }
        }
    }
    return board
}
