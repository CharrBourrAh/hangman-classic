package main

import (
	"hangman-classic/internal/game"
	"hangman-classic/internal/save"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "words.txt" || os.Args[1] == "words2.txt" || os.Args[1] == "words3.txt" {
			// If a words file is specified, the game will launch by using a random word for the given file
			game.ClearCMD()
			game.Init("data/" + os.Args[1])
		} else if os.Args[1] == "--startWith" {
			// If this argument is given, the game will resume the game thanks to the information on the save file given after
			game.ClearCMD()
			game.Resume(save.StartWithFlag("save/" + os.Args[2]))
		}
	} else {
		// If no argument has been given, the game will boot the main menu
		game.ClearCMD()
		game.Menu()
	}
}
