package moves

import "lld/chess/models"

type Horizontal struct{}

func NewHorizontalStrategy() *Horizontal {
	return &Horizontal{}
}

func (*Horizontal) CanMove(from models.Square, to models.Square, board *models.Board) bool {
	if from.Row != to.Row || to.Piece.GetColorType() == from.Piece.GetColorType() {
		return false
	}
	i := from.Col
	step := 1
	if to.Col < from.Col {
		step = -1
	}
	i = i + step
	for i != to.Col {
		if board.Grid[from.Row][i].Piece != nil {
			return false
		}
		i += step
	}
	return true
}
