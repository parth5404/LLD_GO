package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

//var _ models.Piece = (*Rook)(nil)

type Rook struct {
	color  models.Colour
	move_type []models.MoveType
	mf     *moves.MoveFactory
}

func NewRook(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory) *Rook {
	return &Rook{
		move_type: setMoves,
		color:  clr,
		mf:     factory,
	}
}

func (r *Rook) ValidMove(b *models.Board, from *models.Square, to *models.Square) bool {
	for _, m := range r.move_type {
		if r.mf.GetMoveStrategy(m).CanMove(*from, *to, b) {
			return true
		}
	}
	return false
}

func (r *Rook) GetColorType() models.Colour {
	return r.color
}
