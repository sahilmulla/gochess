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
	fmt.Println(b.Move(52, 36))
	fmt.Println(b.Move(12, 28))
	fmt.Println(b.Move(62, 45))
	fmt.Println(b.Move(1, 18))
	fmt.Println(b.Move(61, 34))
	fmt.Println(b.Move(5, 26))
	fmt.Println(b.Move(50, 42))
	fmt.Println(b.Move(6, 21))
	fmt.Println(b.Move(51, 43))
	fmt.Println(b.Move(11, 19))

	fmt.Println(b.Debug(11))
	fmt.Printf("%+v\n", b.MoveLog)
}
