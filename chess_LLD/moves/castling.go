package moves

import (
	"lld/chess/models"
)

type Castling struct{}

func NewCastlingStrategy() *Castling {
	return &Castling{}
}

func (c *Castling) CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool {
	if board == nil || from.Piece == nil || to.Piece != nil {
		return false
	}

	if from.Piece.GetType() != "KING" || from.Row != to.Row || abs(to.Col-from.Col) != 2 {
		return false
	}

	if len(mod_piece) != 2 {
		return false
	}

	var king models.Piece
	var rook models.Piece

	for _, piece := range mod_piece {
		if piece == nil || *piece == nil {
			return false
		}

		switch (*piece).GetType() {
		case "KING":
			if king != nil {
				return false
			}
			king = *piece
		case "ROOK":
			if rook != nil {
				return false
			}
			rook = *piece
		default:
			return false
		}
	}

	if king == nil || rook == nil || king != from.Piece {
		return false
	}

	if king.GetMoveCnt() != 0 || rook.GetMoveCnt() != 0 || king.GetColorType() != rook.GetColorType() {
		return false
	}

	direction := 1
	rookCol := 7
	if to.Col < from.Col {
		direction = -1
		rookCol = 0
	}

	if board.Grid[from.Row][rookCol].Piece != rook {
		return false
	}

	for col := from.Col + direction; col != rookCol; col += direction {
		if board.Grid[from.Row][col].Piece != nil {
			return false
		}
	}

	return true
}
