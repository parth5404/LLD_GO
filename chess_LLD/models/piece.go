package models

type Board struct {
	Grid [8][8]Square
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
