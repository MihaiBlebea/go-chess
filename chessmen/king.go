package chessmen


type King struct {
	id string
	position Position
	color Color
}

func (k King) UpdateX(x int) {
	k.position.x = x
}

func (k King) UpdateY(y int) {
	k.position.y = y
}

func (k King) Equal(king *King) bool {
	return k.id == king.id
}

func (k King) GetId() string {
	return k.id
}

func (k King) GetX() int {
	return k.position.x
}

func (k King) GetY() int {
	return k.position.y
}

// GetValidMoves returns a slice of possible valid positions to move
func (k King) GetValidMoves() []Position {
	var positions []Position
	
	// Go top
	positions = appendPosition(positions, k.position.x, k.position.y + 1)

	// Go bottom
	positions = appendPosition(positions, k.position.x, k.position.y - 1)

	// Go left
	positions = appendPosition(positions, k.position.x - 1, k.position.y)

	// Go right
	positions = appendPosition(positions, k.position.x + 1, k.position.y)

	// Go top right
	positions = appendPosition(positions, k.position.x + 1, k.position.y + 1)

	// Go top left
	positions = appendPosition(positions, k.position.x - 1, k.position.y + 1)

	// Go bottom right
	positions = appendPosition(positions, k.position.x + 1, k.position.y - 1)

	// Go bottom left
	positions = appendPosition(positions, k.position.x - 1, k.position.y - 1)

	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (k King) GetValidTakes() []Position {
	return k.GetValidMoves()
}

