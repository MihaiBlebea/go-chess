package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Player interface {
	Move() (Position, Position)
	Render(board *Board)
	Alert(message string)
}

type HumanPlayer struct {
	name    string
	display Display
}

func (p HumanPlayer) Move() (Position, Position) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Make a move: ")
	input, _ := reader.ReadString('\n')

	// Validate the player input
	start, end , err := StringToPositions(input)

	if err != nil {
		log.Panic(err)
	}
	return start, end

}

func (p HumanPlayer) Render(board *Board) {
	fmt.Println(p.display.Render(board))
}

func (p HumanPlayer) Alert(message string) {
	fmt.Println(message)
}


type SimulatedPlayer struct {
	moves []string
}

func (p SimulatedPlayer) Move() (Position, Position) {
	defer p.removeFirstCommand()

	start, end, err := StringToPositions(p.moves[0])
	if err != nil {
		log.Panic(err)
	}

	return start, end
}

func (p SimulatedPlayer) Render(board *Board) {
	//
}

func (p SimulatedPlayer) Alert(message string) {
	fmt.Println(message)
}

func (p *SimulatedPlayer) removeFirstCommand() {
	copy(p.moves[0:], p.moves[1:])
	p.moves[len(p.moves)-1] = ""
	p.moves = p.moves[:len(p.moves)-1]
}