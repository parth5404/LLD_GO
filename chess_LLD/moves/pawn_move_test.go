package moves

import (
	"lld/chess/models"
	"testing"
)

// pieceMock implements the full models.Piece interface for testing.
type pieceMock struct {
	color     models.Colour
	pieceType string
	mv_cnt    int
}

// Compile-time check: pieceMock must satisfy models.Piece
var _ models.Piece = (*pieceMock)(nil)

func (p *pieceMock) ValidMove(_ *models.Board, _ *models.Square, _ *models.Square, _ []*models.Piece) bool {
	return false
}
func (p *pieceMock) GetColorType() models.Colour { return p.color }
func (p *pieceMock) GetType() string             { return p.pieceType }
func (p *pieceMock) GetMoveCnt() int             { return p.mv_cnt }

// emptyBoard creates a fresh 8x8 models.Board with coordinates set and no pieces.
func emptyBoard() *models.Board {
	var b models.Board
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			b.Grid[r][c] = models.Square{Row: r, Col: c}
		}
	}
	return &b
}

// placePiece places a piece on the board at (row, col).
func placePiece(board *models.Board, row, col int, piece models.Piece) {
	board.Grid[row][col].Piece = piece
}

func TestPawnMove_AllCases(t *testing.T) {
	pawnMove := NewPawnMoveStrategy()

	tests := []struct {
		name        string
		fromRow     int
		fromCol     int
		toRow       int
		toCol       int
		fromColor   models.Colour
		fromMoveCnt int
		toOccupied  bool
		toColor     models.Colour // only used when toOccupied = true
		wantValid   bool
	}{
		// ================= Forward moves =================
		{
			name:    "White single forward to empty square",
			fromRow: 6, fromCol: 4, toRow: 5, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  true,
		},
		{
			name:    "Black single forward to empty square",
			fromRow: 1, fromCol: 4, toRow: 2, toCol: 4,
			fromColor: models.BLACK, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  true,
		},
		{
			name:    "White double forward from starting row (moveCnt=0)",
			fromRow: 6, fromCol: 4, toRow: 4, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  true,
		},
		{
			name:    "White double forward after moving (moveCnt>0) - invalid",
			fromRow: 5, fromCol: 4, toRow: 3, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 1,
			toOccupied: false,
			wantValid:  false,
		},
		{
			name:    "White forward blocked by own piece",
			fromRow: 6, fromCol: 4, toRow: 5, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: true, toColor: models.WHITE,
			wantValid: false,
		},
		{
			name:    "White forward to square occupied by opponent - invalid (can't capture forward)",
			fromRow: 6, fromCol: 4, toRow: 5, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: true, toColor: models.BLACK,
			wantValid: false,
		},
		{
			name:    "White double forward with piece in middle",
			fromRow: 6, fromCol: 4, toRow: 4, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			// blocker at (5,4) is placed in the setup loop below
			wantValid: false,
		},
		// ================= Diagonal captures =================
		{
			name:    "White diagonal capture opponent piece",
			fromRow: 6, fromCol: 4, toRow: 5, toCol: 3,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: true, toColor: models.BLACK,
			wantValid: true,
		},
		{
			name:    "White diagonal capture empty square - invalid",
			fromRow: 6, fromCol: 4, toRow: 5, toCol: 3,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  false,
		},
		{
			name:    "White diagonal capture own piece - invalid",
			fromRow: 6, fromCol: 4, toRow: 5, toCol: 3,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: true, toColor: models.WHITE,
			wantValid: false,
		},
		// ================= Direction & reverse =================
		{
			name:    "White backward move - invalid",
			fromRow: 5, fromCol: 4, toRow: 6, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  false,
		},
		{
			name:    "Black backward move - invalid",
			fromRow: 2, fromCol: 4, toRow: 1, toCol: 4,
			fromColor: models.BLACK, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  false,
		},
		{
			name:    "White moving more than 2 rows forward - invalid",
			fromRow: 6, fromCol: 4, toRow: 3, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  false,
		},
		{
			name:    "Same square - invalid",
			fromRow: 6, fromCol: 4, toRow: 6, toCol: 4,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false,
			wantValid:  false,
		},
		// ================= En Passant =================
		{
			name:    "En passant attempt (not implemented) - should be invalid",
			fromRow: 3, fromCol: 4, toRow: 2, toCol: 3,
			fromColor: models.WHITE, fromMoveCnt: 0,
			toOccupied: false, // en passant square is empty
			wantValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := emptyBoard()

			// Place source piece
			srcPiece := &pieceMock{color: tt.fromColor, mv_cnt: tt.fromMoveCnt, pieceType: "PAWN"}
			placePiece(board, tt.fromRow, tt.fromCol, srcPiece)

			// Place destination piece if occupied
			if tt.toOccupied {
				destPiece := &pieceMock{color: tt.toColor, pieceType: "PAWN"}
				placePiece(board, tt.toRow, tt.toCol, destPiece)
			}

			// Special setup: blocker in the middle for double-move test
			if tt.name == "White double forward with piece in middle" {
				blocker := &pieceMock{color: models.BLACK, pieceType: "PAWN"}
				placePiece(board, 5, 4, blocker)
			}

			from := board.Grid[tt.fromRow][tt.fromCol]
			to := board.Grid[tt.toRow][tt.toCol]

			// Call the REAL production code, not a duplicate
			got := pawnMove.CanMove(from, to, board, nil)
			if got != tt.wantValid {
				t.Errorf("CanMove() = %v, want %v", got, tt.wantValid)
			}
		})
	}
}
