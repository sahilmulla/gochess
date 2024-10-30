package chess

import "fmt"

type Piece struct {
	Role
	FirstMove bool
}

func (p Piece) String() string {
	return fmt.Sprintf("firstmove:%t role:%v", p.FirstMove, p.Role)
}

func NewPiece(r Role) *Piece {
	return &Piece{Role: r, FirstMove: true}
}
func NewEmptyPiece() *Piece {
	return &Piece{Role: Empty}
}

type Role byte

func (r Role) Team() Team {
	if r == Empty {
		return None
	} else if r >= 97 {
		return Black
	} else {
		return White
	}
}

func (r Role) String() string {
	return string(r)
}

const (
	BlackKing   Role = 'k'
	BlackQueen  Role = 'q'
	BlackBishop Role = 'b'
	BlackKnight Role = 'n'
	BlackRook   Role = 'r'
	BlackPawn   Role = 'p'

	WhiteKing   Role = 'K'
	WhiteQueen  Role = 'Q'
	WhiteBishop Role = 'B'
	WhiteKnight Role = 'N'
	WhiteRook   Role = 'R'
	WhitePawn   Role = 'P'

	Empty Role = ' '
)
