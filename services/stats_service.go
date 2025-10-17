package services

import (
	"nav/models"
	"time"
)

// Stats 统计信息结构体
type Stats struct {
	TotalRecords   int
	TodayRecords   int
	WeekRecords    int
	MonthRecords   int
	AverageChars   int
	TotalChars     int
	LongestRecord  models.Record
	ShortestRecord models.Record
}

// CalculateStats 计算统计数据
func CalculateStats() Stats {
	stats := Stats{}
	records := GetRecords()

	// 基本统计
	stats.TotalRecords = len(records)

	// 今日统计
	today := time.Now().Format("2006-01-02")
	todayCount := 0
	for _, record := range records {
		if record.CreatedAt.Format("2006-01-02") == today {
			todayCount++
		}
	}
	stats.TodayRecords = todayCount

	// 本周统计
	weekCount := 0
	year, week := time.Now().ISOWeek()
	for _, record := range records {
		recordYear, recordWeek := record.CreatedAt.ISOWeek()
		if recordYear == year && recordWeek == week {
			weekCount++
		}
	}
	stats.WeekRecords = weekCount

	// 月度统计
	month := time.Now().Format("2006-01")
	monthCount := 0
	for _, record := range records {
		if record.CreatedAt.Format("2006-01") == month {
			monthCount++
		}
	}
	stats.MonthRecords = monthCount

	// 内容长度统计
	totalChars := 0
	if len(records) > 0 {
		stats.LongestRecord = records[0]
		stats.ShortestRecord = records[0]

		for _, record := range records {
			contentLength := len(record.Content)
			totalChars += contentLength

			if contentLength > len(stats.LongestRecord.Content) {
				stats.LongestRecord = record
			}
			if contentLength < len(stats.ShortestRecord.Content) && contentLength > 0 {
				stats.ShortestRecord = record
			}
		}
	}

	stats.TotalChars = totalChars
	if len(records) > 0 {
		stats.AverageChars = totalChars / len(records)
	}

	return stats
}
