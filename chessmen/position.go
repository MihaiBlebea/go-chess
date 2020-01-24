package chessmen

type Position struct {
	x int
	y int
}

func (p *Position) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

func (p *Position) GetX() int {
	return p.x
}

func (p *Position) GetY() int {
	return p.y
}