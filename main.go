package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"nav/handlers"
	"nav/utils"
)

func main() {
	// 创建数据目录
	os.MkdirAll("data", 0755)
	os.MkdirAll("templates", 0755)

	url := "http://localhost:8080"

	// 初始化服务
	handlers.InitServices()

	// 显示当前待办
	handlers.DisplayPendingTasks()

	// 设置路由
	setupRoutes()

	utils.OpenBrowser(url)

	fmt.Println("🚀 记录本应用启动在 http://localhost:8080")
	fmt.Println("📝 功能：添加记录、查看记录、删除记录、搜索记录")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	// 主要页面路由
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)
	http.HandleFunc("/links", handlers.LinksHandler)

	// 记录操作路由
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	http.HandleFunc("/api/update", handlers.UpdateHandler)

	// 链接操作路由
	http.HandleFunc("/links/add", handlers.AddLinkHandler)
	http.HandleFunc("/links/delete", handlers.DeleteLinkHandler)
	http.HandleFunc("/links/api/update", handlers.UpdateLinkHandler)
}
