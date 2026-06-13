package moves

import (
	"lld/chess/models"
	"testing"
)

type castlingPiece struct {
	pieceType string
	color     models.Colour
	moveCount int
}

func (tp *castlingPiece) ValidMove(*models.Board, *models.Square, *models.Square, []*models.Piece) bool {
	return false
}

func (tp *castlingPiece) GetColorType() models.Colour {
	return tp.color
}

func (tp *castlingPiece) GetType() string {
	return tp.pieceType
}

func (tp *castlingPiece) GetMoveCnt() int {
	return tp.moveCount
}

func TestCastlingCanMoveKingSide(t *testing.T) {
	king := models.Piece(&castlingPiece{pieceType: "KING", color: models.WHITE})
	rook := models.Piece(&castlingPiece{pieceType: "ROOK", color: models.WHITE})
	board := &models.Board{}
	board.Grid[7][4] = *models.NewSquare(7, 4, king)
	board.Grid[7][6] = *models.NewSquare(7, 6, nil)
	board.Grid[7][7] = *models.NewSquare(7, 7, rook)

	modifiedPieces := []*models.Piece{&king, &rook}
	if !NewCastlingStrategy().CanMove(board.Grid[7][4], board.Grid[7][6], board, modifiedPieces) {
		t.Fatal("expected king-side castling to be valid")
	}
}

func TestCastlingCanMoveQueenSide(t *testing.T) {
	king := models.Piece(&castlingPiece{pieceType: "KING", color: models.BLACK})
	rook := models.Piece(&castlingPiece{pieceType: "ROOK", color: models.BLACK})
	board := &models.Board{}
	board.Grid[0][4] = *models.NewSquare(0, 4, king)
	board.Grid[0][2] = *models.NewSquare(0, 2, nil)
	board.Grid[0][0] = *models.NewSquare(0, 0, rook)

	modifiedPieces := []*models.Piece{&rook, &king}
	if !NewCastlingStrategy().CanMove(board.Grid[0][4], board.Grid[0][2], board, modifiedPieces) {
		t.Fatal("expected queen-side castling to be valid")
	}
}

func TestCastlingRejectsBlockedPath(t *testing.T) {
	king := models.Piece(&castlingPiece{pieceType: "KING", color: models.WHITE})
	rook := models.Piece(&castlingPiece{pieceType: "ROOK", color: models.WHITE})
	blocker := models.Piece(&castlingPiece{pieceType: "BISHOP", color: models.WHITE})
	board := &models.Board{}
	board.Grid[7][4] = *models.NewSquare(7, 4, king)
	board.Grid[7][5] = *models.NewSquare(7, 5, blocker)
	board.Grid[7][6] = *models.NewSquare(7, 6, nil)
	board.Grid[7][7] = *models.NewSquare(7, 7, rook)

	modifiedPieces := []*models.Piece{&king, &rook}
	if NewCastlingStrategy().CanMove(board.Grid[7][4], board.Grid[7][6], board, modifiedPieces) {
		t.Fatal("expected a blocked castling path to be invalid")
	}
}
