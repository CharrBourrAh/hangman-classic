package main

import (
	"hangman-classic/internal/game"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		game.Init(os.Args[1])
	} else {
		game.Menu()
	}
}
