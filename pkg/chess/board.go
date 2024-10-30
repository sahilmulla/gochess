package chess

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

type Log struct {
	From, To        int
	Moved, Captured Role
	Team            Team
}

type EnPassantInfo struct {
	PassingTileId int
}

type Board struct {
	Tiles         []Tile
	Next          Team
	MoveLog       []Log
	enPassantInfo EnPassantInfo
}

func (b *Board) Move(from, to int) (*Role, error) {
	toMove := b.Tiles[from].Piece

	if toMove.Team() != b.Next {
		return nil, errors.New("illegal turn")
	}

	moves := b.AvailableMoves(from)

	m, has := moves[to]
	if !has {
		return nil, errors.New("illegal move")
	}

	captured := b.Tiles[to].Piece

	b.Tiles[to] = b.Tiles[from]

	b.Tiles[from] = Tile{Piece: NewEmptyPiece()}

	if m == PawnJump {
		if b.Next == Black {
			b.enPassantInfo.PassingTileId = to + int(N)
		} else if b.Next == White {
			b.enPassantInfo.PassingTileId = to + int(S)
		}
	}

	toMove.FirstMove = false
	b.MoveLog = append(b.MoveLog, Log{From: from, To: to, Moved: toMove.Role, Captured: captured.Role, Team: b.Next})

	if b.Next == Black {
		b.Next = White
	} else {
		b.Next = Black
	}

	return &captured.Role, nil
}

func (b *Board) AvailableMoves(tileId int) map[int]Move {
	moves := make(map[int]Move)

	if tileId < 0 {
		return moves
	}

	tile := b.TileAt(tileId)

	if tile.Piece.Role == WhiteRook || tile.Piece.Role == BlackRook {
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
	}

	if tile.Piece.Role == WhiteBishop || tile.Piece.Role == BlackBishop {
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

	if tile.Piece.Role == WhiteQueen || tile.Piece.Role == BlackQueen {
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

	if tile.Piece.Role == WhiteKing || tile.Piece.Role == BlackKing {
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

	if tile.Piece.Role == WhiteKnight || tile.Piece.Role == BlackKnight {
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

	if tile.Piece.Role == BlackPawn {
		skipJump := false
		for _, vec := range []Vector{S, S + S, S + W, S + E} {
			currId := tileId + int(vec)

			if currId < 0 || currId >= NumberOfTiles ||
				(tileId%8 == 7 && vec == S+E) ||
				(tileId%8 == 0 && vec == S+W) {
				continue
			}
			if vec == S+S && (skipJump || tileId/2 != 1 || b.TileAt(currId).Piece.Team() != None) {
				continue
			}
			if vec == S && b.TileAt(currId).Piece.Team() != None {
				skipJump = true
				continue
			}
			if vec == S+E || vec == S+W {
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
				}
				continue
			}

			if vec == S+S {
				moves[currId] = PawnJump
			} else {
				moves[currId] = Advance
			}

			if mod := currId % 8; mod == 0 || mod == 7 {
				continue
			}
		}
	}

	if tile.Piece.Role == WhitePawn {
		skipJump := false
		for _, vec := range []Vector{N, N + N, N + W, N + E} {
			currId := tileId + int(vec)
			if currId < 0 || currId >= NumberOfTiles ||
				(tileId%8 == 7 && vec == N+E) ||
				(tileId%8 == 0 && vec == N+W) {
				continue
			}
			if vec == N+N && (skipJump || tileId/8 != 6 || b.TileAt(currId).Piece.Team() != None) {
				continue
			}
			if vec == N && b.TileAt(currId).Piece.Team() != None {
				skipJump = true
				continue
			}
			if vec == N+E || vec == N+W {
				if other := b.TileAt(currId).Piece.Team(); other != None && other != tile.Piece.Team() {
					moves[currId] = Attack
				}
				continue
			}

			if vec == N+N {
				moves[currId] = PawnJump
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
	fmt.Fprintf(&buffer, "%v \n", b.enPassantInfo)

	moves := b.AvailableMoves(activeId)
	fmt.Fprintln(&buffer, moves)

	for tileId, tile := range b.Tiles {
		if tileId%8 == 0 {
			fmt.Fprintf(&buffer, "%d\t", tileId)
		}

		checkerIt := func(s string) string {
			if tileId/8%2^tileId%2 == 0 {
				return fmt.Sprintf("\033[49m%s\033[0m", s)
			}
			return fmt.Sprintf("\033[100m%s\033[0m", s)
		}

		if move, has := moves[tileId]; has {
			switch move {
			case Attack:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[31m %s \033[0m", tile.Piece.Role)))
			case Advance, PawnJump:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[34m %s \033[0m", "·")))
			default:
				log.Fatalf("unknown move %v", move)
			}
		} else {
			if tileId == activeId {
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[32m %s \033[0m", tile.Piece.Role)))
			} else {
				buffer.WriteString(checkerIt(fmt.Sprintf(" %s ", tile.Piece.Role)))
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
		enPassantInfo: EnPassantInfo{PassingTileId: -1},
	}

	board.Tiles = make([]Tile, NumberOfTiles)

	for t := range board.Tiles {
		board.Tiles[t] = Tile{Piece: NewEmptyPiece()}
	}

	for _, opt := range options {
		opt(board)
	}

	return board
}

type BoardOption func(*Board)

func WithStandardPlacement() BoardOption {
	return WithCustomPlacement(StandardPlacement)
}

func WithCustomPlacement(initPieces map[int]Role) BoardOption {
	return func(b *Board) {
		for pos, role := range initPieces {
			b.Tiles[pos] = Tile{Piece: NewPiece(role)}
		}
	}
}

type Tile struct {
	Piece *Piece
}
