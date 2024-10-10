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
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func ReadFile(nameFile string) []rune {
	fileContent, err := ioutil.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string
	return []rune(string(fileContent))
}

func TransfomTab(wordFile []rune) [][]string {
	// a integrer dans ReadFile
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
	for i := 0; i < len(list[randomWordPos]); i++ {
		data.Word += "_"
		data.ToFind += list[randomWordPos][i]
	}
}

func ShowHangman(hangman [][]string, attempts int) {
	for i := 0 + attempts*7 + 1; i < 7+attempts*7; i++ {
		if len(hangman[i]) == 0 {
			for iBis := 0; iBis < 9; iBis++ {
				fmt.Println("          ")
			}
			i++
		}
		for j := 0; j < len(hangman[i]); j++ {
			if hangman[i][j] != "\n" {
				fmt.Print(hangman[i][j])
			}
		}
		fmt.Print("\n")
	}

}

func Init() {
	var data HangManData
	data.Attempts = 10
	RandomWord(TransfomTab(ReadFile("data/words.txt")), &data)
	Game(&data)
}

func Game(data *HangManData) {
	copyWord := strings.Split(data.Word, "")
	for data.Word != data.ToFind && data.Attempts > 0 {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		//fmt.Println(ShowHangman(data.))
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
		fmt.Println("You've won horray :D\nYou've needed to find " + data.ToFind)
	}
}
