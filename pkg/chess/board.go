package chess

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

type MoveLog struct {
	From, To        int
	Moved, Captured Piece
	Team            Team
	MoveType        Move
}

type EnPassantInfo struct {
	PassingTileId, CaptureTileId int
}

type CastleInfo struct {
	ks, qs map[Team]bool
}

type Board struct {
	Tiles         []Tile
	Next          Team
	MoveLog       []MoveLog
	enPassantInfo EnPassantInfo
	castleInfo    CastleInfo
}

func (b *Board) Move(from, to int) (*Piece, error) {
	toMove := b.Tiles[from].Piece

	if toMove.Team() != b.Next {
		return nil, errors.New("illegal turn")
	}

	moves := b.AvailableMoves(from)

	m, has := moves[to]
	if !has || m == EnPassantCapture {
		return nil, errors.New("illegal move")
	}

	captured := b.Tiles[to].Piece

	if !(m == KingSideCastle || m == QueenSideCastle) {
		b.Tiles[to] = b.Tiles[from]
	} else {
		b.Tiles[to] = Tile{Piece: EmptyPiece}
	}

	b.Tiles[from] = Tile{Piece: EmptyPiece}

	if m == EnPassantAttack {
		captured = b.Tiles[b.enPassantInfo.CaptureTileId].Piece
		b.Tiles[b.enPassantInfo.CaptureTileId] = Tile{Piece: EmptyPiece}
	}

	if m == PawnTwoStep {
		if b.Next == Black {
			b.enPassantInfo = EnPassantInfo{PassingTileId: to + int(N), CaptureTileId: to}
		} else if b.Next == White {
			b.enPassantInfo = EnPassantInfo{PassingTileId: to + int(S), CaptureTileId: to}
		}
	} else {
		b.enPassantInfo = EnPassantInfo{CaptureTileId: -1, PassingTileId: -1}
	}

	if toMove == BlackKing || toMove == WhiteKing || m == KingSideCastle || m == QueenSideCastle {
		b.castleInfo.ks[toMove.Team()] = true
		b.castleInfo.qs[toMove.Team()] = true
	} else if toMove == BlackRook && from == 0 || toMove == WhiteRook && from == 56 {
		b.castleInfo.qs[toMove.Team()] = true
	} else if toMove == BlackRook && from == 7 || toMove == WhiteRook && from == 63 {
		b.castleInfo.ks[toMove.Team()] = true
	}

	if m == KingSideCastle {
		if captured == WhiteRook || captured == BlackRook {
			captured, toMove = toMove, captured
			from, to = to, from
		}
		b.Tiles[to+2] = Tile{Piece: captured}
		b.Tiles[to+1] = Tile{Piece: toMove}
		captured = EmptyPiece
	} else if m == QueenSideCastle {
		if captured == WhiteRook || captured == BlackRook {
			captured, toMove = toMove, captured
			from, to = to, from
		}
		b.Tiles[to-2] = Tile{Piece: captured}
		b.Tiles[to-1] = Tile{Piece: toMove}
		captured = EmptyPiece
	}

	b.MoveLog = append(b.MoveLog, MoveLog{From: from, To: to, Moved: toMove, Captured: captured, Team: b.Next, MoveType: m})

	if b.Next == Black {
		b.Next = White
	} else {
		b.Next = Black
	}

	return &captured, nil
}

