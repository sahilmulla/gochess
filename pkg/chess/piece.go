package chess

type Piece byte

func (r Piece) Team() Team {
	if r == EmptyPiece {
		return None
	} else if r >= 97 {
		return Black
	} else {
		return White
	}
}

func (r Piece) String() string {
	return string(r)
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
