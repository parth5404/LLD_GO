package pieces

import (
	"lld/chess/models"
	"lld/chess/moves"
)

var _ models.Piece = (*Pawn)(nil)

type Pawn struct {
	color     models.Colour
	move_type []models.MoveType
	mf        *moves.MoveFactory
	mv_cnt    int
}

func NewPawn(setMoves []models.MoveType, clr models.Colour, factory *moves.MoveFactory, mv_cnt int) *Pawn {
	return &Pawn{
		move_type: setMoves,
		color:     clr,
		mf:        factory,
		mv_cnt:    mv_cnt,
	}
}

func (p *Pawn) ValidMove(b *models.Board, from *models.Square, to *models.Square, mod_piece []*models.Piece) bool {
	for _, m := range p.move_type {
		if p.mf.GetMoveStrategy(m).CanMove(*from, *to, b, mod_piece) {
			return true
		}
	}
	return false
}

func (p *Pawn) GetColorType() models.Colour {
	return p.color
}

func (p *Pawn) GetType() string {
	return "Pawn"
}

func (p *Pawn) GetMoveCnt() int {
	return p.mv_cnt
}
