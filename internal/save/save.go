package save

import (
	"encoding/json"
	"log"
	"os"
)

type HangManData struct {
	ToFind   string
	Word     string
	Attempts int
}

func StopAndSaveGame(data *HangManData) {
	data.ToFind = "tester"
	data.Word = "t_st__"
	data.Attempts = 6
	mapReturn := map[string]interface{}{
		"ToFind":   data.ToFind,
		"Word":     data.Word,
		"Attempts": data.Attempts,
	}
	DataGameJson, _ := json.MarshalIndent(mapReturn, "", " ")
	err := os.WriteFile("save.json", DataGameJson, 0777)
	if err != nil {
		log.Fatal(err)
	}

}

func StartWithFlag(Start string) Hangman {
	var data Hangman
	JsonData, _ := os.ReadFile(Start)
	err := json.Unmarshal(JsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
