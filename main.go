package main

import (
	"fmt"
	"gochess/pkg/chess"
)

func main() {
	b := chess.NewBoard(chess.WithStandardPlacement())
	fmt.Println(b.Debug(-1))
}

func italianGame() {
	b := chess.NewBoard(chess.WithStandardPlacement())
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
	fmt.Println(b.Move(63, 60))
	fmt.Println(b.Move(8, 16))
	fmt.Println(b.Move(61, 60))
	fmt.Println(b.Move(26, 8))
	fmt.Println(b.Move(48, 32))
	fmt.Println(b.Move(4, 7))
	fmt.Println(b.Move(55, 47))
	fmt.Println(b.Move(18, 12))
	fmt.Println(b.Move(43, 35))
	fmt.Println(b.Move(12, 22))
	fmt.Println(b.Move(57, 51))
	fmt.Println(b.Move(10, 18))
	fmt.Println(b.Move(34, 43))
	fmt.Println(b.Move(5, 4))
	fmt.Println(b.Move(43, 50))
	fmt.Println(b.Move(15, 23))
	fmt.Println(b.Move(51, 61))
	fmt.Println(b.Move(28, 35))
	fmt.Println(b.Move(42, 35))
	fmt.Println(b.Move(18, 26))
	fmt.Println(b.Move(35, 27))
	fmt.Println(b.Move(9, 25))
	fmt.Println(b.Move(32, 25))
	fmt.Println(b.Move(16, 25))
	fmt.Println(b.Move(61, 46))
	fmt.Println(b.Move(2, 11))
	fmt.Println(b.Debug(58))
	fmt.Println(b.Move(58, 44))
	fmt.Println(b.Debug(8))
	fmt.Println(b.Move(8, 17))
	fmt.Println(b.Debug(56))
	fmt.Println(b.Move(56, 0))
	fmt.Println(b.Debug(3))
	fmt.Println(b.Move(3, 0))
	fmt.Println(b.Debug(49))
	fmt.Println(b.Move(49, 33))
	fmt.Println(b.Debug(0))
	fmt.Println(b.Move(0, 8))
	fmt.Println(b.Debug(59))
	fmt.Println(b.Move(59, 56))
	fmt.Println(b.Debug(8))
	fmt.Println(b.Move(8, 10))
	fmt.Println(b.Debug(44))
	fmt.Println(b.Move(44, 23))
	fmt.Println(b.Debug(26))
	fmt.Println(b.Move(26, 33))
	fmt.Println(b.Debug(23))
	fmt.Println(b.Move(23, 14))
	fmt.Println(b.Debug(10))
	fmt.Println(b.Move(10, 50))
	fmt.Println(b.Debug(56))
	fmt.Println(b.Move(56, 21))
	fmt.Println(b.Debug(56))

	// fmt.Printf("%+v\n", b.MoveLog)
}
