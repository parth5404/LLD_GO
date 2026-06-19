package main

import (
	"lld/chess/game"
	"lld/chess/moves"
	"lld/chess/pieces"
)

func main() {
	moveFactory := moves.NewMoveFactory()
	pieceFactory := pieces.NewPieceFactory(moveFactory)
	game := game.GameInstance(pieceFactory)

}
