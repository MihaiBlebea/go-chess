package main

import (
	"fmt"
	"reflect"
)

type Display interface {
	Render(board *Board) string
}

type CommandLine struct {
}

// 	return `
// 	    | a | b | c | d | e | f | g | h |
// 	-----------------------------------------
// 	| 8 |   |   |   |   |   |   |   |   | 8 |
// 	-----------------------------------------
// 	| 7 |   |   |   |   |   |   |   |   | 7 |
// 	-----------------------------------------
// 	| 6 |   |   |   |   |   |   |   |   | 6 |
// 	-----------------------------------------
// 	| 5 |   |   |   |   |   |   |   |   | 5 |
// 	-----------------------------------------
// 	| 4 |   |   |   |   |   |   |   |   | 4 |
// 	-----------------------------------------
// 	| 3 |   |   |   |   |   |   |   |   | 3 |
// 	-----------------------------------------
// 	| 2 |   |   |   |   |   |   |   |   | 2 |
// 	-----------------------------------------
// 	| 1 |   |   |   |   |   |   |   |   | 1 |
// 	-----------------------------------------
// 	    | a | b | c | d | e | f | g | h |
// 	`

func (cl CommandLine) Render(board *Board) string {
	var result string

	result += "    | A | B | C | D | E | F | G | H |    " + "\n"
	result += "-----------------------------------------\n"

	for row := 8; row > 0; row-- {

		line := fmt.Sprintf("| %d |", row)
		for col := 0; col < 8; col++ {
			// Logic to display the chessmen
			spot := board.GetSpot(col, row-1)
			chessman := spot.GetChessman()

			chessmanSign := cl.transformChessmanToSign(chessman)

			line += " " + chessmanSign + " |"
		}
		line += fmt.Sprintf(" %d |", row)

		result += line + "\n"
		result += "-----------------------------------------\n"
	}
	result += "    | A | B | C | D | E | F | G | H |    " + "\n"

	// White prison
	result += "\n"
	result += "Captured: "
	for _, capturedChessman := range board.captured {
		chessmanSign := cl.transformChessmanToSign(capturedChessman)
		result += chessmanSign + " | "
	}
	result += "\n"

	return result
}

func (cl CommandLine) transformChessmanToSign(chessman Chessman) string {

	if chessman == nil {
		return " "
	}

	name := reflect.TypeOf(chessman).Name()

	if chessman.IsWhite() == true {
		switch name {
		case "Pawn":
			return "♙"
		case "Rook":
			return "♖"
		case "Knight":
			return "♘"
		case "Bishop":
			return "♗"
		case "King":
			return "♔"
		case "Queen":
			return "♕"
		default:
			return " "
		}
	}

	switch name {
	case "Pawn":
		return "♟"
	case "Rook":
		return "♜"
	case "Knight":
		return "♞"
	case "Bishop":
		return "♝"
	case "King":
		return "♚"
	case "Queen":
		return "♛"
	default:
		return " "
	}
}
