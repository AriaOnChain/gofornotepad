// 工具函数
package utils

import (
	"encoding/json"
	"nav/models"
	"os"
)

func LoadRecords(dataFile string) ([]models.Record, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return []models.Record{}, nil
	}

	var records []models.Record
	json.Unmarshal(data, &records)
	return records, nil
}

func SaveRecords(dataFile string, records []models.Record) error {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func GetNextID(records []models.Record) int {
	if len(records) == 0 {
		return 1
	}
	maxID := 0
	for _, record := range records {
		if record.ID > maxID {
			maxID = record.ID
		}
	}
	return maxID + 1
}
