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
	board *Board
}

func Create(white, black string) *Game {
	return &Game{white, black, 0, StartGame, &Board{}}
}

func (g *Game) Play() {

	g.board.Reset()

	reader := bufio.NewReader(os.Stdin)

	for g.stage != GameOver {
		switch g.stage {
		case StartGame:
			fmt.Println("StartGame")
			g.stage = StartTurn
			
		case StartTurn:
			fmt.Printf("StartTurn %d", g.round)

			fmt.Println("")
			display := render(g.board)
			fmt.Print(display)

			g.stage = CheckChess

		case CheckChess:
			fmt.Println("CheckChess")
			g.stage = MoveChessman

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
			
			res := g.board.Move(Position{startX, startY}, Position{endX, endY})
			if res == true {
				g.stage = EndTurn
			}

			fmt.Println("Something went wrong, try again")

		case EndTurn:
			fmt.Println("EndTurn")

			g.round++

			g.stage = StartTurn
		}
	}
}
