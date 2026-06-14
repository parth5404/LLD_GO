package game

import (
	"lld/chess/models"
	"lld/chess/pieces"
)

var standardStartingLayout = []string{
	"rnbqkbnr", // Row 0 (Black)
	"pppppppp", // Row 1 (Black)
	"........", // Row 2 (Empty)
	"........", // Row 3
	"........", // Row 4
	"........", // Row 5
	"PPPPPPPP", // Row 6 (White)
	"RNBQKBNR", // Row 7 (White)
}

func SetupBoardFromLayout(layout []string, pieceFactory *pieces.PieceFactory) *models.Board {
	var b models.Board

	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			b.Grid[r][c] = models.Square{Row: r, Col: c}

			char := layout[r][c]
			if char == '.' {
				continue // empty square
			}

			color := models.BLACK
			if char >= 'A' && char <= 'Z' {
				color = models.WHITE
			}

			switch char {
			case 'r', 'R':
				b.Grid[r][c].Piece = pieceFactory.CreatePiece("ROOK", color)
			case 'p', 'P':
				b.Grid[r][c].Piece = pieceFactory.CreatePiece("PAWN", color)
			case 'n', 'N':
				b.Grid[r][c].Piece = pieceFactory.CreatePiece("KNIGHT", color)
			case 'b', 'B':
				b.Grid[r][c].Piece = pieceFactory.CreatePiece("BISHOP", color)
			case 'q', 'Q':
				b.Grid[r][c].Piece = pieceFactory.CreatePiece("QUEEN", color)
			case 'k', 'K':
				b.Grid[r][c].Piece = pieceFactory.CreatePiece("KING", color)
			}
		}
	}
	return &b
}
