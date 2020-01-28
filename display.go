package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func render(board *Board) string {
	var result string

	// Black prison
	result += "Prison: "
	for _, capturedChessman := range board.whitePrison.GetCaptured() {
		chessmanSign := TransformChessmanToSign(capturedChessman)
		result += chessmanSign + " | "
	}
	result += "\n"

	result += "    | A | B | C | D | E | F | G | H |    " + "\n"
	result += "-----------------------------------------\n"

	for row := 8; row > 0; row-- {

		line := fmt.Sprintf("| %d |", row)
		for col := 0; col < 8; col++ {
			// Logic to display the chessmen
			spot := board.GetSpot(col, row-1)
			chessman := spot.GetChessman()

			chessmanSign := TransformChessmanToSign(chessman)

			line += " " + chessmanSign + " |"
		}
		line += fmt.Sprintf(" %d |", row)

		result += line + "\n"
		result += "-----------------------------------------\n"
	}
	result += "    | A | B | C | D | E | F | G | H |    " + "\n"

	// White prison
	result += "\n"
	result += "Prison: "
	for _, capturedChessman := range board.blackPrison.GetCaptured() {
		chessmanSign := TransformChessmanToSign(capturedChessman)
		result += chessmanSign + " | "
	}
	result += "\n"

	return result

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
}

func TransformChessmanToSign(chessman Chessman) string {

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

func TransformPosition(input string) (int, int, error) {

	defaultErr := errors.New("Position transformation went wrong for " + input)

	// Remove the new line character
	input = strings.TrimSuffix(input, "\n")

	// Validate input
	regexTempl := regexp.MustCompile(`(?m)[A-Ha-h]-[1-8]`)
	inputRegexResult := regexTempl.FindAllString(input, -1)

	if len(inputRegexResult) < 1 {
		return 0, 0, defaultErr
	}

	// Split the coordonates into x and y
	coordinates := strings.Split(input, "-")

	if len(coordinates) != 2 {
		return 0, 0, defaultErr
	}

	y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		return 0, 0, defaultErr
	}
	var x int

	switch strings.ToUpper(coordinates[0]) {
	case "A":
		x = 0
	case "B":
		x = 1
	case "C":
		x = 2
	case "D":
		x = 3
	case "E":
		x = 4
	case "F":
		x = 5
	case "G":
		x = 6
	case "H":
		x = 7
	default:
		return 0, 0, defaultErr
	}

	return x, y - 1, nil
}