func (b *Board) AvailableMoves(tileId int) map[int]Move {
	moves := make(map[int]Move)

	if tileId < 0 {
		return moves
	}

	tile := b.TileAt(tileId)

	if tile.Piece == WhiteRook || tile.Piece == BlackRook {
		for _, vec := range []Vector{N, S, E, W} {
			for currId := tileId + int(vec); ; currId += int(vec) {
				if currId < 0 || currId >= NumberOfTiles || vec == E && currId%8 == 0 || vec == W && currId%8 == 7 {
					break
				}
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
					break
				}
				if b.TileAt(currId).Piece.Team() == tile.Piece.Team() {
					break
				}
				moves[currId] = Advance
			}
		}

		if tileId == 63 || tileId == 7 {
			_, cantCastleDueToMove := b.castleInfo.ks[tile.Piece.Team()]

			clearPathToCastle := true
			if !cantCastleDueToMove {
				if tile.Piece.Team() == White {
					for i := 61; i < 63; i++ {
						if b.TileAt(i).Piece != EmptyPiece {
							clearPathToCastle = false
							break
						}
					}
				}

				if tile.Piece.Team() == Black {
					for i := 5; i < 7; i++ {
						if b.TileAt(i).Piece != EmptyPiece {
							clearPathToCastle = false
							break
						}
					}
				}
			}

			if !cantCastleDueToMove && clearPathToCastle {
				if tile.Piece.Team() == White && b.TileAt(60).Piece == WhiteKing {
					moves[60] = KingSideCastle
				} else if tile.Piece.Team() == Black && b.TileAt(4).Piece == BlackKing {
					moves[4] = KingSideCastle
				}
			}
		}
	}

	if tile.Piece == WhiteBishop || tile.Piece == BlackBishop {
		for _, vec := range []Vector{N + E, N + W, S + E, S + W} {
			for currId := tileId + int(vec); ; currId += int(vec) {
				if currId < 0 || currId >= NumberOfTiles || ((vec == N+E || vec == S+E) && currId%8 == 0) || ((vec == N+W || vec == S+W) && currId%8 == 7) {
					break
				}
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
					break
				}
				if b.TileAt(currId).Piece.Team() == tile.Piece.Team() {
					break
				}
				moves[currId] = Advance

				if mod := currId % 8; mod == 0 || mod == 7 {
					break
				}
			}
		}
	}

	if tile.Piece == WhiteQueen || tile.Piece == BlackQueen {
		for _, vec := range []Vector{N, S, E, W, N + E, N + W, S + E, S + W} {
			for currId := tileId + int(vec); ; currId += int(vec) {
				if currId < 0 || currId >= NumberOfTiles || ((vec == E || vec == N+E || vec == S+E) && currId%8 == 0) || ((vec == W || vec == N+W || vec == S+W) && currId%8 == 7) {
					break
				}
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
					break
				}
				if b.TileAt(currId).Piece.Team() == tile.Piece.Team() {
					break
				}
				moves[currId] = Advance

				if mod := currId % 8; mod == 0 || mod == 7 {
					break
				}
			}
		}
	}

	if tile.Piece == WhiteKing || tile.Piece == BlackKing {
		if tileId == 4 {
			_, cantCastleDueToMove := b.castleInfo.ks[tile.Piece.Team()]

			clearPathToCastle := true
			if !cantCastleDueToMove {
				if tile.Piece.Team() == White {
					for i := 61; i < 63; i++ {
						if b.TileAt(i).Piece != EmptyPiece {
							clearPathToCastle = false
							break
						}
					}
				}

				if tile.Piece.Team() == Black {
					for i := 5; i < 7; i++ {
						if b.TileAt(i).Piece != EmptyPiece {
							clearPathToCastle = false
							break
						}
					}
				}
			}

			if !cantCastleDueToMove && clearPathToCastle {
				if tile.Piece.Team() == White {
					moves[63] = KingSideCastle
				} else if tile.Piece.Team() == Black {
					moves[7] = KingSideCastle
				}
			}

		}

		for _, vec := range []Vector{N, S, E, W, N + E, N + W, S + E, S + W} {
			currId := tileId + int(vec)
			if currId < 0 || currId >= NumberOfTiles || (currId%8 == 0 && (vec == E || vec == N+E || vec == S+E)) || (currId%8 == 7 && (vec == W || vec == N+W || vec == S+W)) {
				continue
			}
			if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
				moves[currId] = Attack
				continue
			}
			if b.TileAt(currId).Piece.Team() == tile.Piece.Team() {
				continue
			}
			moves[currId] = Advance

			if mod := currId % 8; mod == 0 || mod == 7 {
				continue
			}
		}
	}

	if tile.Piece == WhiteKnight || tile.Piece == BlackKnight {
		for _, vec := range []Vector{N + N + E, N + N + W, S + S + E, S + S + W, W + W + N, W + W + S, E + E + N, E + E + S} {
			currId := tileId + int(vec)
			if currId < 0 || currId >= NumberOfTiles ||
				((tileId%8 == 0 || tileId%8 == 1) && (vec == W+W+N || vec == W+W+S)) ||
				((tileId%8 == 7 || tileId%8 == 6) && (vec == E+E+N || vec == E+E+S)) ||
				(tileId%8 == 0 && (vec == N+N+W || vec == S+S+W)) ||
				(tileId%8 == 7 && (vec == N+N+E || vec == S+S+E)) {
				continue
			}
			if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
				moves[currId] = Attack
				continue
			}
			if b.TileAt(currId).Piece.Team() == tile.Piece.Team() {
				continue
			}
			moves[currId] = Advance

			if mod := currId % 8; mod == 0 || mod == 7 {
				continue
			}
		}
	}

	if tile.Piece == BlackPawn {
		skipTwoStep := false
		for _, vec := range []Vector{S, S + S, S + W, S + E} {
			currId := tileId + int(vec)

			if currId < 0 || currId >= NumberOfTiles ||
				(tileId%8 == 7 && vec == S+E) ||
				(tileId%8 == 0 && vec == S+W) {
				continue
			}
			if vec == S+S && (skipTwoStep || tileId/8 != 1 || b.TileAt(currId).Piece.Team() != None) {
				continue
			}
			if vec == S && b.TileAt(currId).Piece.Team() != None {
				skipTwoStep = true
				continue
			}
			if vec == S+E || vec == S+W {
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
				} else if currId == b.enPassantInfo.PassingTileId && b.TileAt(b.enPassantInfo.CaptureTileId).Piece.Team() != tile.Piece.Team() {
					moves[currId] = EnPassantAttack
					moves[b.enPassantInfo.CaptureTileId] = EnPassantCapture
				}
				continue
			}

			if vec == S+S {
				moves[currId] = PawnTwoStep
			} else {
				moves[currId] = Advance
			}

			if mod := currId % 8; mod == 0 || mod == 7 {
				continue
			}
		}
	}

	if tile.Piece == WhitePawn {
		skipTwoStep := false
		for _, vec := range []Vector{N, N + N, N + W, N + E} {
			currId := tileId + int(vec)
			if currId < 0 || currId >= NumberOfTiles ||
				(tileId%8 == 7 && vec == N+E) ||
				(tileId%8 == 0 && vec == N+W) {
				continue
			}
			if vec == N+N && (skipTwoStep || tileId/8 != 6 || b.TileAt(currId).Piece.Team() != None) {
				continue
			}
			if vec == N && b.TileAt(currId).Piece.Team() != None {
				skipTwoStep = true
				continue
			}
			if vec == N+E || vec == N+W {
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
				} else if currId == b.enPassantInfo.PassingTileId && b.TileAt(b.enPassantInfo.PassingTileId).Piece.Team() != tile.Piece.Team() {
					moves[currId] = EnPassantAttack
					moves[b.enPassantInfo.CaptureTileId] = EnPassantCapture
				}
				continue
			}

			if vec == N+N {
				moves[currId] = PawnTwoStep
			} else {
				moves[currId] = Advance
			}
			if mod := currId % 8; mod == 0 || mod == 7 {
				continue
			}
		}
	}

	return moves
}

