package main

import (
	"hangman-classic/internal/game"
	"hangman-classic/internal/save"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "words.txt" || os.Args[1] == "words2.txt" || os.Args[1] == "words3.txt" {
			game.ClearCMD()
			game.Init("data/" + os.Args[1])
		} else if os.Args[1] == "--startWith" {
			game.ClearCMD()
			game.Resume(save.StartWithFlag(os.Args[2]))
		}
	} else {
		game.Menu()
	}
}
