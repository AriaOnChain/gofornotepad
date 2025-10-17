package handlers

import (
	"fmt"
	"html/template"
	"nav/models"
	"nav/services"
	"nav/templates"
	"net/http"
	"os"
	"strings"
	"time"
)

// InitServices 初始化服务
func InitServices() {
	services.LoadRecords()
	services.LoadLinkRecords()
}

// IndexHandler 首页处理器
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		createDefaultTemplate()
		tmpl, _ = template.ParseFiles("templates/index.html")
	}

	searchQuery := r.URL.Query().Get("q")
	var displayRecords []models.Record

	if searchQuery != "" {
		displayRecords = services.SearchRecords(searchQuery)
	} else {
		displayRecords = services.GetRecords()
	}

	data := map[string]interface{}{
		"Title":       "demo",
		"Title1":      "待实现的零碎想法~",
		"Records":     displayRecords,
		"TotalCount":  len(services.GetRecords()),
		"SearchQuery": searchQuery,
		"Success":     r.URL.Query().Get("success"),
		"Error":       r.URL.Query().Get("error"),
	}

	tmpl.Execute(w, data)
}

// LinksHandler 链接页面处理器
func LinksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/links.html")
	if err != nil {
		createLinksTemplate()
		tmpl, _ = template.ParseFiles("templates/links.html")
	}

	searchQuery := r.URL.Query().Get("q")
	var displayRecords []models.LinkRecord

	if searchQuery != "" {
		displayRecords = services.SearchLinkRecords(searchQuery)
	} else {
		displayRecords = services.GetLinkRecords()
	}

	data := map[string]interface{}{
		"Title":       "链接收藏",
		"Title1":      "🔗 我的链接库",
		"Records":     displayRecords,
		"TotalCount":  len(services.GetLinkRecords()),
		"SearchQuery": searchQuery,
		"Success":     r.URL.Query().Get("success"),
		"Error":       r.URL.Query().Get("error"),
	}

	tmpl.Execute(w, data)
}

// AboutHandler 关于页面处理器
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		createAboutTemplate()
		tmpl, _ = template.ParseFiles("templates/about.html")
	}

	data := map[string]interface{}{
		"Title":        "关于我们",
		"AppName":      "零碎想法记录本",
		"Version":      "1.1.0",
		"TotalRecords": len(services.GetRecords()),
	}

	tmpl.Execute(w, data)
}

// StatsHandler 统计页面处理器
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	stats := services.CalculateStats()

	tmpl, err := template.ParseFiles("templates/stats.html")
	if err != nil {
		createStatsTemplate()
		tmpl, _ = template.ParseFiles("templates/stats.html")
	}

	data := map[string]interface{}{
		"Title":   "数据统计",
		"Stats":   stats,
		"Records": services.GetRecords(),
	}

	tmpl.Execute(w, data)
}

// DisplayPendingTasks 显示待办任务
func DisplayPendingTasks() {
	records := services.GetRecords()
	if len(records) == 0 {
		fmt.Println("\n列表暂无记录")
		fmt.Println("提示: 在浏览器中添加你的第一条记录吧！")
		return
	}

	fmt.Printf("\n 当前列表 (共 %d 项):\n", len(records))
	fmt.Println(strings.Repeat("=", 50))

	for i, record := range records {
		timeStr := record.CreatedAt.Format("01-02 15:04")
		contentPreview := ""
		if record.Content != "" {
			if len(record.Content) > 100 {
				contentPreview = " - " + record.Content[:100] + "..."
			} else {
				contentPreview = " - " + record.Content
			}
		}

		fmt.Printf("%d. %s%s\n", i+1, record.Title, contentPreview)
		fmt.Printf("创建时间: %s | ID: %d\n", timeStr, record.ID)

		if i < len(records)-1 {
			fmt.Println(strings.Repeat("-", 50))
		}
	}
	fmt.Println(strings.Repeat("=", 80))

	// 显示统计信息
	today := time.Now().Format("2006-01-02")
	todayCount := 0
	for _, record := range records {
		if record.CreatedAt.Format("2006-01-02") == today {
			todayCount++
		}
	}

	fmt.Printf("统计: 今日新增 %d 条记录 | 总计 %d 条记录\n", todayCount, len(records))
}

// DisplayLinkStats 显示链接统计
func DisplayLinkStats() {
	linkRecords := services.GetLinkRecords()
	fmt.Printf("\n🔗 链接库统计: 共 %d 个链接\n", len(linkRecords))
	if len(linkRecords) > 0 {
		fmt.Println("最新链接:")
		for i := len(linkRecords) - 1; i >= max(0, len(linkRecords)-3); i-- {
			record := linkRecords[i]
			fmt.Printf("  • %s → %s\n", record.Title, record.Link)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func createDefaultTemplate() {
	os.WriteFile("templates/index.html", []byte(templates.IndexTemplate), 0644)
}

// 创建关于页面模板
func createAboutTemplate() {
	os.WriteFile("templates/about.html", []byte(templates.AboutTemplate), 0644)
}

// 创建统计页面模板
func createStatsTemplate() {
	os.WriteFile("templates/stats.html", []byte(templates.StatsTemplate), 0644)
}

// 创建链接页面模板
func createLinksTemplate() {
	os.WriteFile("templates/links.html", []byte(templates.LinksTemplate), 0644)
}
