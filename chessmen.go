package main

import "math"

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
	// Check if the end spot has a piece of the same color
	if end.GetChessman() != nil && end.GetChessman().IsWhite() == p.color.isWhite() {
		return false
	}

	// Check if it can move in the end position
	x := end.position.GetX() - start.position.GetX()
	y := end.position.GetY() - start.position.GetY()

	if x != 0 {
		return false
	}

	if start.position.GetY() < 4 {
		return y <= 2
	} else {
		return y == 1
	}
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
	if end.GetChessman() != nil && end.GetChessman().IsWhite() == r.color.isWhite() {
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
	if end.GetChessman() != nil && end.GetChessman().IsWhite() == k.color.isWhite() {
		return false
	}

	// Check if it can move in the end position
	x := end.position.GetX() - start.position.GetX()
	y := end.position.GetY() - start.position.GetY()
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
	if end.GetChessman() != nil && end.GetChessman().IsWhite() == b.color.isWhite() {
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
	if end.GetChessman() != nil && end.GetChessman().IsWhite() == k.color.isWhite() {
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
	if end.GetChessman() != nil && end.GetChessman().IsWhite() == q.color.isWhite() {
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
