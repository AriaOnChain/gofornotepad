// main.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"nav/templates"
)

// 记录结构体
type Record struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var records []Record
var dataFile = "data/records.json"

// 链接记录结构体
type LinkRecord struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var linkRecords []LinkRecord
var linkDataFile = "data/links.json"

func main() {
	// 创建数据目录
	os.MkdirAll("data", 0755)
	os.MkdirAll("templates", 0755)

	// 加载数据
	loadRecords()
	loadLinkRecords()
	// 显示当前待办
	displayPendingTasks()

	// 设置路由
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/api/update", updateHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/stats", statsHandler)

	// 添加链接页面路由
	http.HandleFunc("/links", linksHandler)
	http.HandleFunc("/links/add", addLinkHandler)
	http.HandleFunc("/links/delete", deleteLinkHandler)
	http.HandleFunc("/links/api/update", updateLinkHandler)

	fmt.Println("🚀 记录本应用启动在 http://localhost:8080")
	fmt.Println("📝 功能：添加记录、查看记录、删除记录、搜索记录")
	fmt.Println("规划助力来也~！")
	fmt.Println("你好呀！！！ 今天是美好的一天 开心点！")
	fmt.Println("把想法写下来 然后慢慢行动吧！你是最棒的！")
	fmt.Println("立马行动！gogogo!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 首页处理器
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 渲染模板
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// 如果模板不存在，创建默认模板
		createDefaultTemplate()
		tmpl, _ = template.ParseFiles("templates/index.html")
	}

	searchQuery := r.URL.Query().Get("q")
	var displayRecords []Record

	if searchQuery != "" {
		// 执行搜索
		for _, record := range records {
			if strings.Contains(strings.ToLower(record.Title), strings.ToLower(searchQuery)) ||
				strings.Contains(strings.ToLower(record.Content), strings.ToLower(searchQuery)) {
				displayRecords = append(displayRecords, record)
			}
		}
	} else {
		displayRecords = records
	}

	data := map[string]interface{}{
		"Title":       "demo",
		"Title1":      "待实现的零碎想法~",
		"Records":     displayRecords,
		"TotalCount":  len(records),
		"SearchQuery": searchQuery,
		"Success":     r.URL.Query().Get("success"),
		"Error":       r.URL.Query().Get("error"),
	}

	tmpl.Execute(w, data)
}

// 关于页面处理器
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		// 如果模板不存在，创建默认关于模板
		createAboutTemplate()
		tmpl, _ = template.ParseFiles("templates/about.html")
	}

	data := map[string]interface{}{
		"Title":        "关于我们",
		"AppName":      "零碎想法记录本",
		"Version":      "1.1.0",
		"TotalRecords": len(records),
	}

	tmpl.Execute(w, data)
}

// 统计页面处理器
func statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 计算统计信息
	stats := calculateStats()

	tmpl, err := template.ParseFiles("templates/stats.html")
	if err != nil {
		// 如果模板不存在，创建默认统计模板
		createStatsTemplate()
		tmpl, _ = template.ParseFiles("templates/stats.html")
	}

	data := map[string]interface{}{
		"Title":   "数据统计",
		"Stats":   stats,
		"Records": records,
	}

	tmpl.Execute(w, data)
}

// 链接页面处理器
func linksHandler(w http.ResponseWriter, r *http.Request) {
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
	var displayRecords []LinkRecord

	if searchQuery != "" {
		// 执行搜索
		for _, record := range linkRecords {
			if strings.Contains(strings.ToLower(record.Title), strings.ToLower(searchQuery)) ||
				strings.Contains(strings.ToLower(record.Content), strings.ToLower(searchQuery)) ||
				strings.Contains(strings.ToLower(record.Link), strings.ToLower(searchQuery)) {
				displayRecords = append(displayRecords, record)
			}
		}
	} else {
		displayRecords = linkRecords
	}

	data := map[string]interface{}{
		"Title":       "链接收藏",
		"Title1":      "🔗 我的链接库",
		"Records":     displayRecords,
		"TotalCount":  len(linkRecords),
		"SearchQuery": searchQuery,
		"Success":     r.URL.Query().Get("success"),
		"Error":       r.URL.Query().Get("error"),
	}

	tmpl.Execute(w, data)
}

// 添加链接处理器
func addLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))
	link := strings.TrimSpace(r.FormValue("link"))

	// 验证链接格式
	if link != "" && !isValidURL(link) {
		// 如果不是完整URL，添加https前缀
		if !strings.Contains(link, "://") {
			link = "https://" + link
		}
	}

	// 创建新链接记录
	newRecord := LinkRecord{
		ID:        getNextLinkID(),
		Title:     title,
		Content:   content,
		Link:      link,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	linkRecords = append(linkRecords, newRecord)
	saveLinkRecords()

	fmt.Printf("\n✅ 新链接已添加！标题: %s\n", title)
	displayLinkStats()

	http.Redirect(w, r, "/links?success=链接添加成功", http.StatusSeeOther)
}

// 删除链接处理器
func deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/links?error=无效的记录ID", http.StatusSeeOther)
		return
	}

	// 查找并删除记录
	for i, record := range linkRecords {
		if record.ID == id {
			linkRecords = append(linkRecords[:i], linkRecords[i+1:]...)
			saveLinkRecords()
			http.Redirect(w, r, "/links?success=链接删除成功", http.StatusSeeOther)
			return
		}
	}

	http.Redirect(w, r, "/links?error=链接未找到", http.StatusSeeOther)
}

