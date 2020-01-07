package main

import "fmt"

// main function
func main() {
	rook := NewRook(black, 1, 3)
	fmt.Println(rook.GetValidMoves())
	// fmt.Println(rook.GetValidTakes())
}
