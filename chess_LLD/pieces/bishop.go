package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

var _ models.Piece = (*Bishop)(nil)

type Bishop struct {
	color     models.Colour
	move_type []models.MoveType
	mf        *moves.MoveFactory
	mv_cnt    int
}

func NewBishop(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory, mv_cnt int) *Bishop {
	return &Bishop{
		move_type: setMoves,
		color:     clr,
		mf:        factory,
		mv_cnt:    mv_cnt,
	}
}

func (bi *Bishop) ValidMove(b *models.Board, from *models.Square, to *models.Square, mod_piece []*models.Piece) bool {
	for _, m := range bi.move_type {
		if bi.mf.GetMoveStrategy(m).CanMove(*from, *to, b, mod_piece) {
			return true
		}
	}
	return false
}

func (bi *Bishop) GetColorType() models.Colour {
	return bi.color
}

func (bi *Bishop) GetType() string {
	return "BISHOP"
}

func (bi *Bishop) GetMoveCnt() int {
	return bi.mv_cnt
}
