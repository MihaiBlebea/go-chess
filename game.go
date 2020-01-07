package main

type Game struct {
	turn  int
	white *Player
	black *Player
}

func NewGame(white, black *Player) *Game {
	return &Game{0, white, black}
}

func (g *Game) PassTurn() {
	g.turn++
}

func (g *Game) WhosTurn() *Player {
	if g.turn%2 == 0 {
		return g.white
	}
	return g.black
}

func (g *Game) IsWhiteTurn() bool {
	return g.WhosTurn().id == g.white.id
}

func (g *Game) IsBlackTurn() bool {
	return g.WhosTurn().id == g.black.id
}

func (g *Game) Move(chessman IChessman, newX, newY int) error {
	// See if chessman can move onto the new position
	if g.validateNewPosition(chessman, newX, newY) == false {
		// return error
	}

	// See if there is no other chessman in the path of the move


	return nil
}

// validateNewPosition of the chessman, return true if the new position is valid from the chessman point of view
func (g *Game) validateNewPosition(chessman IChessman, x, y int) bool {
	// validXs := chessman.GetValidXMoves()
	// validYs := chessman.GetValidYMoves()

	// for _, validY := range validYs {
	// 	for _, validX := range validXs {
	// 		if validX == x && validY == y {
	// 			return true
	// 		}
	// 	}
	// }
	return false
}

func (g *Game) validatePath(chessman IChessman, x, y int) bool {
	return true
}
