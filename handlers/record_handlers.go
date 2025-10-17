package handlers

import (
	"encoding/json"
	"fmt"
	"nav/services"
	"net/http"
	"strconv"
	"strings"
)

// AddHandler 添加记录处理器
func AddHandler(w http.ResponseWriter, r *http.Request) {
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

	services.AddRecord(title, content)

	fmt.Printf("\n✅ 新记录已添加！标题: %s\n", title)
	DisplayPendingTasks()

	http.Redirect(w, r, "/?success=记录添加成功", http.StatusSeeOther)
}

// DeleteHandler 删除记录处理器
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
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

	if services.DeleteRecord(id) {
		http.Redirect(w, r, "/?success=记录删除成功", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/?error=记录未找到", http.StatusSeeOther)
	}
}

// UpdateHandler 更新记录处理器
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

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

	if services.UpdateRecord(id, title, content) {
		response := map[string]interface{}{
			"success": true,
			"message": "记录更新成功",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		fmt.Printf("\n✏️ 记录已更新！ID: %d\n", id)
		DisplayPendingTasks()
	} else {
		http.Error(w, "记录未找到", http.StatusNotFound)
	}
}

// SearchHandler 搜索处理器
func SearchHandler(w http.ResponseWriter, r *http.Request) {
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
