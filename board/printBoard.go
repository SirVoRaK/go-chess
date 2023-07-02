package board

import (
	"fmt"
	"piece"
	"strings"
)

const (
    ColorReset = "\033[0m"

    ColorRed = "\033[31m"
    ColorGreen = "\033[32m"
    ColorYellow = "\033[33m"
    ColorBlue = "\033[34m"
    ColorPurple = "\033[35m"
    ColorCyan = "\033[36m"
    ColorWhite = "\033[37m"
)
func (b *Board) Print() {
    lineLength := uint(9 * 4 + 1)
    rowNumber := uint(8)
    PrintLine(lineLength)
    for row := uint(0); row < 8; row++ {
        for col := uint(0); col < 8; col++ {
            currentPiece, _ := b.PieceAt(col, row)
            letter := " "
            if currentPiece != nil {
                letter = strings.ToUpper(currentPiece.Letter())
            }
            pieceLetterColor := ColorWhite
            if currentPiece != nil && currentPiece.Color == piece.BLACK {
                pieceLetterColor = ColorYellow
            }
            fmt.Printf("%s|%s %s%s ", ColorCyan, pieceLetterColor, letter, ColorReset)

            if col == 7 {
                fmt.Printf("%s|%s %d%s ", ColorCyan, ColorBlue, rowNumber, ColorReset)
                rowNumber--
            }
        }
        fmt.Printf("%s|%s\n", ColorCyan, ColorReset)
        PrintLine(lineLength)
    }
    letters := "abcdefgh "
    for i := 0; i < len(letters); i++ {
        fmt.Printf("%s|%s %c%s ", ColorCyan, ColorBlue, letters[i], ColorReset)
    }
    fmt.Printf("%s|%s\n", ColorCyan, ColorReset)
    PrintLine(lineLength)
}

func PrintLine(size uint) {
    fmt.Printf("%s+", ColorCyan)
    for i := uint(0); i < size - 2; i++ {
        fmt.Print("-")
    }
    fmt.Printf("+%s\n", ColorReset)
}
