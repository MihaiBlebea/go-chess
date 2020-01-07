package main

type Rook struct {
	*Chessman
}

func NewRook(color Color, x, y int) *Rook {
	return &Rook{NewChessman(color, x, y)}
}

func (r *Rook) UpdateX(x int) {
	r.x = x
}

func (r *Rook) UpdateY(y int) {
	r.y = y
}

// GetValidMoves returns a slice of possible valid positions to move
func (r *Rook) GetValidMoves() []Position {
	var positions []Position

	for y := 1; y < 9; y++ {
		if y != r.y {
			positions = append(positions, Position{r.x, y})
		}
	}

	for x := 1; x < 9; x++ {
		if x != r.x {
			positions = append(positions, Position{x, r.y})
		}
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (r *Rook) GetValidTakes() []Position {
	return r.GetValidMoves()
}