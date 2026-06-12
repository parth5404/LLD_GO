package moves

import (
	"lld/chess/models"
	"log"
)

type Lshape struct{}

func NewLshapeStrategy() *Lshape {
	return &Lshape{}
}

func (l *Lshape) CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool {
	if from == to {
		log.Println("Same src and destination")
		return false
	}
	if from.Row == to.Row || from.Col == to.Col {
		log.Println("Invalid Move")
		return false
	}
	if (abs(to.Row-from.Row)+abs(to.Col-from.Col) != 3) || to.Piece != nil && to.Piece.GetColorType() == from.Piece.GetColorType() {
		log.Println("Invalid Move")
		return false
	}

	return true
}
