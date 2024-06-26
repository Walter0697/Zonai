package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func SaveJsonFile[T interface{}](data T, filename string) {
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadJsonFile[T interface{}](filename string) T {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var data T
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
