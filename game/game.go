package game

import (
    "board"
    "piece"
)

type Game struct {
    Board board.Board
    Turn piece.Color
}

func NewGame() *Game {
    game := new(Game)
    game.Board = board.InitialPositionBoard()
    game.Turn = piece.WHITE
    game.Board.Print()
    return game
}
