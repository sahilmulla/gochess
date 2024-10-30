package chess

type Piece byte

func (p Piece) String() string {
	return string(p)
}

func (p Piece) Team() Team {
	if p == EmptyPiece {
		return None
	} else if p >= 97 {
		return Black
	} else {
		return White
	}
}

const (
	BlackKing   Piece = 'k'
	BlackQueen  Piece = 'q'
	BlackBishop Piece = 'b'
	BlackKnight Piece = 'n'
	BlackRook   Piece = 'r'
	BlackPawn   Piece = 'p'

	WhiteKing   Piece = 'K'
	WhiteQueen  Piece = 'Q'
	WhiteBishop Piece = 'B'
	WhiteKnight Piece = 'N'
	WhiteRook   Piece = 'R'
	WhitePawn   Piece = 'P'

	EmptyPiece Piece = ' '
)
