package moves

import "lld/chess/models"

type Move interface {
	CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool
}

type MoveFactory struct {
	strategies map[models.MoveType]Move
}

func NewMoveFactory() *MoveFactory {
	return &MoveFactory{
		strategies: map[models.MoveType]Move{
			models.Horizontal: &Horizontal{},
			models.Castling:   &Castling{},
			models.Pawn:       &PawnMove{},
			models.LShaped:    &Lshape{},
			models.Vertical:   &Vertical{},
			models.Diagonal:   &Diagonal{},
		},
	}
}

func (m *MoveFactory) GetMoveStrategy(mt models.MoveType) Move {
	return m.strategies[mt]
}
