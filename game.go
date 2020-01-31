package main

import (
	"bufio"
	"fmt"
	"os"
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
	round       int
	stage       Stage
	board       *Board
}

func Create(white, black string) *Game {
	return &Game{white, black, 1, StartGame, &Board{}}
}

func (g *Game) IsWhiteTurn() bool {
	return g.round%2 != 0
}

func (g *Game) Play() {

	g.board.Reset()

	reader := bufio.NewReader(os.Stdin)

	cl := CommandLine{}

	for g.stage != GameOver {
		switch g.stage {
		case StartGame:
			fmt.Println("StartGame")
			g.stage = StartTurn

		case StartTurn:
			fmt.Println("StartTurn stage")
			if g.IsWhiteTurn() {
				fmt.Println("It is WHITE player turn")
			} else {
				fmt.Println("It is BLACK player turn")
			}

			// Display the board
			fmt.Print(cl.Render(g.board))

			g.stage = CheckChess

		case CheckChess:
			fmt.Println("CheckChess stage")
			fmt.Println("No chess found")

			g.stage = MoveChessman

		case MoveChessman:
			fmt.Println("MoveChessman stage")

			fmt.Print("Which chessman you want to move: ")
			start, _ := reader.ReadString('\n')

			// Validate the player input
			startX, startY, err := cl.TransformPosition(start)
			if err != nil {
				fmt.Println("Wrong move. Please try again")
				g.stage = MoveChessman
				continue
			}

			// Validate starting position
			startSpot := g.board.GetSpot(startX, startY)
			startChessman := startSpot.GetChessman()
			if startChessman != nil {
				if startChessman.IsWhite() != g.IsWhiteTurn() {
					fmt.Println("You have selected a wrong chessman. Please try again")
					g.stage = MoveChessman
					continue
				}
			}

			fmt.Print("Where do you want to move it: ")
			end, _ := reader.ReadString('\n')

			// Validate the player input
			endX, endY, err := cl.TransformPosition(end)
			if err != nil {
				fmt.Println("Wrong move. Please try again")
				g.stage = MoveChessman
				continue
			}

			res := g.board.Move(Position{startX, startY}, Position{endX, endY})
			if res == false {
				fmt.Println("Wrong move. Please try again")
				g.stage = MoveChessman
				continue
			}

			g.stage = EndTurn

		case EndTurn:
			fmt.Println("EndTurn stage")
			fmt.Println("-------------")
			fmt.Println("")

			g.round++

			g.stage = StartTurn

		case GameOver:
			fmt.Println("GameOver stage")
		}

	}
}
