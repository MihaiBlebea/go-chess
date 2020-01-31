package main

func main() {

	white := HumanPlayer{"Mihai", CommandLine{}}
	// black := HumanPlayer{"Cristina", CommandLine{}}
	black := SimulatedPlayer{[]string{"b-7 b-5", "e-7 e-5"}}

	game := Create(white, black)
	game.Play()
}
