package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

var _ models.Piece = (*King)(nil)

type King struct {
	color     models.Colour
	move_type []models.MoveType
	mf        *moves.MoveFactory
	mv_cnt    int
}

func NewKing(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory, mv_cnt int) *King {
	return &King{
		move_type: setMoves,
		color:     clr,
		mf:        factory,
		mv_cnt:    mv_cnt,
	}
}

func (k *King) ValidMove(b *models.Board, from *models.Square, to *models.Square, mod_piece []*models.Piece) bool {
	for _, m := range k.move_type {
		if k.mf.GetMoveStrategy(m).CanMove(*from, *to, b, mod_piece) {
			return true
		}
	}
	return false
}

func (k *King) GetColorType() models.Colour {
	return k.color
}

func (k *King) GetType() string {
	return "KING"
}

func (k *King) GetMoveCnt() int {
	return k.mv_cnt
}
