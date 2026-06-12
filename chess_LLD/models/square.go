package models

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
