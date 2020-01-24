package main

import (
	"fmt"
	"go-chess/chessmen"
	"log"
)


type Color interface {
	
}

type Position interface {
}


// main function
func main() {
	queen, err := chessmen.New("queen", "white", 1, 1)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(queen.GetValidMoves())
	// fmt.Println(rook.GetValidTakes())
}
