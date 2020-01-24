package chessmen

type Rook struct {
	id string
	position Position
	color Color
}

func (r Rook) UpdateX(x int) {
	r.position.x = x
}

func (r Rook) UpdateY(y int) {
	r.position.y = y
}

func (r Rook) Equal(rook *Rook) bool {
	return r.id == rook.id
}

func (r Rook) GetId() string {
	return r.id
}

func (r Rook) GetX() int {
	return r.position.x
}

func (r Rook) GetY() int {
	return r.position.y
}

// GetValidMoves returns a slice of possible valid positions to move
func (r Rook) GetValidMoves() []Position {
	var positions []Position

	for y := 1; y < 9; y++ {
		if y != r.position.y {
			positions = append(positions, Position{r.position.x, y})
		}
	}

	for x := 1; x < 9; x++ {
		if x != r.position.x {
			positions = append(positions, Position{x, r.position.y})
		}
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (r Rook) GetValidTakes() []Position {
	return r.GetValidMoves()
}