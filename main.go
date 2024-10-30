package main

import (
	"fmt"

	"gochess/pkg/chess"
)

func main() {
	b := chess.NewBoard(chess.WithCustomStartPosition(map[int]chess.Piece{54: chess.BlackRook, 49: chess.WhitePawn, 22: chess.WhiteQueen, 47: chess.WhitePawn, 40: chess.BlackBishop}))
	// b := chess.NewBoard(chess.WithStandardStartPosition())

	fmt.Println(b.Debug(49))
}
