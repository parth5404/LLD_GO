package moves

import "lld/chess/models"

type KingMove struct{}

func NewKingMoveStrategy() *Horizontal {
	return &Horizontal{}
}

func (k *KingMove) CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool {
	if from == to {
		return false
	}
	if to.Piece != nil && to.Piece.GetColorType() == from.Piece.GetColorType() {
		return false
	}
	dr, dc := abs(to.Row-from.Row), abs(to.Col-from.Col)

	return (dr == 1 && dc == 1) || (dr == 0 && dc == 1) || (dr == 1 && dc == 0)
}
