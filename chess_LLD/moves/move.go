package moves

import "lld/chess/models"

type Move interface {
	CanMove(from models.Square, to models.Square, board *models.Board) bool
}

type MoveFactory struct {
	strategies map[models.MoveType]Move
}

func NewMoveFactory() *MoveFactory {
	return &MoveFactory{
		strategies: map[models.MoveType]Move{
			models.Horizontal: &Horizontal{},
		},
	}
}

func (m *MoveFactory) GetMoveStrategy(mt models.MoveType) Move {
	return m.strategies[mt]
}
