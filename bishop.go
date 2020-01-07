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
	for y := b.y; y < 9; y++ {
		positions = append(positions, Position{b.x + 1, b.y + 1})
	}

	for x := b.y; x > 1; x-- {
		positions = append(positions, Position{b.x - 1, b.y - 1})
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (b *Bishop) GetValidTakes() []Position {
	return b.GetValidMoves()
}