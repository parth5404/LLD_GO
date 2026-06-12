package moves

import (
	"lld/chess/models"
	"log"
)

type Diagonal struct{}

func NewDiagonalStrategy() *Diagonal {
	return &Diagonal{}
}

func (d *Diagonal) CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool {
	if from == to {
		log.Println("Same src and destination")
		return false
	}
	if (abs(to.Row-from.Row) != abs(to.Col-from.Col)) || to.Piece != nil && to.Piece.GetColorType() == from.Piece.GetColorType() {
		log.Println("Invalid Move")
		return false
	}
	i, j := from.Row, from.Col
	stepi, stepj := 1, 1
	if to.Row < from.Row {
		stepi = -1
	}
	if to.Col < from.Col {
		stepj = -1
	}
	i = i + stepi
	j = j + stepj
	for i != to.Row {
		if board.Grid[i][j].Piece != nil {
			return false
		}
		i += stepi
		j += stepj
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
