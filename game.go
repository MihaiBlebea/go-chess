package main

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
	white Player
	black Player
	round int
	stage Stage
	board *Board
}

func NewGame(white, black Player) *Game {
	return &Game{white, black, 1, StartGame, &Board{}}
}

func (g *Game) IsWhiteTurn() bool {
	return g.round%2 != 0
}

func (g *Game) Play() {

	g.board.Reset()

	for g.stage != GameOver {
		switch g.stage {
		case StartGame:
			g.stage = StartTurn

		case StartTurn:
			player := g.playerTurn()
			player.Alert("StartTurn stage")

			// Display the board
			player.Render(g.board)

			g.stage = CheckChess

		case CheckChess:
			player := g.playerTurn()
			player.Alert("CheckChess stage")
			player.Alert("No chess found")

			g.stage = MoveChessman

		case MoveChessman:
			player := g.playerTurn()

			player.Alert("MoveChessman stage")

			start, end := player.Move()

			res := g.board.Move(start, end)
			if res == false {
				player.Alert("Something went wrong")
				g.stage = MoveChessman
				continue
			}

			g.stage = EndTurn

		case EndTurn:
			player := g.playerTurn()
			player.Alert("EndTurn stage")

			g.round++

			g.stage = StartTurn

		case GameOver:
			//
		}

	}
}

func (g *Game) playerTurn() Player {
	if g.IsWhiteTurn() {
		return g.white
	} else {
		return g.black
	}
}
