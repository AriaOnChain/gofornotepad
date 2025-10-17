package handlers

import (
	"encoding/json"
	"fmt"
	"nav/services"
	"net/http"
	"strconv"
	"strings"
)

// AddLinkHandler 添加链接处理器
func AddLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))
	link := strings.TrimSpace(r.FormValue("link"))

	services.AddLinkRecord(title, content, link)

	fmt.Printf("\n✅ 新链接已添加！标题: %s\n", title)
	DisplayLinkStats()

	http.Redirect(w, r, "/links?success=链接添加成功", http.StatusSeeOther)
}

// DeleteLinkHandler 删除链接处理器
func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
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

	if services.DeleteLinkRecord(id) {
		http.Redirect(w, r, "/links?success=链接删除成功", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/links?error=链接未找到", http.StatusSeeOther)
	}
}

// UpdateLinkHandler 更新链接处理器
func UpdateLinkHandler(w http.ResponseWriter, r *http.Request) {
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

	if services.UpdateLinkRecord(id, title, content, link) {
		response := map[string]interface{}{
			"success": true,
			"message": "链接更新成功",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		fmt.Printf("\n✏️ 链接已更新！ID: %d\n", id)
		DisplayLinkStats()
	} else {
		http.Error(w, "链接未找到", http.StatusNotFound)
	}
}
