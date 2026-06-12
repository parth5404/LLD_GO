package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

var _ models.Piece = (*Rook)(nil)

type Rook struct {
	color     models.Colour
	move_type []models.MoveType
	mf        *moves.MoveFactory
	mv_cnt    int
}

func NewRook(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory, mv_cnt int) *Rook {
	return &Rook{
		move_type: setMoves,
		color:     clr,
		mf:        factory,
		mv_cnt:    mv_cnt,
	}
}

func (r *Rook) ValidMove(b *models.Board, from *models.Square, to *models.Square, mod_piece []*models.Piece) bool {
	for _, m := range r.move_type {
		if r.mf.GetMoveStrategy(m).CanMove(*from, *to, b, mod_piece) {
			return true
		}
	}
	return false
}

func (r *Rook) GetColorType() models.Colour {
	return r.color
}

func (r *Rook) GetType() string {
	return "ROOK"
}

func (r *Rook) GetMoveCnt() int {
	return r.mv_cnt
}
