package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func render(board *Board) string {
	
	var result string
	for row := 0; row < 10; row++ {
		if row == 0 || row == 9 {
			result += "    | A | B | C | D | E | F | G | H |    " + "\n"
		} else {

			line := fmt.Sprintf("| %d |", 9-row)
			for col := 0; col < 8; col++ {
				// Logic to display the chessmen
				spot := board.GetSpot(col, row-1)
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
			line += fmt.Sprintf(" %d |", 9-row)

			result += line + "\n"
		}
		if row != 9 {
			result += "-----------------------------------------\n"
		}
	}

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

func TransformPosition(code string) (int, int, error) {
	// Remove the new line character
	code = strings.TrimSuffix(code, "\n")
	
	// Split the coordonates into x and y
	coordinates := strings.Split(code, "-")

	if len(coordinates) != 2 {
		return 0, 0, errors.New("Position transformation went wrong for " + code)
	}

	Y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		log.Panic(err)
	}
	switch coordinates[0] {
	case "A":
		return 0, Y - 1, nil
	case "B":
		return 0, Y - 1, nil
	case "C":
		return 0, Y - 1, nil
	case "D":
		return 0, Y - 1, nil
	case "E":
		return 0, Y - 1, nil
	case "F":
		return 0, Y - 1, nil
	case "G":
		return 0, Y - 1, nil
	case "H":
		return 0, Y - 1, nil
	default:
		return 0, 0, nil
	}
}
