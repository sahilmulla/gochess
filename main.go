package main

import (
	"fmt"

	"gochess/pkg/chess"
)

func main() {
	b := chess.NewBoard(chess.WithCustomStartPosition(map[int]chess.Piece{54: chess.BlackRook, 47: chess.WhiteKnight, 22: chess.WhiteQueen, 56: chess.BlackRook}))
	// b := chess.NewBoard(chess.WithStandardStartPosition())

	fmt.Println(b.Debug(47))
}
