package board

import (
	"fmt"
	"piece"
	"strconv"
	"strings"
)

func BoardFromFen(fen string) Board {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("Invalid FEN: %s", r))
		}
	}()
	board := Board{}
	board.fillDefaults()

	segments := strings.Split(fen, " ")

	position := segments[0]
	x := uint(0)
	y := uint(0)
	rows := strings.Split(position, "/")
	if len(rows) != 8 {
		panic("Wrong number of rows")
	}
	for _, row := range rows {
	letters_loop:
		for i := 0; i < len(row); i++ {
			if x > 7 {
				panic("Too many pieces in a row")
			}
			letter := string(row[i])
			number, err := strconv.ParseUint(letter, 10, 16)
			if err == nil {
				x += uint(number)
				continue letters_loop
			}
			newPiece := piece.LetterToPiece(letter, piece.Position{X: x, Y: y})
			board.Pieces[y][x] = &newPiece
			x++
		}
		y++
		x = 0
	}
	if len(segments) == 1 {
		board.turn = piece.WHITE
		return board
	}

	turn := string(segments[1])
	if turn == "w" {
		board.turn = piece.WHITE
	} else if turn == "b" {
		board.turn = piece.BLACK
	} else {
		panic(fmt.Sprintf("Invalid turn color - %s", turn))
	}

	if len(segments) == 2 {
		return board
	}

	castling := string(segments[2])
	if len(castling) > 4 {
		panic(fmt.Sprintf("Invalid castling - %s", castling))
	}
	if castling == "-" {
		board.Castling = Castling{false, false, false, false}
	} else {
		for _, letter := range castling {
			switch letter {
			case 'K':
				board.Castling.WhiteKingSide = true
			case 'Q':
				board.Castling.WhiteQueenSide = true
			case 'k':
				board.Castling.BlackKingSide = true
			case 'q':
				board.Castling.BlackQueenSide = true
			default:
				panic(fmt.Sprintf("Invalid castling - %s", castling))
			}
		}
	}

	return board
}
