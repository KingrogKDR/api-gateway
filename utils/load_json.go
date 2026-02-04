package utils

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJson[T any](path string) ([]T, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening json files")
		return nil, err
	}
	defer file.Close()

	var data []T
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	return data, err
}
