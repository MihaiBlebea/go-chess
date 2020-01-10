package main


type Bishop struct {
	*Chessman
}

func NewBishop(color Color, x, y int) *Bishop {
	return &Bishop{NewChessman(color, x, y)}
}

func (b *Bishop) UpdateX(x int) {
	b.x = x
}

func (b *Bishop) UpdateY(y int) {
	b.y = y
}

// GetValidMoves returns a slice of possible valid positions to move
func (b *Bishop) GetValidMoves() []Position {
	var positions []Position
	// Go top right
	x := b.x
	y := b.y
	for x < 8 && y < 8 {
		x++
		y++
		positions = append(positions, Position{x, y})
	}

	// Go top left
	x = b.x
	y = b.y
	for x > 1 && y < 8 {
		x--
		y++
		positions = append(positions, Position{x, y})
	}

	// Go bottom right
	x = b.x
	y = b.y
	for x < 8 && y > 1 {
		x++
		y--
		positions = append(positions, Position{x, y})
	}

	// Go bottom left
	x = b.x
	y = b.y
	for x > 1 && y > 1 {
		x--
		y--
		positions = append(positions, Position{x, y})
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (b *Bishop) GetValidTakes() []Position {
	return b.GetValidMoves()
}
