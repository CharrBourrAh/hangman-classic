package game

import (
	"fmt"
	"io/ioutil"
	"log"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func Init() {
	var data HangManData
	//data.ToFind = RandomWord("hangman-classic/data/word.txt")
	//data.Word = len(data.ToFind)
	data.Attempts = 10
}

func ReadFile(nameFile string) []byte {
	fileContent, err := ioutil.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string
	return fileContent
}

func ShowHangman(hangman []rune, attempts int) [7][9]string {
	var tab [7][9]string
	//var tabReturn [10]string
	for i := 0; i < 7; i++ {
		for j := 0; j < 9; j++ {
			if rune(hangman[i*j]) != 10 {
				tab[i][j] = string(hangman[i*j])
			}
		}
	}
	fmt.Println(tab)
	return tab
	//sfpijfohizrzhfuofheu
}

func RandomWord(wordFile []rune) [][]string {
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
