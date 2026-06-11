package main

import (
	"fmt"
	"lld/chess/models"
	"lld/chess/moves"
	"lld/chess/pieces"
)

func main() {
	movefac := moves.NewMoveFactory()
	var mt []models.MoveType = make([]models.MoveType, 0)
	mt = append(mt, models.Horizontal)
	rook := pieces.NewRook(mt[0:], models.BLACK, movefac)
	fmt.Println(rook.GetColorType())

}
