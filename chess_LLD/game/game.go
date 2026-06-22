package game

import (
	"fmt"
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

func (g *Game) checkmateCheck(clr models.Colour) bool {
	kingcheck := g.IsKingCheck(clr, nil)
	if !kingcheck {
		return false
	}
	grid := g.board.Grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].Piece == nil {
				continue
			}
			if grid[i][j].Piece.GetColorType() != clr {
				continue
			}
			if g.checkviaTempBoard(&grid[i][j], clr) {
				return false
			}

		}
	}
	return true
}

func (g *Game) checkviaTempBoard(from *models.Square, clr models.Colour) bool {
	grid := g.board.Grid

	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			to := grid[r][c]

			if from.Piece.ValidMove(g.board, from, &to, nil) {
				
				tempboard := g.board.Clone()
				
				tempboard.Grid[r][c].Piece = from.Piece
				tempboard.Grid[from.Row][from.Col].Piece = nil

				if !g.IsKingCheck(clr, tempboard) {
					return true
				}
			}
		}
	}
	
	
	return false
}

func (g *Game) GetCurrentTurn() string {
	if g.moveside == models.WHITE {
		return "White"
	}
	return "Black"
}

func (g *Game) PrintBoard() {
	fmt.Println("\n  a b c d e f g h")
	for r := 0; r < 8; r++ {
		fmt.Printf("%d ", 8-r)
		for c := 0; c < 8; c++ {
			p := g.board.Grid[r][c].Piece
			if p == nil {
				fmt.Print(". ")
			} else {
				char := "?"
				switch p.GetType() {
				case "PAWN":
					char = "p"
				case "ROOK":
					char = "r"
				case "KNIGHT":
					char = "n"
				case "BISHOP":
					char = "b"
				case "QUEEN":
					char = "q"
				case "KING":
					char = "k"
				}
				if p.GetColorType() == models.WHITE {
					// uppercase for white
					char = string(char[0] - 32)
				}
				fmt.Print(char + " ")
			}
		}
		fmt.Printf("%d\n", 8-r)
	}
	fmt.Println("  a b c d e f g h\n")
}

func parseCoord(s string) (int, int) {
	if len(s) != 2 {
		return -1, -1
	}
	c := int(s[0] - 'a')
	r := 8 - int(s[1]-'0')
	if r < 0 || r > 7 || c < 0 || c > 7 {
		return -1, -1
	}
	return r, c
}

// PlayMoveStr takes standard coordinates like "e2" to "e4".
func (g *Game) PlayMoveStr(fromStr, toStr string) bool {
	r1, c1 := parseCoord(fromStr)
	r2, c2 := parseCoord(toStr)

	if r1 == -1 || r2 == -1 {
		fmt.Println("Invalid coordinates. Use format 'e2'.")
		return false
	}

	fromSq := g.board.Grid[r1][c1]
	toSq := g.board.Grid[r2][c2]

	if fromSq.Piece == nil {
		fmt.Println("No piece at source square.")
		return false
	}

	movemap := make(map[string][]models.Square)
	
	// Special Case: Castling detected if King moves 2 steps horizontally
	if fromSq.Piece.GetType() == "KING" && (c2-c1 == 2 || c1-c2 == 2) {
		// Find the rook
		rRook := r1
		cRookFrom := 7
		cRookTo := 5
		if c2 < c1 {
			cRookFrom = 0
			cRookTo = 3
		}
		
		rookFromSq := g.board.Grid[rRook][cRookFrom]
		rookToSq := g.board.Grid[rRook][cRookTo]
		
		if rookFromSq.Piece == nil || rookFromSq.Piece.GetType() != "ROOK" {
			fmt.Println("Castling failed: Rook not found")
			return false
		}
		
		movemap["KING"] = []models.Square{fromSq, toSq}
		movemap["ROOK"] = []models.Square{rookFromSq, rookToSq}
		
	} else {
		// Normal move
		movemap[fromSq.Piece.GetType()] = []models.Square{fromSq, toSq}
	}

	return g.MakeMove(movemap)
}
