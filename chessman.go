package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Chessman
type Chessman struct {
	id    string
	color Color
	Position
}

type IChessman interface {
	GetValidMoves() []Position
	GetId() string
	GetX() int
	GetY() int
	equal(chessman *Chessman) bool
}

func NewChessman(color Color, x, y int) *Chessman {
	uuid := uuid.New()
	id := uuid.ID()

	return &Chessman{
		id:       fmt.Sprint(id),
		color:    color,
		Position: Position{x, y},
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