// 更新链接处理器
func updateLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Link    string `json:"link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	id := requestData.ID
	title := strings.TrimSpace(requestData.Title)
	content := strings.TrimSpace(requestData.Content)
	link := strings.TrimSpace(requestData.Link)

	if title == "" {
		http.Error(w, "标题不能为空", http.StatusBadRequest)
		return
	}

	// 验证链接格式
	if link != "" && !isValidURL(link) {
		if !strings.Contains(link, "://") {
			link = "https://" + link
		}
	}

	// 更新记录
	found := false
	for i := range linkRecords {
		if linkRecords[i].ID == id {
			linkRecords[i].Title = title
			linkRecords[i].Content = content
			linkRecords[i].Link = link
			linkRecords[i].UpdatedAt = time.Now()
			saveLinkRecords()
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "链接未找到", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "链接更新成功",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Printf("\n✏️ 链接已更新！ID: %d\n", id)
	displayLinkStats()
}

// 链接相关工具函数
func loadLinkRecords() {
	data, err := os.ReadFile(linkDataFile)
	if err != nil {
		linkRecords = []LinkRecord{}
		return
	}
	json.Unmarshal(data, &linkRecords)
}

func saveLinkRecords() {
	data, err := json.MarshalIndent(linkRecords, "", "  ")
	if err != nil {
		log.Printf("保存链接数据失败: %v", err)
		return
	}
	os.WriteFile(linkDataFile, data, 0644)
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

func displayLinkStats() {
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

// 计算统计数据
func calculateStats() map[string]interface{} {
	stats := make(map[string]interface{})

	// 基本统计
	stats["TotalRecords"] = len(records)

	// 今日统计
	today := time.Now().Format("2006-01-02")
	todayCount := 0
	for _, record := range records {
		if record.CreatedAt.Format("2006-01-02") == today {
			todayCount++
		}
	}
	stats["TodayRecords"] = todayCount

	// 本周统计
	weekCount := 0
	year, week := time.Now().ISOWeek()
	for _, record := range records {
		recordYear, recordWeek := record.CreatedAt.ISOWeek()
		if recordYear == year && recordWeek == week {
			weekCount++
		}
	}
	stats["WeekRecords"] = weekCount

	// 月度统计
	month := time.Now().Format("2006-01")
	monthCount := 0
	for _, record := range records {
		if record.CreatedAt.Format("2006-01") == month {
			monthCount++
		}
	}
	stats["MonthRecords"] = monthCount

	// 内容长度统计
	totalChars := 0
	longestRecord := Record{}
	shortestRecord := Record{}

	if len(records) > 0 {
		longestRecord = records[0]
		shortestRecord = records[0]

		for _, record := range records {
			contentLength := len(record.Content)
			totalChars += contentLength

			if contentLength > len(longestRecord.Content) {
				longestRecord = record
			}
			if contentLength < len(shortestRecord.Content) && contentLength > 0 {
				shortestRecord = record
			}
		}
	}

	stats["AverageChars"] = 0
	if len(records) > 0 {
		stats["AverageChars"] = totalChars / len(records)
	}
	stats["LongestRecord"] = longestRecord
	stats["ShortestRecord"] = shortestRecord
	stats["TotalChars"] = totalChars

	return stats
}

// 添加记录处理器
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))

	if title == "" {
		http.Redirect(w, r, "/?error=标题不能为空", http.StatusSeeOther)
		return
	}

	// 创建新记录
	newRecord := Record{
		ID:      getNextID(),
		Title:   title,
		Content: content,
		// Category:  category,
		CreatedAt: time.Now(),
	}

	records = append(records, newRecord)
	saveRecords()

	fmt.Println("\n✅ 新记录已添加！")
	displayPendingTasks()

	http.Redirect(w, r, "/?success=记录添加成功", http.StatusSeeOther)
}

// 删除记录处理器
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/?error=无效的记录ID", http.StatusSeeOther)
		return
	}

	// 查找并删除记录
	for i, record := range records {
		if record.ID == id {
			records = append(records[:i], records[i+1:]...)
			saveRecords()
			http.Redirect(w, r, "/?success=记录删除成功", http.StatusSeeOther)
			return
		}
	}

	http.Redirect(w, r, "/?error=记录未找到", http.StatusSeeOther)
}

// 搜索处理器
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/?q="+query, http.StatusSeeOther)
}

// 更新处理
func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 解析JSON请求
	var requestData struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	id := requestData.ID
	title := strings.TrimSpace(requestData.Title)
	content := strings.TrimSpace(requestData.Content)

	if title == "" {
		http.Error(w, "标题不能为空", http.StatusBadRequest)
		return
	}

	// 更新记录
	found := false
	for i := range records {
		if records[i].ID == id {
			records[i].Title = title
			records[i].Content = content
			records[i].UpdatedAt = time.Now()
			saveRecords()
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "记录未找到", http.StatusNotFound)
		return
	}

	// 返回成功响应
	response := map[string]interface{}{
		"success": true,
		"message": "记录更新成功",
		"record":  records,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Printf("\n✏️ 记录已更新！ID: %d\n", id)
	displayPendingTasks()
}

// 显示待办任务
func displayPendingTasks() {
	if len(records) == 0 {
		fmt.Println("\n列表暂无记录")
		fmt.Println("提示: 在浏览器中添加你的第一条记录吧！")
		return
	}

	fmt.Printf("\n 当前列表 (共 %d 项):\n", len(records))
	fmt.Println(strings.Repeat("=", 50))

	for i, record := range records {
		// 格式化时间
		timeStr := record.CreatedAt.Format("01-02 15:04")

		// 显示内容预览（如果有）
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

// 工具函数
func loadRecords() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		// 文件不存在，使用空记录
		records = []Record{}
		return
	}

	json.Unmarshal(data, &records)
}

func saveRecords() {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		log.Printf("保存数据失败: %v", err)
		return
	}

	os.WriteFile(dataFile, data, 0644)
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
