package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dhowden/tag"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Music struct {
	Id     uint   `gorm:"primaryKey"`
	Artist string `json:"artist"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}
type Music2User struct {
	UserId  uint
	MusicId uint
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
	Db.AutoMigrate(&Music{}, &Music2User{})
}

func main() {
	// 设置文件服务的根目录

	dbInit()
	// 创建HTTP处理器函数
	router := mux.NewRouter()
	router.HandleFunc("/assets/{filename}", getMusic).Methods("GET")
	router.HandleFunc("/getlist/{id}", getList)
	router.HandleFunc("/upload/{id}", uploadMusic)
	// 启动HTTP服务器
	log.Printf("Server started on :%s\n", serverIP)
	err := http.ListenAndServe(serverIP, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func uploadMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域名访问
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		log.Printf("wrong userid: %s\n", err.Error())
	}
	url := "http://" + serverIP + "/assets/" + head.Filename
	newMusic := Music{Name: meta.Title(), Artist: meta.Artist(), Url: url}
	Db.Model(&Music{}).Create(&newMusic)
	Db.Model(&Music2User{}).Create(&Music2User{UserId: uint(id), MusicId: newMusic.Id})
	response := map[string]string{"message": "File uploaded successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getMusic(w http.ResponseWriter, r *http.Request) {
	// 获取请求的文件名
	vars := mux.Vars(r)
	filename := vars["filename"]
	filePath := filepath.Join(saveDir, filename)
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
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		log.Printf("wrong userid: %s\n", err.Error())
	}
	var (
		musicIds []uint
		allMusic []Music
	)
	Db.Model(&Music2User{}).Where("user_id = ?", uint(id)).Pluck("music_id", &musicIds)
	Db.Model(&Music{}).Find(&allMusic).Where("id = ?", musicIds)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listResponse{Count: len(allMusic), Musics: allMusic})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
