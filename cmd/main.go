package main

import (
	"hangman-classic/internal/save"
)

func main() {
	/*if len(os.Args) > 1 {
		game.Init(os.Args[1])
	} else {
		game.Init("")
	}*/
	save.StopAndSaveGame(save.Hangman{})
}
