package main

import (
	"hangman-classic/internal/game"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "data/words.txt" || os.Args[1] == "data/words2.txt" || os.Args[1] == "data/words3.txt" {
			game.Init(os.Args[1])
		}
	} else {
		game.Menu()
	}
}
