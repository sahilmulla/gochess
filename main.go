package main

import (
	"fmt"
	"gochess/pkg/chess"
)

func main() {
	b := chess.NewBoard(chess.WithCustomStartPosition(map[int]chess.Piece{52: chess.WhiteRook, 50: chess.BlackPawn}))

	fmt.Println(b)
}
