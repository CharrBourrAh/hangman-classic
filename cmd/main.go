package main

import (
	"fmt"
	"hangman-classic/internal/game"
)

func main() {

	game.ShowHangman([]rune("         \n         \n         \n         \n         \n         \n=========\n\n         \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n\n  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n\n  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n\n  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n\n  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n\n  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n\n  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n\n  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n\n  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n\n"), 1)
	fmt.Println(game.ReadFile("data/words.txt"))
	test := []rune{98, 97, 110, 99, 10, 98, 117, 114, 101, 97, 117, 10, 99, 97, 98, 105, 110, 101, 116, 10, 99, 97, 114, 114, 101, 97, 117, 10}
	fmt.Print(game.RandomWord(test))
}
