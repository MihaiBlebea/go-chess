package main

import "testing"


func TestPawnMovementCorrect( t *testing.T) {

	board := Board{}
	board.Reset()

	// Move white Pawn 2 spaces in front
	res := board.Move(Position{0, 1}, Position{0, 3})
	if res != true {
		t.Errorf("White Pawn should be able to move 2 spaces in front if it starts in it's own half of the board")
	}

	// Move white Pawn one space in front beyond the middle of the board
	res = board.Move(Position{0, 3}, Position{0, 4})
	if res != true {
		t.Errorf("White Pawn should be able to move 1 space in front beyong the middle of the board")
	}

	// White Pawn captures a Black Pawn
	board.Move(Position{7, 1}, Position{7, 3})
	board.Move(Position{6, 6}, Position{6, 4})
	res = board.Move(Position{7, 3}, Position{6, 4})
	if res != true {
		t.Errorf("White Pawn should be able to capture an black pawn if it's situated one space across from it")
	}

	// Move black Pawn 2 spaces in front
	res = board.Move(Position{1, 6}, Position{1, 4})
	if res != true {
		t.Errorf("Black Pawn should be able to move 2 spaces in front if it starts in it's own half of the board")
	}

	// Move black Pawn one space in front beyond the middle of the board
	res = board.Move(Position{1, 4}, Position{1, 3})
	if res != true {
		t.Errorf("Black Pawn should be able to move 1 space in front beyong the middle of the board")
	}
}

func TestPawnMovementWrong( t *testing.T) {

	board := Board{}
	board.Reset()

	// Move white Pawn 1 space in front and 2 spaces to the right
	res := board.Move(Position{0, 1}, Position{2, 2})
	if res != false {
		t.Errorf("White Pawn should not move to the side")
	}

	// Move white Pawn 5 spaces in front
	res = board.Move(Position{0, 1}, Position{0, 6})
	if res != false {
		t.Errorf("White Pawn should not be able to move 5 spaces in front at once")
	}

	// Move white Pawn backwords
	board.Move(Position{0, 1}, Position{0, 3})
	res = board.Move(Position{0, 3}, Position{0, 1})
	if res != false {
		t.Errorf("White Pawn should not be moving backwards")
	}


	// Move black Pawn 1 space in front and 2 spaces to the right
	res = board.Move(Position{1, 6}, Position{2, 5})
	if res != false {
		t.Errorf("Black Pawn should not move to the side")
	}

	// Move black Pawn 5 spaces in front
	res = board.Move(Position{1, 6}, Position{1, 2})
	if res != false {
		t.Errorf("Black Pawn should not be able to move 5 spaces in front at once")
	}

	// Move black Pawn backwords
	board.Move(Position{1, 6}, Position{1, 4})
	res = board.Move(Position{1, 4}, Position{1, 6})
	if res != false {
		t.Errorf("Black Pawn should not be moving backwards")
	}
}