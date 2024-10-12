package game

import (
	"fmt"
	"hangman-classic/internal/input"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [][]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func ReadFile(nameFile string) [][]string {
	content, err := ioutil.ReadFile(nameFile)
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
	hangmanReturn := hangman[attempts*7 : (attempts+1)*7]
	for i := 0; i < len(hangmanReturn); i++ {
		for j := 0; j < len(hangmanReturn[i]); j++ {
			fmt.Print(hangmanReturn[i][j])
		}
		fmt.Print("\n")
	}
}
func Init() {
	var data HangManData
	data.Attempts = 10
	RandomWord(ReadFile("data/words.txt"), &data)
	data.HangmanPositions = ReadFile("data/hangman.txt")
	Game(&data)
}

func Game(data *HangManData) {
	copyWord := strings.Split(data.Word, "")
	for data.Word != data.ToFind && data.Attempts > 0 {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		ShowHangman(data.HangmanPositions, data.Attempts)
		fmt.Println(data.Word)
		userInput := input.Input()
		for i := 0; i < len(data.ToFind); i++ {
			if len(userInput) > 1 {
				//menu.Menu()
				if userInput == "/r" {
					Init()
				}
			} else {
				for j := 0; j < len(data.Word); j++ {
					if userInput == string(data.ToFind[j]) {
						copyWord[j] = userInput
					}
				}
			}
		}
		if strings.Join(copyWord, "") == data.Word {
			data.Attempts -= 1
		} else {
			data.Word = strings.Join(copyWord, "")
		}
		println(data.Attempts)
	}
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	if data.Attempts == 0 {
		fmt.Println(data.Word + "\nYou loose :( \nYou've needed to find " + data.ToFind)
	} else {
		fmt.Println("You've won horray :D\nYou've successfully found " + data.ToFind)
	}
}
