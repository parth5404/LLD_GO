package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

var _ models.Piece = (*Knight)(nil)

type Knight struct {
	color     models.Colour
	move_type []models.MoveType
	mf        *moves.MoveFactory
	mv_cnt    int
}

func NewKnight(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory, mv_cnt int) *Knight {
	return &Knight{
		move_type: setMoves,
		color:     clr,
		mf:        factory,
		mv_cnt:    mv_cnt,
	}
}

func (kn *Knight) ValidMove(b *models.Board, from *models.Square, to *models.Square, mod_piece []*models.Piece) bool {
	for _, m := range kn.move_type {
		if kn.mf.GetMoveStrategy(m).CanMove(*from, *to, b, mod_piece) {
			return true
		}
	}
	return false
}

func (kn *Knight) GetColorType() models.Colour {
	return kn.color
}

func (kn *Knight) GetType() string {
	return "Knight"
}

func (kn *Knight) GetMoveCnt() int {
	return kn.mv_cnt
}
