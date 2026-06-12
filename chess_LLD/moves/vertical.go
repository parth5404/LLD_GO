package moves

import (
	"lld/chess/models"
	"log"
)

type Vertical struct{}

func NewVerticalStrategy() *Vertical {
	return &Vertical{}
}

func (v *Vertical) CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool {
	if from == to {
		log.Println("Same src and destination")
		return false
	}
	if from.Col != to.Col || (to.Piece != nil && to.Piece.GetColorType() == from.Piece.GetColorType()) {
		log.Println("Invalid Move")
		return false
	}
	i := from.Row
	step := 1
	if to.Row < from.Row {
		step = -1
	}
	i = i + step
	for i != to.Row {
		if board.Grid[i][from.Col].Piece != nil {
			return false
		}
		i += step
	}
	return true
}
