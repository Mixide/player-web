package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/dhowden/tag"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Music struct {
	Id     uint   `gorm:"primaryKey"`
	Artist string `json:"artist"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}
type listResponse struct {
	Count  int     `json:"count"`
	Musics []Music `json:"musics"`
}

var saveDir = "./assets/"
var serverIP = "localhost:7986"
var Db *gorm.DB

func dbInit() {
	var err error
	Db, err = gorm.Open(sqlite.Open(saveDir + "database.db"))
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Music{})
}

func main() {
	// 设置文件服务的根目录

	dbInit()
	// 创建HTTP处理器函数
	http.HandleFunc("/assets/", getMusic)
	http.HandleFunc("/getlist", getList)
	http.HandleFunc("/upload", uploadMusic)
	// 启动HTTP服务器
	log.Printf("Server started on :%s\n", serverIP)
	err := http.ListenAndServe(serverIP, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func uploadMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域名访问
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			log.Printf("get data fail: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		newFile, err := os.Create(saveDir + head.Filename)
		if err != nil {
			log.Printf("create file fail: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			os.Remove(saveDir + head.Filename)
			log.Printf("save data fail: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = newFile.Seek(0, io.SeekStart)
		if err != nil {
			os.Remove(saveDir + head.Filename) // 删除文件
			log.Printf("seek file fail: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		meta, err := tag.ReadFrom(newFile)
		if err != nil {
			os.Remove(saveDir + head.Filename)
			log.Printf("get metadata fail: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		url := "http://" + serverIP + "/assets/" + head.Filename
		Db.Create(&Music{Name: meta.Title(), Artist: meta.Artist(), Url: url})
		response := map[string]string{"message": "File uploaded successfully"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func getMusic(w http.ResponseWriter, r *http.Request) {
	// 获取请求的文件名
	requestedFile := r.URL.Path
	// 构建完整的文件路径
	fileName := path.Base(requestedFile)
	filePath := filepath.Join(saveDir, fileName)
	// 设置Content-Disposition头，指定文件名和下载方式
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))

	// 设置Content-Type头，根据需要设置适当的MIME类型
	w.Header().Set("Content-Type", "application/octet-stream")

	// 读取文件并将其写入响应
	http.ServeFile(w, r, filePath)
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域名访问
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var allMusic []Music
	Db.Find(&allMusic)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(listResponse{Count: len(allMusic), Musics: allMusic})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
