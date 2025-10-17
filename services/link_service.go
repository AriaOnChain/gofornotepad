package services

import (
	"encoding/json"
	"log"
	"nav/models"
	"os"
	"strings"
	"time"
)

var linkRecords []models.LinkRecord
var linkDataFile = "data/links.json"

// LoadLinkRecords 加载链接记录
func LoadLinkRecords() {
	data, err := os.ReadFile(linkDataFile)
	if err != nil {
		linkRecords = []models.LinkRecord{}
		return
	}
	json.Unmarshal(data, &linkRecords)
}

// SaveLinkRecords 保存链接记录
func SaveLinkRecords() {
	data, err := json.MarshalIndent(linkRecords, "", "  ")
	if err != nil {
		log.Printf("保存链接数据失败: %v", err)
		return
	}
	os.WriteFile(linkDataFile, data, 0644)
}

// GetLinkRecords 获取所有链接记录
func GetLinkRecords() []models.LinkRecord {
	return linkRecords
}

// AddLinkRecord 添加新链接记录
func AddLinkRecord(title, content, link string) {
	// 验证链接格式
	if link != "" && !isValidURL(link) {
		if !strings.Contains(link, "://") {
			link = "https://" + link
		}
	}

	newRecord := models.LinkRecord{
		ID:        getNextLinkID(),
		Title:     title,
		Content:   content,
		Link:      link,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	linkRecords = append(linkRecords, newRecord)
	SaveLinkRecords()
}

// DeleteLinkRecord 删除链接记录
func DeleteLinkRecord(id int) bool {
	for i, record := range linkRecords {
		if record.ID == id {
			linkRecords = append(linkRecords[:i], linkRecords[i+1:]...)
			SaveLinkRecords()
			return true
		}
	}
	return false
}

// UpdateLinkRecord 更新链接记录
func UpdateLinkRecord(id int, title, content, link string) bool {
	// 验证链接格式
	if link != "" && !isValidURL(link) {
		if !strings.Contains(link, "://") {
			link = "https://" + link
		}
	}

	for i := range linkRecords {
		if linkRecords[i].ID == id {
			linkRecords[i].Title = title
			linkRecords[i].Content = content
			linkRecords[i].Link = link
			linkRecords[i].UpdatedAt = time.Now()
			SaveLinkRecords()
			return true
		}
	}
	return false
}

// SearchLinkRecords 搜索链接记录
func SearchLinkRecords(query string) []models.LinkRecord {
	var results []models.LinkRecord
	for _, record := range linkRecords {
		if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(record.Content), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(record.Link), strings.ToLower(query)) {
			results = append(results, record)
		}
	}
	return results
}

func getNextLinkID() int {
	if len(linkRecords) == 0 {
		return 1
	}
	maxID := 0
	for _, record := range linkRecords {
		if record.ID > maxID {
			maxID = record.ID
		}
	}
	return maxID + 1
}

func isValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
