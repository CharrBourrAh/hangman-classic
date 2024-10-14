package save

import (
	"encoding/json"
	"hangman-classic/pkg/structs"
	"log"
	"os"
)

func StopAndSaveGame(data *structs.HangManData) {
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

func StartWithFlag(Start string) structs.HangManData {
	var data structs.HangManData
	JsonData, _ := os.ReadFile(Start)
	err := json.Unmarshal(JsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
