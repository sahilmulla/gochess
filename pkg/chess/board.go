package chess

import (
	"bytes"
	"errors"
	"fmt"
)

type Board struct {
	Tiles []Tile
	Next  Color
}

func (b *Board) Move(from, to int) (*Piece, error) {
	moves := b.AvailableMoves(from)

	if _, has := moves[to]; !has {
		return nil, errors.New("illegal move")
	}

	captured := b.Tiles[to].Piece

	b.Tiles[to] = b.Tiles[from]

	b.Tiles[from] = Tile{EmptyPiece}

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
				if other := b.TileAt(currId).Piece.Color(); other != None && other != tile.Piece.Color() {
					moves[currId] = Attack
					break
				}
				if b.TileAt(currId).Piece.Color() == tile.Piece.Color() {
					break
				}
				moves[currId] = Advance
			}
		}
	}

	if tile.Piece == WhiteBishop || tile.Piece == BlackBishop {
		for _, vec := range []Vector{N + E, N + W, S + E, S + W} {
			for currId := tileId + int(vec); ; currId += int(vec) {
				if currId < 0 || currId >= NumberOfTiles || ((vec == N+E || vec == S+E) && currId%8 == 0) || ((vec == N+W || vec == S+W) && currId%8 == 7) {
					break
				}
				if other := b.TileAt(currId).Piece.Color(); other != None && other != tile.Piece.Color() {
					moves[currId] = Attack
					break
				}
				if b.TileAt(currId).Piece.Color() == tile.Piece.Color() {
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
				if other := b.TileAt(currId).Piece.Color(); other != None && other != tile.Piece.Color() {
					moves[currId] = Attack
					break
				}
				if b.TileAt(currId).Piece.Color() == tile.Piece.Color() {
					break
				}
				moves[currId] = Advance

				if mod := currId % 8; mod == 0 || mod == 7 {
					break
				}
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

	moves := b.AvailableMoves(activeId)
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
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[31m %s \033[0m", tile.Piece)))
			case Advance:
				buffer.WriteString(checkerIt(fmt.Sprintf("\033[34m %s \033[0m", "+")))
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
		Tiles: make([]Tile, NumberOfTiles),
		Next:  White,
	}

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

func WithStandardStartPosition() BoardOption {
	return WithCustomStartPosition(StandardStart)
}

func WithCustomStartPosition(initPieces map[int]Piece) BoardOption {
	return func(b *Board) {
		for pos, piece := range initPieces {
			b.Tiles[pos] = Tile{Piece: piece}
		}
	}
}

type Tile struct {
	Piece Piece
}
