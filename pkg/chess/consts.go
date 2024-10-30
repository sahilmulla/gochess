package chess

const (
	NumberOfTiles = 64
)

type Move int

const (
	Advance Move = iota
	PawnJump

	Attack
)

type Team byte

func (t Team) String() string {
	return string(t)
}

const (
	None  Team = 'n'
	White Team = 'w'
	Black Team = 'b'
)

type Vector int

const (
	N Vector = -8
	S Vector = 8
	E Vector = 1
	W Vector = -1
)

var (
	StandardPlacement = map[int]Piece{
		0:  BlackRook,
		1:  BlackKnight,
		2:  BlackBishop,
		3:  BlackQueen,
		4:  BlackKing,
		5:  BlackBishop,
		6:  BlackKnight,
		7:  BlackRook,
		8:  BlackPawn,
		9:  BlackPawn,
		10: BlackPawn,
		11: BlackPawn,
		12: BlackPawn,
		13: BlackPawn,
		14: BlackPawn,
		15: BlackPawn,

		48: WhitePawn,
		49: WhitePawn,
		50: WhitePawn,
		51: WhitePawn,
		52: WhitePawn,
		53: WhitePawn,
		54: WhitePawn,
		55: WhitePawn,
		56: WhiteRook,
		57: WhiteKnight,
		58: WhiteBishop,
		59: WhiteQueen,
		60: WhiteKing,
		61: WhiteBishop,
		62: WhiteKnight,
		63: WhiteRook,
	}
)
