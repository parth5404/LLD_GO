package models

type Board struct {
	Grid [8][8]Square // nil means empty square
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
	Castling
	Pawn
)

type Piece interface {
	ValidMove(b *Board, from *Square, to *Square) bool
	GetColorType() Colour
}

type Square struct {
	Row, Col int
	Piece    Piece
}

func NewSquare(row int, col int, pc Piece) *Square {
	return &Square{
		Row:   row,
		Col:   col,
		Piece: pc,
	}
}

func (s *Square) GetPiece() Piece {
	return s.Piece
}
