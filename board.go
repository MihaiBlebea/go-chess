package main

import (
	"errors"
	"log"
)

//! Spot
type Spot struct {
	chessman Chessman
	position Position
}

func (s *Spot) HasChessman() bool {
	return s.chessman != nil
}

func (s *Spot) GetChessman() Chessman {
	return s.chessman
}

func (s *Spot) RemoveChessman() (bool, error) {
	s.chessman = nil
	return true, nil
}

func (s *Spot) AddChessman(chessman *Chessman) (bool, error) {
	// This is a capture if chessman is of different color
	if s.HasChessman() {
		return false, errors.New("There is already a chessman on this spot")
	}

	s.chessman = *chessman
	return true, nil
}


//! Prison
type Prison struct {
	captured []Chessman
}

func (p *Prison) AddChessman(chessman Chessman) {
	p.captured = append(p.captured, chessman)
}

func (p *Prison) GetCaptured() []Chessman {
	return p.captured
}


//! Board
type Board struct {
	spots [8][8]Spot
	whitePrison Prison
	blackPrison Prison
}

func (b *Board) Reset() {
	// Insert white 2nd line of chessmen
	b.spots[0][0] = Spot{Rook{White}, Position{0, 0}}
	b.spots[1][0] = Spot{Knight{White}, Position{1, 0}}
	b.spots[2][0] = Spot{Bishop{White}, Position{2, 0}}
	b.spots[3][0] = Spot{Queen{White}, Position{3, 0}}
	b.spots[4][0] = Spot{King{White}, Position{4, 0}}
	b.spots[5][0] = Spot{Bishop{White}, Position{5, 0}}
	b.spots[6][0] = Spot{Knight{White}, Position{6, 0}}
	b.spots[7][0] = Spot{Rook{White}, Position{7, 0}}

	for x := 0; x < 8; x++ {
		// Insert white line of pawns
		b.spots[x][1] = Spot{Pawn{White}, Position{x, 1}}

		// Insert black line of pawns
		b.spots[x][6] = Spot{Pawn{Black}, Position{x, 6}}
	}

	// Insert empty spots on the middle of the board
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

	// Insert black 2nd line of chessmen
	b.spots[0][7] = Spot{Rook{Black}, Position{0, 7}}
	b.spots[1][7] = Spot{Knight{Black}, Position{1, 7}}
	b.spots[2][7] = Spot{Bishop{Black}, Position{2, 7}}
	b.spots[3][7] = Spot{Queen{Black}, Position{3, 7}}
	b.spots[4][7] = Spot{King{Black}, Position{4, 7}}
	b.spots[5][7] = Spot{Bishop{Black}, Position{5, 7}}
	b.spots[6][7] = Spot{Knight{Black}, Position{6, 7}}
	b.spots[7][7] = Spot{Rook{Black}, Position{7, 7}}
}

func (b *Board) GetSpot(x, y int) *Spot {
	return &b.spots[x][y]
}

func (b *Board) Move(start, end Position) bool {
	startSpot := b.GetSpot(start.GetX(), start.GetY())
	endSpot := b.GetSpot(end.GetX(), end.GetY())

	// validate the move
	if b.validateMove(startSpot, endSpot) == false {
		return false
	}

	// Make the move
	startPositionChessman := startSpot.GetChessman()
	endPositionChessman := endSpot.GetChessman()

	// This is a capture
	if endSpot.HasChessman() {
		endSpot.RemoveChessman()
		_, err := endSpot.AddChessman(&startPositionChessman)
		if err != nil {
			log.Panic(err)
		}
		startSpot.RemoveChessman()

		if endPositionChessman.IsWhite() {
			b.blackPrison.AddChessman(endPositionChessman)
		} else {
			b.whitePrison.AddChessman(endPositionChessman)
		}
	} else {

		// This is a move
		//! this has to be an atomical operation like a transaction that can be rolled back if any errors
		startSpot.RemoveChessman()
		_, err := endSpot.AddChessman(&startPositionChessman)
		if err != nil {
			log.Panic(err)
		}
	}
	return true
}

func (b *Board) getPositionsBetween(start, end *Spot) []Position {
	diffX := end.position.GetX() - start.position.GetX()
	diffY := end.position.GetY() - start.position.GetY()

	patternX := 0
	patternY := 0

	if diffX > 0 {
		patternX = 1
	} 
	if diffX < 0 {
		patternX = -1
	}

	if diffY > 0 {
		patternY = 1
	} 
	if diffY < 0 {
		patternY = -1
	}
	
	var span []Position

	currentX := start.position.GetX()
	currentY := start.position.GetY()

	for currentX <= end.position.GetX() && currentY <= end.position.GetY() {

		currentPosition := Position{currentX, currentY}
		span = append(span, currentPosition)

		currentX += patternX
		currentY += patternY
	}

	// Strip the first and last positions from the slice as they are the start and end positions
	return span[1 : len(span)-1]
}

func (b *Board) validateMove(start, end *Spot) bool {
	isValid := true

	positionsBetween := b.getPositionsBetween(start, end)

	// Validate if there are any chessmen between start and end position
	for _, position := range(positionsBetween) {
		spot := b.GetSpot(position.GetX(), position.GetY())

		if spot.HasChessman() {
			isValid = false
		}
	}

	// Validate if the selected chessman can move to the end position
	chessman := start.GetChessman()
	canMove := chessman.CanMove(b, start, end)
	if canMove == false {
		isValid = false
	}

	return isValid
}