func (b *Board) TileAt(idx int) Tile {
	return b.Tiles[idx]
}

func (b *Board) Debug(activeId int) string {
	var buffer bytes.Buffer

	buffer.WriteRune('\n')
	fmt.Fprintf(&buffer, "%+v \n", b.enPassantInfo)
	fmt.Fprintf(&buffer, "%+v \n", b.castleInfo)

	moves := b.AvailableMoves(activeId)
	fmt.Fprintln(&buffer, moves)

	for tileId, tile := range b.Tiles {
		if tileId%8 == 0 {
			fmt.Fprintf(&buffer, "%d\t", tileId)
		}

		checkerIt := func(s string) string {
			if tileId/8%2^tileId%2 == 0 {
				return fmt.Sprintf("\033[100m%s\033[0m", s)
			}
			return fmt.Sprintf("\033[49m%s\033[0m", s)
		}

		if move, has := moves[tileId]; has {
			switch move {
			case Attack, EnPassantCapture:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[31m %s \033[0m", tile.Piece)))
			case EnPassantAttack:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[31m %s \033[0m", "*")))
			case Advance, PawnTwoStep:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[34m %s \033[0m", "·")))
			case KingSideCastle, QueenSideCastle:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[34m %s \033[0m", tile.Piece)))
			default:
				log.Fatalf("unknown move %v", move)
			}
		} else {
			if tileId == activeId {
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[32m %s \033[0m", tile.Piece)))
			} else {
				buffer.WriteString(checkerIt(fmt.Sprintf(" %s ", tile.Piece)))
			}
		}

		if tileId%8 == 7 && tileId < 8*7 {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()

}

func NewBoard(options ...BoardOption) *Board {
	board := &Board{
		Tiles:         make([]Tile, NumberOfTiles),
		Next:          White,
		enPassantInfo: EnPassantInfo{PassingTileId: -1, CaptureTileId: -1},
	}

	board.castleInfo = CastleInfo{ks: make(map[Team]bool), qs: make(map[Team]bool)}

	board.Tiles = make([]Tile, NumberOfTiles)

	for t := range board.Tiles {
		board.Tiles[t] = Tile{Piece: EmptyPiece}
	}

	for _, opt := range options {
		opt(board)
	}

	return board
}

type BoardOption func(*Board)

func WithStartTeam(t Team) BoardOption {
	return func(b *Board) {
		b.Next = t
	}
}

func WithStandardPlacement() BoardOption {
	return WithCustomPlacement(StandardPlacement)
}

func WithCustomPlacement(initPieces map[int]Piece) BoardOption {
	return func(b *Board) {
		for tileId, p := range initPieces {
			b.Tiles[tileId] = Tile{Piece: p}
		}
	}
}

type Tile struct {
	Piece Piece
}
