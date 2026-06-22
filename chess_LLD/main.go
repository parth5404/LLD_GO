package main

import (
	"bufio"
	"fmt"
	"lld/chess/game"
	"lld/chess/moves"
	"lld/chess/pieces"
	"os"
	"strings"
)

func main() {
	moveFactory := moves.NewMoveFactory()
	pieceFactory := pieces.NewPieceFactory(moveFactory)
	g := game.GameInstance(pieceFactory)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Command Line Chess!")
	fmt.Println("Enter moves in the format 'e2 e4' or 'quit' to exit.")

	for {
		g.PrintBoard()
		fmt.Printf("%s's turn to move: ", g.GetCurrentTurn())

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "quit" || input == "exit" {
			fmt.Println("Thanks for playing!")
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please provide source and destination (e.g., 'e2 e4').")
			continue
		}

		success := g.PlayMoveStr(parts[0], parts[1])
		if !success {
			fmt.Println("Invalid move! Try again.")
		}
	}
}
