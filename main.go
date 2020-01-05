package main

import "fmt"

// main function
func main() {
	white := NewPlayer(0)
	black := NewPlayer(1)

	game := NewGame(white, black)
	game.PassTurn()
	game.PassTurn()

	// println(game.turn, game.WhosTurn().id, game.IsBlackTurn(), game.IsWhiteTurn())

	// pawn := NewPawn(1, 1)
	// err := game.Move(pawn, 2, 2)
	// if err != nil {
	// 	panic(err)
	// }

	for _, chessman := range game.white.chessmen {
		fmt.Println(chessman.GetId(), chessman.GetX(), chessman.GetY())
	}

	for _, chessman := range game.black.chessmen {
		fmt.Println(chessman.GetId(), chessman.GetX(), chessman.GetY())
	}
}
