package chessmen


type Bishop struct {
	id string
	position Position
	color Color
}

func (b Bishop) UpdateX(x int) {
	b.position.x = x
}

func (b Bishop) UpdateY(y int) {
	b.position.y = y
}

func (b Bishop) Equal(biship *Bishop) bool {
	return b.id == biship.id
}

func (b Bishop) GetId() string {
	return b.id
}

func (b Bishop) GetX() int {
	return b.position.x
}

func (b Bishop) GetY() int {
	return b.position.y
}

// GetValidMoves returns a slice of possible valid positions to move
func (b Bishop) GetValidMoves() []Position {
	var positions []Position
	// Go top right
	x := b.position.x
	y := b.position.y
	for x < 8 && y < 8 {
		x++
		y++
		positions = append(positions, Position{x, y})
	}

	// Go top left
	x = b.position.x
	y = b.position.y
	for x > 1 && y < 8 {
		x--
		y++
		positions = append(positions, Position{x, y})
	}

	// Go bottom right
	x = b.position.x
	y = b.position.y
	for x < 8 && y > 1 {
		x++
		y--
		positions = append(positions, Position{x, y})
	}

	// Go bottom left
	x = b.position.x
	y = b.position.y
	for x > 1 && y > 1 {
		x--
		y--
		positions = append(positions, Position{x, y})
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (b Bishop) GetValidTakes() []Position {
	return b.GetValidMoves()
}
