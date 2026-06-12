package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

var _ models.Piece = (*Queen)(nil)

type Queen struct {
	color     models.Colour
	move_type []models.MoveType
	mf        *moves.MoveFactory
	mv_cnt    int
}

func NewQueen(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory, mv_cnt int) *Queen {
	return &Queen{
		move_type: setMoves,
		color:     clr,
		mf:        factory,
		mv_cnt:    mv_cnt,
	}
}

func (q *Queen) ValidMove(b *models.Board, from *models.Square, to *models.Square, mod_piece []*models.Piece) bool {
	for _, m := range q.move_type {
		if q.mf.GetMoveStrategy(m).CanMove(*from, *to, b, mod_piece) {
			return true
		}
	}
	return false
}

func (q *Queen) GetColorType() models.Colour {
	return q.color
}

func (q *Queen) GetType() string {
	return "QUEEN"
}

func (q *Queen) GetMoveCnt() int {
	return q.mv_cnt
}
