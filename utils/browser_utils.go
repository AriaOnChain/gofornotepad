package utils

import (
	"log"
	"os/exec"
	"runtime"
)

// OpenBrowser 打开浏览器
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		log.Printf("请手动打开浏览器访问: %s", url)
	}

	if err != nil {
		log.Printf("自动打开浏览器失败: %v", err)
	}
}
