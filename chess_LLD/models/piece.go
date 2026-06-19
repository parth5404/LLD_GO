package models

type Board struct {
	Grid [8][8]Square
}

func (b *Board) Clone() *Board {
	var nb Board

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			nb.Grid[i][j] = b.Grid[i][j]
		}
	}

	return &nb
}

type Colour int

type MoveType int

const (
	WHITE Colour = iota
	BLACK
)

const (
	LShaped MoveType = iota
	Diagonal
	Horizontal
	Vertical
	Castling
	Pawn
	King
)

type Piece interface {
	ValidMove(b *Board, from *Square, to *Square, mod_piece []*Piece) bool
	GetColorType() Colour
	GetType() string
	GetMoveCnt() int
}
