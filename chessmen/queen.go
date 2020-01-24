package chessmen


type Queen struct {
	id string
	position Position
	color Color
}

func (q Queen) UpdateX(x int) {
	q.position.x = x
}

func (q Queen) UpdateY(y int) {
	q.position.y = y
}

func (q Queen) Equal(queen *Queen) bool {
	return q.id == queen.id
}

func (q Queen) GetId() string {
	return q.id
}

func (q Queen) GetX() int {
	return q.position.x
}

func (q Queen) GetY() int {
	return q.position.y
}

// GetValidMoves returns a slice of possible valid positions to move
func (q Queen) GetValidMoves() []Position {
	var positions []Position
	// Go top right
	x := q.position.x
	y := q.position.y
	for x < 8 && y < 8 {
		x++
		y++
		positions = append(positions, Position{x, y})
	}

	// Go top left
	x = q.position.x
	y = q.position.y
	for x > 1 && y < 8 {
		x--
		y++
		positions = append(positions, Position{x, y})
	}

	// Go bottom right
	x = q.position.x
	y = q.position.y
	for x < 8 && y > 1 {
		x++
		y--
		positions = append(positions, Position{x, y})
	}

	// Go bottom left
	x = q.position.x
	y = q.position.y
	for x > 1 && y > 1 {
		x--
		y--
		positions = append(positions, Position{x, y})
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (q Queen) GetValidTakes() []Position {
	return q.GetValidMoves()
}
