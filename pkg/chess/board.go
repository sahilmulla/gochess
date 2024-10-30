package chess

import (
	"bytes"
	"fmt"
)

type Board struct {
	Tiles []Tile
	Next  Color
}

func (b *Board) Move(from, to int) (Piece, error) {
	captured := b.Tiles[to].Piece

	b.Tiles[to] = b.Tiles[from]

	b.Tiles[from] = Tile{EmptyPiece}

	return captured, nil
}

func (b *Board) AvailableMoves(tileId int) map[int]Move {
	moves := make(map[int]Move)

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
				fmt.Println(currId)
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
				fmt.Println(currId)
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
	t := b.Tiles[idx]
	copy := Tile{Piece: t.Piece}

	return copy
}

func (b *Board) String() string {
	var buffer bytes.Buffer

	moves := b.AvailableMoves(22)
	for tileId, tile := range b.Tiles {
		if tileId%8 == 0 {
			fmt.Fprintf(&buffer, "%d\t", tileId/8+1)
		}

		if move, has := moves[tileId]; has {
			switch move {
			case Attack:
				fmt.Fprintf(&buffer, "\033[31m%s \033[39m", tile.Piece)
			case Advance:
				fmt.Fprintf(&buffer, "%s ", "+")
			}
		} else {
			fmt.Fprintf(&buffer, "%s ", tile.Piece)
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
