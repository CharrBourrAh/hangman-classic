package game

import (
	"fmt"
	"hangman-classic/internal/input"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"slices"
	"strings"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [][]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

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

func RandomWord(list [][]string, data *HangManData) {
	randomWordPos := rand.Intn(len(list))
	for i := 0; i < len(list[randomWordPos])-1; i++ {
		data.Word += "_"
		data.ToFind += list[randomWordPos][i]
	}
}

func ShowHangman(hangman [][]string, attempts int) {
	var hangmanReturn [][]string
	lenHangman := 7
	for i := 0; i < lenHangman; i++ {
		if len(hangman[i]) != 0 {
			hangmanReturn = append(hangmanReturn, hangman[i*attempts])
			for j := 0; j < 9; j++ {
				hangmanReturn[i] = append(hangmanReturn[i], hangman[i*attempts][j])
			}
		} else {
			i++
			lenHangman++
		}
	}
	for i2 := 0; i2 < len(hangmanReturn); i2++ {
		if len(hangmanReturn[i2]) != 0 {
			for j2 := 0; j2 < len(hangmanReturn[i2]); j2++ {
				fmt.Print(hangmanReturn[i2][j2])
			}
		}
		fmt.Print("\n")
	}
}

func Init(WordFile string) {
	if WordFile == "" {
		WordFile = "data/words.txt"
	}
	var data HangManData
	data.Attempts = 10
	RandomWord(ReadFile(WordFile), &data)
	data.HangmanPositions = ReadFile("data/hangman.txt")
	Game(&data)
}

func Game(data *HangManData) {
	copyWord := strings.Split(data.Word, "")
	var usedLetters []string
	for data.Word != data.ToFind && data.Attempts > 0 {
		//ShowHangman(data.HangmanPositions, data.Attempts)
		fmt.Println(data.Word)
		fmt.Println(data.ToFind)
		userInput := strings.ToLower(input.Input())
		for i := 0; i < len(data.ToFind); i++ {
			if len(userInput) == 2 {
				//menu.Menu()
				if userInput == "/r" {
					Init("")
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
