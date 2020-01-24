package chessmen


type Knight struct {
	id string
	position Position
	color Color
}

func (k Knight) UpdateX(x int) {
	k.position.x = x
}

func (k Knight) UpdateY(y int) {
	k.position.y = y
}

func (k Knight) Equal(knight *Knight) bool {
	return k.id == knight.id
}

func (k Knight) GetId() string {
	return k.id
}

func (k Knight) GetX() int {
	return k.position.x
}

func (k Knight) GetY() int {
	return k.position.y
}

// GetValidMoves returns a slice of possible valid positions to move
func (k Knight) GetValidMoves() []Position {
	var positions []Position
	
	// Go top right
	positions = appendPosition(positions, k.position.x + 1, k.position.y + 2)
	positions = appendPosition(positions, k.position.x + 2, k.position.y + 1)

	// Go top left
	positions = appendPosition(positions, k.position.x - 1, k.position.y + 2)
	positions = appendPosition(positions, k.position.x - 2, k.position.y + 1)

	// Go bottom right
	positions = appendPosition(positions, k.position.x + 1, k.position.y - 2)
	positions = appendPosition(positions, k.position.x + 2, k.position.y - 1)

	// Go bottom left
	positions = appendPosition(positions, k.position.x - 1, k.position.y - 2)
	positions = appendPosition(positions, k.position.x - 2, k.position.y - 1)

	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (k Knight) GetValidTakes() []Position {
	return k.GetValidMoves()
}

func appendPosition(positions []Position, x, y int) []Position {
	if x > 0 && x < 8 && y > 0 && y < 8 {
		positions = append(positions, Position{x, y})
	}
	return positions
}