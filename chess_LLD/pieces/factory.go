package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

type PieceFactory struct {
	moveFactory *moves.MoveFactory
	moveRules   map[string][]models.MoveType
}

func NewPieceFactory(mf *moves.MoveFactory) *PieceFactory {
	return &PieceFactory{
		moveFactory: mf,
		moveRules: map[string][]models.MoveType{
			"PAWN":   {models.Pawn},
			"ROOK":   {models.Horizontal, models.Vertical},
			"KNIGHT": {models.LShaped},
			"KING":   {models.Castling, models.King},
			"QUEEN":  {models.Vertical, models.Horizontal, models.Diagonal},
			"BISHOP": {models.Diagonal},
		},
	}
}

func (pf *PieceFactory) CreatePiece(pieceType string, color models.Colour) models.Piece {
	rules := pf.moveRules[pieceType]

	switch pieceType {
	case "ROOK":
		return NewRook(rules, color, pf.moveFactory, 0)
	case "PAWN":
		return NewPawn(rules, color, pf.moveFactory, 0)
	case "KNIGHT":
		return NewKnight(rules, color, pf.moveFactory, 0)
	case "KING":
		return NewKing(rules, color, pf.moveFactory, 0)
	case "QUEEN":
		return NewQueen(rules, color, pf.moveFactory, 0)
	case "BISHOP":
		return NewBishop(rules, color, pf.moveFactory, 0)
	}
	return nil
}
