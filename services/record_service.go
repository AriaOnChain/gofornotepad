package services

import (
	"encoding/json"
	"log"
	"nav/models"
	"os"
	"strings"
	"time"
)

var records []models.Record
var dataFile = "data/records.json"

// LoadRecords 加载记录数据
func LoadRecords() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		records = []models.Record{}
		return
	}
	json.Unmarshal(data, &records)
}

// SaveRecords 保存记录数据
func SaveRecords() {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		log.Printf("保存数据失败: %v", err)
		return
	}
	os.WriteFile(dataFile, data, 0644)
}

// GetRecords 获取所有记录
func GetRecords() []models.Record {
	return records
}

// AddRecord 添加新记录
func AddRecord(title, content string) {
	newRecord := models.Record{
		ID:        getNextID(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	records = append(records, newRecord)
	SaveRecords()
}

// DeleteRecord 删除记录
func DeleteRecord(id int) bool {
	for i, record := range records {
		if record.ID == id {
			records = append(records[:i], records[i+1:]...)
			SaveRecords()
			return true
		}
	}
	return false
}

// UpdateRecord 更新记录
func UpdateRecord(id int, title, content string) bool {
	for i := range records {
		if records[i].ID == id {
			records[i].Title = title
			records[i].Content = content
			records[i].UpdatedAt = time.Now()
			SaveRecords()
			return true
		}
	}
	return false
}

// SearchRecords 搜索记录
func SearchRecords(query string) []models.Record {
	var results []models.Record
	for _, record := range records {
		if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(record.Content), strings.ToLower(query)) {
			results = append(results, record)
		}
	}
	return results
}

func getNextID() int {
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
