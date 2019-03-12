package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))

}

func getCurrentIP(r http.Request) (string) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		return r.RemoteAddr
	}
	return ip
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 实现多文件接收
	//上传结果以以json格式返回
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s := ""
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		newfile := UniqueId() + path.Ext(part.FileName())
		loginfo := fmt.Sprintf("%s %s uploadfile %s ==> %s", time.Now().Format("2006-01-02 15:04:05"), getCurrentIP(*r), newfile, part.FileName())
		log.Println(loginfo)
		if part.FileName() != "" {
			dst, _ := os.Create("c:/var/" + newfile)
			defer dst.Close()
			io.Copy(dst, part)
			s += fmt.Sprintf("'%s':'%s',", part.FormName(), newfile)
		}
	}
	res := fmt.Sprintf("{'code':'success',results:{%s}}", strings.Trim(s, ","))
	w.Header().Add("Content-Type","application/json;charset:utf-8;")
	fmt.Fprintf(w, res)
}

func main() {
	log.Print("Server is starting")
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Service:", err)
	}

}
