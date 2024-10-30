package main

import (
	"fmt"
	"gochess/pkg/chess"
)

func main() {
	b := chess.NewBoard(chess.WithCustomPlacement(map[int]chess.Piece{27: chess.WhitePawn, 29: chess.WhitePawn, 22: chess.WhiteQueen, 12: chess.BlackPawn}), chess.WithStartTeam(chess.Black))

	fmt.Println(b.Debug(12))
	fmt.Println(b.Move(12, 28))
	fmt.Println(b.Debug(27))
	fmt.Println(b.Move(27, 20))

	fmt.Println(b.MoveLog)

	// b := chess.NewBoard(chess.WithStandardPlacement())

	// p := chess.StandardPlacement
	// b := chess.NewBoard(chess.WithCustomPlacement(p))
	// fmt.Println(b.Move(52, 36))
	// fmt.Println(b.Move(12, 28))
	// fmt.Println(b.Move(62, 45))
	// fmt.Println(b.Move(1, 18))
	// fmt.Println(b.Move(61, 34))
	// fmt.Println(b.Move(5, 26))
	// fmt.Println(b.Move(50, 42))
	// fmt.Println(b.Move(6, 21))
	// fmt.Println(b.Move(51, 43))
	// fmt.Println(b.Move(11, 19))

	// fmt.Println(b.Debug(52))
	// fmt.Printf("%+v\n", b.MoveLog)
}
