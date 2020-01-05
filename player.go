package main

import (
	"fmt"

	"github.com/google/uuid"
)

type PlayerColor int 

const (
	white PlayerColor = iota
	black
)

type Player struct {
	id string
	color PlayerColor
	chessmen []IChessman
}

func NewPlayer(color PlayerColor) *Player {
	uuid := uuid.New()
	id := uuid.ID()

	player := &Player{
		id: fmt.Sprint(id),
		color: color,
	}

	// Generate pawns
	for _, x := range []int{1,2,3,4,5,6,7,8} {
		if player.color == white {
			player.AddChessman(NewPawn(x, 2))
		} else {
			player.AddChessman(NewPawn(x, 7))
		}
	}

	// Generate second row chessmen
	if player.color == white {
		player.AddChessman(NewRook(1, 1))
		player.AddChessman(NewRook(8, 1))

		player.AddChessman(NewKnight(2, 1))
		player.AddChessman(NewKnight(7, 1))

		player.AddChessman(NewBishop(3, 1))
		player.AddChessman(NewBishop(6, 1))

		player.AddChessman(NewKing(4, 8))
		player.AddChessman(NewQueen(5, 8))
	} else {
		player.AddChessman(NewRook(1, 8))
		player.AddChessman(NewRook(8, 8))

		player.AddChessman(NewKnight(2, 8))
		player.AddChessman(NewKnight(7, 8))

		player.AddChessman(NewBishop(3, 8))
		player.AddChessman(NewBishop(6, 8))

		player.AddChessman(NewKing(4, 8))
		player.AddChessman(NewQueen(5, 8))
	}

	return player
}

func (p *Player) AddChessman(chessman IChessman) {
	p.chessmen = append(p.chessmen, chessman)
}

func (p *Player) RemoveChessman(chessman *Chessman) {
	res := []IChessman{}
	for _, cm := range p.chessmen {
		if cm.equal(chessman) {
			res = append(res, cm)
		}
	}

	p.chessmen = res
}