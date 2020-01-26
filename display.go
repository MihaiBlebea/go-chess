package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"regexp"
)

func render(board *Board) string {
	var result string

	result += "    | A | B | C | D | E | F | G | H |    " + "\n"
	result += "-----------------------------------------\n"

	for row := 8; row > 0; row-- {
	
		line := fmt.Sprintf("| %d |", row)
		for col := 0; col < 8; col++ {
			// Logic to display the chessmen
			spot := board.GetSpot(col, row - 1)
			chessman := spot.GetChessman()

			var chessmanSign string
			if chessman != nil {
				name := reflect.TypeOf(chessman).Name()
				chessmanSign = GetEmoticon(name, chessman.IsWhite())
			} else {
				chessmanSign = " "
			}

			line += " " + chessmanSign + " |"
		}
		line += fmt.Sprintf(" %d |", row)

		result += line + "\n"
		result += "-----------------------------------------\n"
	}
	result += "    | A | B | C | D | E | F | G | H |    " + "\n"
	
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

func GetEmoticon(chessmanName string, isWhite bool) string {
	if isWhite == true {
		switch chessmanName {
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

	switch chessmanName {
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
	regexTempl := regexp.MustCompile(`(?m)[A-H]-[1-8]`)
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
		log.Panic(err)
	}
	var x int

	switch coordinates[0] {
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

	fmt.Println(input, x, y - 1)
	return x, y - 1, nil
}
