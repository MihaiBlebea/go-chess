package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func StringToPositions(input string) (Position, Position, error) {

	defaultErr := errors.New("Position transformation went wrong for " + input)

	// Remove the new line character
	input = strings.TrimSuffix(input, "\n")

	// Validate input
	regexTempl := regexp.MustCompile(`(?m)[A-Ha-h]-[1-8] [A-Ha-h]-[1-8]`)
	inputRegexResult := regexTempl.FindAllString(input, -1)

	if len(inputRegexResult) < 1 {
		return Position{}, Position{}, defaultErr
	}

	inputs := strings.Split(input, " ")

	var positions []Position

	for _, input := range inputs {
		// Split the coordonates into x and y
		coordinates := strings.Split(input, "-")

		if len(coordinates) != 2 {
			return Position{}, Position{}, defaultErr
		}

		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			return Position{}, Position{}, defaultErr
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
			return Position{}, Position{}, defaultErr
		}

		positions = append(positions, Position{x, y-1})
	}

	return positions[0], positions[1], nil
}
