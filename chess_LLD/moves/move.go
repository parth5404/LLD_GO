package moves

import "lld/chess/models"

type MoveFactory struct{}

type Move interface {
	CanMove(from models.Square, to models.Square, board *models.Board) bool
}

func NewMoveFactory() *MoveFactory {
	return &MoveFactory{}
}

func (m *MoveFactory) GetMoveStrategy(moveType models.MoveType) Move {
	if moveType == models.Horizontal {
		return NewHorizontalStrategy()
	}
	return nil
}
