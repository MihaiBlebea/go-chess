package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Chessman
type Chessman struct {
	id string
	x  int
	y  int
}

type IChessman interface {
	MoveChessman
	GetId() string
	GetX() int
	GetY() int
	equal(chessman *Chessman) bool
}

func NewChessman(x, y int) *Chessman {
	uuid := uuid.New()
	id := uuid.ID()

	return &Chessman{
		id: fmt.Sprint(id),
		x: x,
		y: y,
	}
}

func (c *Chessman) equal(chessman *Chessman) bool {
	return c.id == chessman.id
}

func (c *Chessman) GetId() string {
	return c.id
}

func (c *Chessman) GetX() int {
	return c.x
}

func (c *Chessman) GetY() int {
	return c.y
}

type Pawn struct {
	*Chessman
}

type Rook struct {
	*Chessman
}

type Bishop struct {
	*Chessman
}

type Knight struct {
	*Chessman
}

type King struct {
	*Chessman
}

type Queen struct {
	*Chessman
}

type MoveChessman interface {
	GetValidXMoves() []int
	GetValidYMoves() []int
}

func NewPawn(x, y int) *Pawn {
	return &Pawn{NewChessman(x, y)}
}

func NewRook(x, y int) *Pawn {
	return &Pawn{NewChessman(x, y)}
}

func NewBishop(x, y int) *Pawn {
	return &Pawn{NewChessman(x, y)}
}

func NewKnight(x, y int) *Pawn {
	return &Pawn{NewChessman(x, y)}
}

func NewKing(x, y int) *Pawn {
	return &Pawn{NewChessman(x, y)}
}

func NewQueen(x, y int) *Pawn {
	return &Pawn{NewChessman(x, y)}
}

func (p *Pawn) UpdateX(x int) {
	p.x = x
}

func (p *Pawn) UpdateY(y int) {
	p.y = y
}

func (p *Pawn) GetValidXMoves() []int {
	return []int{1, 2, 3, 4}
}

func (p *Pawn) GetValidYMoves() []int {
	return []int{1, 2, 3, 4}
}
