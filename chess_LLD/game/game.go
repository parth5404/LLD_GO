package game

import (
	"lld/chess/models"
	"lld/chess/pieces"
)

type Game struct {
	whiteKingLoc *models.Square
	blackKingLoc *models.Square
	board        *models.Board
	pieceFactory *pieces.PieceFactory
	moveside     models.Colour
}

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
var ins *Game = nil

func GameInstance(factory *pieces.PieceFactory) *Game {
	if ins == nil {
		ins = &Game{pieceFactory: factory}
		ins.moveside = models.WHITE
		ins.setupGame(standardStartingLayout, ins.pieceFactory)
	}
	return ins
}
func (g *Game) setupGame(layout []string, pieceFactory *pieces.PieceFactory) {
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
				switch char {
				case 'K':
					g.whiteKingLoc = &b.Grid[r][c]
				case 'k':
					g.blackKingLoc = &b.Grid[r][c]

				}
			}
		}
	}
	g.board = &b

}

func (g *Game) MakeMove(movemap map[string][]models.Square) bool {
	tempboard := g.board.Clone()
	kingcheck := g.IsKingCheck(g.moveside, tempboard)
	// CASTLING
	if len(movemap) == 2 && !kingcheck {
		k1, ok1 := movemap["KING"]
		r1, ok2 := movemap["ROOK"]
		if !ok1 || !ok2 {
			return false
		}
		if len(k1) != 2 || len(r1) != 2 {
			return false
		}

		king := k1[0].Piece
		rookPiece := r1[0].Piece

		var pc []*models.Piece
		pc = append(pc, &rookPiece)

		if king.ValidMove(tempboard, &k1[0], &k1[1], pc) {

			tempboard.Grid[k1[1].Row][k1[1].Col].Piece = king
			tempboard.Grid[k1[0].Row][k1[0].Col].Piece = nil

			tempboard.Grid[r1[1].Row][r1[1].Col].Piece = rookPiece
			tempboard.Grid[r1[0].Row][r1[0].Col].Piece = nil

			if g.IsKingCheck(g.moveside, tempboard) {
				return false
			}
			g.board = tempboard

			if king.GetColorType() == models.WHITE {
				g.whiteKingLoc = &g.board.Grid[k1[1].Row][k1[1].Col]
			} else {
				g.blackKingLoc = &g.board.Grid[k1[1].Row][k1[1].Col]
			}
			if g.moveside == models.WHITE {
				g.moveside = models.BLACK
			} else {
				g.moveside = models.WHITE
			}
			return true
		}
		return false
	}

	// Normal Move Case: 1 piece involved
	for _, v := range movemap {
		if len(v) != 2 {
			return false
		}
		pc := v[0].Piece
		if pc == nil {
			return false
		}

		if pc.ValidMove(g.board, &v[0], &v[1], nil) {

			tempboard.Grid[v[1].Row][v[1].Col].Piece = pc
			tempboard.Grid[v[0].Row][v[0].Col].Piece = nil
			if g.IsKingCheck(g.moveside, tempboard) {
				return false
			}
			g.board = tempboard
			if pc.GetType() == "KING" {
				if pc.GetColorType() == models.WHITE {
					g.whiteKingLoc = &g.board.Grid[v[1].Row][v[1].Col]
				} else {
					g.blackKingLoc = &g.board.Grid[v[1].Row][v[1].Col]
				}
			}
			if g.moveside == models.WHITE {
				g.moveside = models.BLACK
			} else {
				g.moveside = models.WHITE
			}
			return true
		}
	}
	return false
}
func (g *Game) findKing(clr models.Colour, tempboard *models.Board) *models.Square {
	grid := g.board.Grid
	if tempboard != nil {
		grid = tempboard.Grid
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].Piece == nil {
				continue
			}
			if grid[i][j].Piece.GetColorType() == clr && grid[i][j].Piece.GetType() == "KING" {
				return &grid[i][j]
			}

		}
	}
	return nil
}

func (g *Game) IsKingCheck(clr models.Colour, tempboard *models.Board) bool {
	board := g.board
	grid := g.board.Grid
	if tempboard != nil {
		board = tempboard
		grid = board.Grid
	}
	king := g.findKing(clr, board)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].Piece == nil {
				continue
			}
			if grid[i][j].Piece.GetColorType() == clr {
				continue
			}

			if grid[i][j].Piece.ValidMove(board, &grid[i][j], king, nil) == true {
				return true
			}

		}
	}
	return false
}
