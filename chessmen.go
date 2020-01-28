package main

import (
	"math"
	// "fmt"
)

//! Color
type Color int

const (
	White Color = iota
	Black
)

func (c *Color) isWhite() bool {
	return *c == White
}

//! Position
type Position struct {
	x int
	y int
}

func (p *Position) GetX() int {
	return p.x
}

func (p *Position) GetY() int {
	return p.y
}

//! Chessmen interface
type Chessman interface {
	CanMove(board *Board, start, end *Spot) bool
	IsWhite() bool
}

//! Pawn
type Pawn struct {
	color Color
}

func (p Pawn) IsWhite() bool {
	return p.color.isWhite()
}

func (p Pawn) CanMove(board *Board, start, end *Spot) bool {

	// Check the absolute difference between start and end positions
	x := math.Abs(float64(end.position.GetX() - start.position.GetX()))
	y := math.Abs(float64(end.position.GetY() - start.position.GetY()))

	// Check if this is a capture or a move
	// If the end spot has a chessman of a different color then yours
	if end.HasChessman() && end.GetChessman().IsWhite() != p.color.isWhite() {
		if x == 1 && y == 1 {
			return true
		}
	}

	// Check if the end spot has a chessman on it
	if end.HasChessman() {
		return false
	}

	if x != 0 {
		return false
	}

	startSpotChessman := start.GetChessman()

	//! Come up with a better rule here for Pawn movement
	if startSpotChessman.IsWhite() {
		if end.position.GetY() < start.position.GetY() {
			return false
		}

		if end.position.GetY() <= 3 {
			return y <= 2
		} else {
			return y == 1
		}
	} else {

		if end.position.GetY() > start.position.GetY() {
			return false
		}

		if end.position.GetY() >= 4 {
			return y <= 2
		} else {
			return y == 1
		}
	}
}

func getSpan(start, end *Spot) []Position {
	x, y := 0, 1
	var span []Position

	currentX := start.position.GetX()
	currentY := start.position.GetY()

	for currentX <= end.position.GetX() && currentY <= end.position.GetY() {

		currentPosition := Position{currentX, currentY}
		span = append(span, currentPosition)

		currentX += x
		currentY += y
	}

	// Strip the first and last positions from the slice as they are the start and end positions
	return span[1 : len(span)-1]
}

//! Rook
type Rook struct {
	color Color
}

func (r Rook) IsWhite() bool {
	return r.color.isWhite()
}

func (r Rook) CanMove(board *Board, start, end *Spot) bool {

	// Check if the end spot has a piece of the same color
	if end.HasChessman() {
		return false
	}

	// Check if it can move in the end position
	x := math.Abs(float64(end.position.GetX() - start.position.GetX()))
	y := math.Abs(float64(end.position.GetY() - start.position.GetY()))

	if (x == 0 && y <= 7) || (y == 0 && x <= 7) {
		return true
	}

	return false
}

//! Knight
type Knight struct {
	color Color
}

func (k Knight) IsWhite() bool {
	return k.color.isWhite()
}

func (k Knight) CanMove(board *Board, start, end *Spot) bool {

	// Check if the end spot has a piece of the same color
	if end.HasChessman() {
		return false
	}

	// Check if it can move in the end position
	x := math.Abs(float64(end.position.GetX() - start.position.GetX()))
	y := math.Abs(float64(end.position.GetY() - start.position.GetY()))

	return x*y == 2
}

//! Bishop
type Bishop struct {
	color Color
}

func (b Bishop) IsWhite() bool {
	return b.color.isWhite()
}

func (b Bishop) CanMove(board *Board, start, end *Spot) bool {

	// Check if the end spot has a piece of the same color
	if end.HasChessman() {
		return false
	}

	// Check if it can move in the end position
	x := end.position.GetX() - start.position.GetX()
	y := end.position.GetY() - start.position.GetY()
	return x == y
}

//! King
type King struct {
	color Color
}

func (k King) IsWhite() bool {
	return k.color.isWhite()
}

func (k King) CanMove(board *Board, start, end *Spot) bool {

	// Check if the end spot has a piece of the same color
	if end.HasChessman() {
		return false
	}

	// Check if it can move in the end position
	x := math.Abs(float64(end.position.GetX() - start.position.GetX()))
	y := math.Abs(float64(end.position.GetY() - start.position.GetY()))
	return x+y == 1
}

//! Queen
type Queen struct {
	color Color
}

func (q Queen) IsWhite() bool {
	return q.color.isWhite()
}

func (q Queen) CanMove(board *Board, start, end *Spot) bool {

	// Check if the end spot has a piece of the same color
	if end.HasChessman() {
		return false
	}

	// Check if it can move in the end position
	x := math.Abs(float64(end.position.GetX() - start.position.GetX()))
	y := math.Abs(float64(end.position.GetY() - start.position.GetY()))
	if x == y {
		return true
	}

	if (x == 0 && y <= 7) || (y == 0 && x <= 7) {
		return true
	}

	return false
}
