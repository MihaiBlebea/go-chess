package chessmen

import (
	"fmt"

	"github.com/google/uuid"
	"strings"
	"errors"
)

type Chessman interface {
	GetValidMoves() []Position
	GetValidTakes() []Position
	GetId() string
	GetX() int
	GetY() int
}

func New(name, color string, x, y int) (Chessman, error) {
	id := uuid.New().ID()

	name = strings.ToLower(name)
	position := Position{x, y}

	var col Color
	switch color {
	case "white":
		col = White
	case "black":
		col = Black
	default:
		return nil, errors.New("Color is undefined. Please choose between `white` or `black`") 
	}

	switch name {
	case "pawn":
		return Pawn{fmt.Sprint(id), position, col}, nil
	case "rook":
		return Rook{fmt.Sprint(id), position, col}, nil
	case "knight":
		return Knight{fmt.Sprint(id), position, col}, nil
	case "bishop":
		return Bishop{fmt.Sprint(id), position, col}, nil
	case "queen":
		return &Queen{fmt.Sprint(id), position, col}, nil
	case "king":
		return King{fmt.Sprint(id), position, col}, nil
	default:
		return nil, errors.New("Chessman name is invalid. Please supply a valid type") 
	}
}

