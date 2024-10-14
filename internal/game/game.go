package game

import (
	"fmt"
	"hangman-classic/internal/input"
	"hangman-classic/pkg/structs"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func ReadFile(nameFile string) [][]string {
	content, err := os.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string
	wordFile := []rune(string(content))
	var wordTab [][]string
	counter := 0
	wordTab = append(wordTab, []string{})
	for i := 0; i < len(wordFile); i++ {
		if wordFile[i] != 10 {
			wordTab[counter] = append(wordTab[counter], string(wordFile[i]))
		} else {
			counter++
			wordTab = append(wordTab, []string{})
		}
	}
	return wordTab
}

func RandomWord(list [][]string, data *structs.HangManData) {
	randomWordPos := rand.Intn(len(list))
	for i := 0; i < len(list[randomWordPos]); i++ {
		data.Word += "_"
		data.ToFind += list[randomWordPos][i]
	}
}

func ShowHangman(hangman [][]string, attempts int) {
	for i := 8 * attempts; i < 8+8*attempts-1; i++ {
		for j := 0; j < len(hangman[i]); j++ {
			fmt.Print(hangman[i][j])
		}
		fmt.Print("\n")
	}
}

func Init(WordFile string) {
	if WordFile == "" {
		WordFile = "data/words.txt"
	}
	var data structs.HangManData
	data.Attempts = 10
	RandomWord(ReadFile(WordFile), &data)
	data.HangmanPositions = ReadFile("data/hangman.txt")
	Game(&data)
}

func Game(data *structs.HangManData) {
	copyWord := strings.Split(data.Word, "")
	var usedLetters []string
	for data.Word != data.ToFind && data.Attempts > 0 {
		ShowHangman(data.HangmanPositions, 10-data.Attempts)
		fmt.Println(data.Word)
		fmt.Println(data.ToFind)
		//fmt.Println(data.HangmanPositions)
		userInput := strings.ToLower(input.Input())
		for i := 0; i < len(data.ToFind); i++ {
			if len(userInput) == 2 {
				//menu.Menu()
				if userInput == "/r" {
					Init("")
				} else if userInput == "/m" {
					Menu()
				}
			} else if len(userInput) == 1 {
				for j := 0; j < len(data.Word); j++ {
					if userInput == string(data.ToFind[j]) {
						copyWord[j] = userInput
					}
				}
			} else if len(userInput) > 2 {
				if userInput == data.ToFind {
					data.Word = data.ToFind
					break
				} else {
					data.Attempts -= 1
					break
				}
			}
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("This type of terminal is not supported by this game. Please use Windows' newer or classic Terminal app")
				return
			}
		}
		if strings.Join(copyWord, "") == data.Word && slices.Contains(usedLetters, userInput) == false {
			data.Attempts -= 1
			fmt.Println("Not present in the word,", data.Attempts, "attempts remaining")
		} else if slices.Contains(usedLetters, userInput) == true {
			fmt.Println("You've already used this letter before")
		} else {
			data.Word = strings.Join(copyWord, "")
		}
		usedLetters = append(usedLetters, userInput)
	}
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("This type of terminal is not supported by this game. Please use Windows' newer or classic Terminal app")
		return
	}
	if data.Attempts == 0 {
		fmt.Println(data.Word + "\nYou loose :( \nYou've needed to find " + data.ToFind)
	} else {
		fmt.Println("You've won horray :D\nYou've successfully found " + data.ToFind)
	}
}

func Menu() {
	clearCMD()
	mainMenuAscii := "  __  __       _                                    \n |  \\/  |     (_)                                   \n | \\  / | __ _ _ _ __    _ __ ___   ___ _ __  _   _ \n | |\\/| |/ _` | | '_ \\  | '_ ` _ \\ / _ \\ '_ \\| | | |\n | |  | | (_| | | | | | | | | | | |  __/ | | | |_| |\n |_|  |_|\\__,_|_|_| |_| |_| |_| |_|\\___|_| |_|\\__,_|\n                                                    \n                                                    "
	fmt.Print(mainMenuAscii)
	fmt.Println("\n" + "\033[32m" + "s" + "\033[0m" + " : launch a game")
	fmt.Println("\x1b[33m" + "o" + "\033[0m" + " : opens the game's settings")
	fmt.Println("\033[31m" + "q / e" + "\033[0m" + " : exit the game")
	choice := input.Input()
	if choice == "s" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("This type of terminal is not supported by this game. Please use Windows' newer or classic Terminal app")
			return
		}
		Init("") // launches the game
	}
	if choice == "o" {
		clearCMD()
		fmt.Println("\n   _____      _   _   _                 \n  / ____|    | | | | (_)                \n | (___   ___| |_| |_ _ _ __   __ _ ___ \n  \\___ \\ / _ \\ __| __| | '_ \\ / _` / __|\n  ____) |  __/ |_| |_| | | | | (_| \\__ \\\n |_____/ \\___|\\__|\\__|_|_| |_|\\__, |___/\n                               __/ |    \n                              |___/     \n")
		fmt.Println("w : Choose which word file you want to use to play")
		choice = input.Input()
	}
	if choice == "q" || choice == "e" {
		clearCMD()
		os.Exit(3) // Exit the program
	}
}

func clearCMD() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("This type of terminal is not supported by this game. Please use Windows' newer or classic Terminal app")
		return
	}
}
