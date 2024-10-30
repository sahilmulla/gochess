package main

import (
	"fmt"

	"gochess/pkg/chess"
)

func main() {
	b := chess.NewBoard(chess.WithCustomStartPosition(map[int]chess.Piece{54: chess.BlackRook, 50: chess.BlackKing, 22: chess.WhiteQueen}))

	fmt.Println(b)
}
