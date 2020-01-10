package main


type King struct {
	*Chessman
}

func NewKing(color Color, x, y int) *King {
	return &King{NewChessman(color, x, y)}
}

func (k *King) UpdateX(x int) {
	k.x = x
}

func (k *King) UpdateY(y int) {
	k.y = y
}

// GetValidMoves returns a slice of possible valid positions to move
func (k *King) GetValidMoves() []Position {
	var positions []Position
	
	// Go top
	positions = appendPosition(positions, k.x, k.y + 1)

	// Go bottom
	positions = appendPosition(positions, k.x, k.y - 1)

	// Go left
	positions = appendPosition(positions, k.x - 1, k.y)

	// Go right
	positions = appendPosition(positions, k.x + 1, k.y)

	// Go top right
	positions = appendPosition(positions, k.x + 1, k.y + 1)

	// Go top left
	positions = appendPosition(positions, k.x - 1, k.y + 1)

	// Go bottom right
	positions = appendPosition(positions, k.x + 1, k.y - 1)

	// Go bottom left
	positions = appendPosition(positions, k.x - 1, k.y - 1)

	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (k *King) GetValidTakes() []Position {
	return k.GetValidMoves()
}

