package board

import (
	"fmt"
	"piece"
	"strings"
)

const (
	ColorReset = "\033[0m"

	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func (b *Board) Print() {
	lineLength := uint(8*4 + 1)
	rowNumber := uint(8)
	fmt.Print("\n\n")
	PrintLine(lineLength)
	for row := uint(0); row < 8; row++ {
		for col := uint(0); col < 8; col++ {
			currentPiece := b.PieceAt(col, row)
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
		fmt.Print("\n")
		PrintLine(lineLength)
	}
	letters := "abcdefgh "
	for i := 0; i < len(letters); i++ {
		fmt.Printf(" %s %c%s ", ColorBlue, letters[i], ColorReset)
	}
	fmt.Print("\n")
	fmt.Println("FEN:", b.ToFen())
}

func PrintLine(size uint) {
	for i := uint(0); i < size; i++ {
		if i%4 == 0 {
			fmt.Print(ColorCyan, "+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("\n", ColorReset)
}
