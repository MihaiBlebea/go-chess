package main


type Knight struct {
	*Chessman
}

func NewKnight(color Color, x, y int) *Knight {
	return &Knight{NewChessman(color, x, y)}
}

func (k *Knight) UpdateX(x int) {
	k.x = x
}

func (k *Knight) UpdateY(y int) {
	k.y = y
}

// GetValidMoves returns a slice of possible valid positions to move
func (k *Knight) GetValidMoves() []Position {
	var positions []Position
	
	// Go top right
	positions = appendPosition(positions, k.x + 1, k.y + 2)
	positions = appendPosition(positions, k.x + 2, k.y + 1)

	// Go top left
	positions = appendPosition(positions, k.x - 1, k.y + 2)
	positions = appendPosition(positions, k.x - 2, k.y + 1)

	// Go bottom right
	positions = appendPosition(positions, k.x + 1, k.y - 2)
	positions = appendPosition(positions, k.x + 2, k.y - 1)

	// Go bottom left
	positions = appendPosition(positions, k.x - 1, k.y - 2)
	positions = appendPosition(positions, k.x - 2, k.y - 1)

	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (k *Knight) GetValidTakes() []Position {
	return k.GetValidMoves()
}

func appendPosition(positions []Position, x, y int) []Position {
	if x > 0 && x < 8 && y > 0 && y < 8 {
		positions = append(positions, Position{x, y})
	}
	return positions
}