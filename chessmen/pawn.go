package chessmen


type Pawn struct {
	id string
	position Position
	color Color
}

func (p *Pawn) UpdateX(x int) {
	p.position.x = x
}

func (p Pawn) UpdateY(y int) {
	p.position.y = y
}

func (p Pawn) Equal(pawn Pawn) bool {
	return p.id == pawn.id
}

func (p Pawn) GetId() string {
	return p.id
}

func (p Pawn) GetX() int {
	return p.position.x
}

func (p Pawn) GetY() int {
	return p.position.y
}

// GetValidMoves returns a slice of possible valid positions to move
func (p Pawn) GetValidMoves() []Position {
	var positions []Position
	if p.color == White {
		if p.position.y < 5 {
			for y := p.position.y; y < 4; y++ {
				positions = append(positions, Position{p.position.x, y + 1})
			}
		} else {
			positions = append(positions, Position{p.position.x, p.position.y + 1})
		}
	} else {
		if p.position.y > 4 {
			for y := p.position.y; y > 5; y-- {
				positions = append(positions, Position{p.position.x, y - 1})
			}
		} else {
			positions = append(positions, Position{p.position.x, p.position.y - 1})
		}
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (p Pawn) GetValidTakes() []Position {
	var positions []Position
	if p.color == White {
		if p.position.y < 8 {
			if p.position.x > 1 {
				positions = append(positions, Position{p.position.x - 1, p.position.y + 1})
			}
			
			if p.position.x < 7 {
				positions = append(positions, Position{p.position.x + 1, p.position.y + 1})
			}
		}
	} else {
		if p.position.y > 1 {
			if p.position.x > 1 {
				positions = append(positions, Position{p.position.x - 1, p.position.y - 1})
			}
			
			if p.position.x < 7 {
				positions = append(positions, Position{p.position.x + 1, p.position.y - 1})
			}
		}
	}
	return positions
}