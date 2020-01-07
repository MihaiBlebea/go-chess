package main


type Pawn struct {
	*Chessman
}

func NewPawn(color Color, x, y int) *Pawn {
	return &Pawn{NewChessman(color, x, y)}
}

func (p *Pawn) UpdateX(x int) {
	p.x = x
}

func (p *Pawn) UpdateY(y int) {
	p.y = y
}

// GetValidMoves returns a slice of possible valid positions to move
func (p *Pawn) GetValidMoves() []Position {
	var positions []Position
	if p.color == white {
		if p.y < 5 {
			for y := p.y; y < 4; y++ {
				positions = append(positions, Position{p.x, y + 1})
			}
		} else {
			positions = append(positions, Position{p.x, p.y + 1})
		}
	} else {
		if p.y > 4 {
			for y := p.y; y > 5; y-- {
				positions = append(positions, Position{p.x, y - 1})
			}
		} else {
			positions = append(positions, Position{p.x, p.y - 1})
		}
	}
	return positions
}

// GetValidTakes returns a slice of possible valid positions to move and capture the enemy chessman
func (p *Pawn) GetValidTakes() []Position {
	var positions []Position
	if p.color == white {
		if p.y < 8 {
			if p.x > 1 {
				positions = append(positions, Position{p.x - 1, p.y + 1})
			}
			
			if p.x < 7 {
				positions = append(positions, Position{p.x + 1, p.y + 1})
			}
		}
	} else {
		if p.y > 1 {
			if p.x > 1 {
				positions = append(positions, Position{p.x - 1, p.y - 1})
			}
			
			if p.x < 7 {
				positions = append(positions, Position{p.x + 1, p.y - 1})
			}
		}
	}
	return positions
}