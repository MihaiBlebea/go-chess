package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

type Stage int

const (
	StartGame Stage = iota
	StartTurn
	CheckChess
	MoveChessman
	EndTurn
	GameOver
)

type Game struct {
	playerWhite string
	playerBlack string
	round int
	stage Stage
}

func NewGame() *Game {

	game := Game{"Mihai", "Serban", 0, StartGame}

	board := Board{}
	board.Reset()

	reader := bufio.NewReader(os.Stdin)

	for game.stage != GameOver {
		switch game.stage {
		case StartGame:
			fmt.Println("StartGame")
			game.stage = StartTurn
			
		case StartTurn:
			fmt.Printf("StartTurn %d", game.round)

			fmt.Println("")
			display := render(&board)
			fmt.Print(display)

			game.stage = CheckChess

		case CheckChess:
			fmt.Println("CheckChess")
			game.stage = MoveChessman

		case MoveChessman:
			fmt.Println("MoveChessman")

			fmt.Print("Which chessman you want to move: ")
			start, _ := reader.ReadString('\n')

			// Validate the player input
			startX, startY, err := TransformPosition(start)
			if err != nil {
				//
			}

			fmt.Print("Where do you want to move it: ")
			end, _ := reader.ReadString('\n')

			// Validate the player input
			endX, endY, err := TransformPosition(end)
			if err != nil {
				log.Panic(err)
			}
			
			res := board.Move(Position{startX, startY}, Position{endX, endY})
			if res == true {
				game.stage = EndTurn
			}

			fmt.Println("Something went wrong, try again")

		case EndTurn:
			fmt.Println("EndTurn")

			game.round++

			game.stage = StartTurn
		}
	}
	return &game
}
