package main

import "fmt"

// main function
func main() {
	king := NewKing(black, 4, 2)
	fmt.Println(king.GetValidMoves())
	// fmt.Println(rook.GetValidTakes())
}
