package models

import (
	"encoding/csv"
	"os"
)

var StoreData = map[string]string{}

func LoadStoreData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] { // Skip header
		StoreData[record[0]] = record[1]
	}
	return nil
}
