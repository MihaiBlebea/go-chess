package main


//! Spot
type Spot struct {
	chessman Chessman
	position Position
}

func (s *Spot) GetChessman() Chessman {
	return s.chessman
}


//! Board 
type Board struct {
	spots [8][8]Spot
}

func (b *Board) Reset() {
	b.spots[0][0] = Spot{Rook{ White }, Position{0, 0}}
	b.spots[1][0] = Spot{Knight{ White }, Position{1, 0}}
	b.spots[2][0] = Spot{Bishop{ White }, Position{2, 0}}
	b.spots[3][0] = Spot{nil, Position{3, 0}}
	b.spots[4][0] = Spot{nil, Position{4, 0}}
	b.spots[5][0] = Spot{Bishop{ White }, Position{5, 0}}
	b.spots[6][0] = Spot{Knight{ White }, Position{6, 0}}
	b.spots[7][0] = Spot{Rook{ White }, Position{7, 0}}

	for x := 0; x < 8; x++ {
		b.spots[x][1] = Spot{Pawn{ White }, Position{x, 1}}
		b.spots[x][6] = Spot{Pawn{ Black }, Position{x, 6}}
	}

	for y := 2; y < 6; y++ {
		b.spots[0][y] = Spot{nil, Position{0, y}}
		b.spots[1][y] = Spot{nil, Position{1, y}}
		b.spots[2][y] = Spot{nil, Position{2, y}}
		b.spots[3][y] = Spot{nil, Position{3, y}}
		b.spots[4][y] = Spot{nil, Position{4, y}}
		b.spots[5][y] = Spot{nil, Position{5, y}}
		b.spots[6][y] = Spot{nil, Position{6, y}}
		b.spots[7][y] = Spot{nil, Position{7, y}}
	}

	b.spots[0][7] = Spot{Rook{ Black }, Position{0, 7}}
	b.spots[1][7] = Spot{Knight{ Black }, Position{1, 7}}
	b.spots[2][7] = Spot{Bishop{ Black }, Position{2, 7}}
	b.spots[3][7] = Spot{nil, Position{3, 7}}
	b.spots[4][7] = Spot{nil, Position{4, 7}}
	b.spots[5][7] = Spot{Bishop{ Black }, Position{5, 7}}
	b.spots[6][7] = Spot{Knight{ Black }, Position{6, 7}}
	b.spots[7][7] = Spot{Rook{ Black }, Position{7, 7}}
}

func (b *Board) GetSpot(x, y int) Spot {
	return b.spots[x][y]
}

