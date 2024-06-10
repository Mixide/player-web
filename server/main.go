package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// 设置文件服务的根目录
	fileServerRoot := "./"

	// 创建HTTP处理器函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求的文件名
		requestedFile := r.URL.Path
		// 构建完整的文件路径
		filePath := filepath.Join(fileServerRoot, requestedFile)
		// 设置Content-Disposition头，指定文件名和下载方式
		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))

		// 设置Content-Type头，根据需要设置适当的MIME类型
		w.Header().Set("Content-Type", "application/octet-stream")

		// 读取文件并将其写入响应
		http.ServeFile(w, r, filePath)
	})

	// 启动HTTP服务器
	log.Println("Server started on :7986")
	err := http.ListenAndServe(":7986", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
