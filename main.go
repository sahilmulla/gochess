package main

import (
	"fmt"

	"gochess/pkg/chess"
)

func main() {
	// b := chess.NewBoard(chess.WithCustomPlacement(map[int]chess.Piece{42: chess.BlackRook, 49: chess.WhitePawn, 22: chess.WhiteQueen, 47: chess.WhitePawn, 40: chess.BlackBishop}))
	// b := chess.NewBoard(chess.WithStandardPlacement())

	p := chess.StandardPlacement
	b := chess.NewBoard(chess.WithCustomPlacement(p))
	fmt.Println(b.Move(52, 44))
	fmt.Println(b.Move(12, 20))
	fmt.Println(b.Move(44, 36))
	fmt.Println(b.Move(20, 28))
	fmt.Println(b.Move(57, 42))

	fmt.Println(b.Debug(42))
	fmt.Printf("%+v\n", b.MoveLog)
}
