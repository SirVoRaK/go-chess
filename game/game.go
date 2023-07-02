package game

import (
	"board"
	"fmt"
	"os"
	"piece"
	"strconv"
)

type State string
const (
    PLAYING State = "playing"
    STARTING State = "starting"
    FINISHED State = "finished"
)
type Game struct {
    Board board.Board
    Turn piece.Color
    State State
}

func NewGame() *Game {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
            os.Exit(1)
        }
    }()
    game := new(Game)
    game.Board = board.InitialPositionBoard()
    game.Turn = game.Board.GetTurn()
    game.Board.Print()
    return game
}

func (g *Game) Start() {
    g.State = PLAYING
    g.gameLoop()
}
func (g *Game) gameLoop() {
    for g.State == PLAYING {
        fmt.Printf("%s's turn\n", g.Turn)
        g.playTurn()
        g.Board.Print()
    }
}

func (g *Game) playTurn() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Invalid move: %s\n", r)
            g.playTurn()
            g.Board.Print()
        }
    }()

    from, to := inputMove()
    g.Board.MovePiece(from, to)
    g.Turn = g.Turn.Opposite()
    g.Board.ChangeTurn()
}

func inputMove() (piece.Position, piece.Position) {
    var move string

    fmt.Print("Enter move: ")
    fmt.Scanln(&move)

    if len(move) != 4 {
        panic("Invalid move")
    }

    from := piece.Position{}
    to := piece.Position{}

    from.X = piece.LetterToColumn(string(move[0]))
    from.Y = piece.NumberToRow(toUint(string(move[1])))
    to.X = piece.LetterToColumn(string(move[2]))
    to.Y = piece.NumberToRow(toUint(string(move[3])))

    return from, to
}
func toUint(text string) uint {
    value, err := strconv.ParseUint(text, 10, 32)
    if err != nil {
        panic("Invalid number")
    }
    return uint(value)
}
