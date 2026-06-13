package moves

import (
	"lld/chess/models"
	"log"
)

type PawnMove struct{}

func NewPawnMoveStrategy() *PawnMove {
	return &PawnMove{}
}

func (pm *PawnMove) CanMove(from models.Square, to models.Square, board *models.Board, mod_piece []*models.Piece) bool {
	if from == to {
		log.Println("Same src and destination")
		return false
	}
	if to.Piece != nil && to.Piece.GetColorType() == from.Piece.GetColorType() {
		log.Println("Invalid Move")
		return false
	}
	pc := from.GetPiece()
	dr := to.Row - from.Row
	dc := to.Col - from.Col
	step := 1
	if pc.GetColorType() == models.WHITE {
		step = -1
	}

	if abs(dc) == 1 {
		if dr == step {
			if to.GetPiece() != nil && to.GetPiece().GetColorType() != pc.GetColorType() {
				return true
			}
		}
	} else {
		if dr != step && dr != 2*step {
			return false
		}
		if pc.GetMoveCnt() != 0 && abs(dr) == 2 {
			return false
		}

		for i := from.Row + step; i != to.Row; i = i + step {
			if board.Grid[i][to.Col].Piece != nil {
				return false
			}
		}
		if to.Piece != nil {
			return false
		}
		return true
	}
	return false
}
