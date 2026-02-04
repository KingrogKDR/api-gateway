package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/KingrogKDR/api-gateway/models"
)

// loaded sequentially for now, can be loaded parallely using go routines
func LoadJson[T1 models.User, T2 models.Order](path1 string, path2 string) ([]T1, []T2, error) {
	userFile, err := os.Open(path1)
	if err != nil {
		log.Fatal("Error opening user file:", err)
		return nil, nil, err
	}
	defer userFile.Close()

	orderFile, err := os.Open(path1)
	if err != nil {
		log.Fatal("Error opening order file:", err)
		return nil, nil, err
	}
	defer orderFile.Close()

	var userData []T1
	decoder := json.NewDecoder(userFile)
	err = decoder.Decode(&userData)

	var orderData []T2
	decoder = json.NewDecoder(orderFile)
	err = decoder.Decode(&orderData)

	return userData, orderData, err
}
