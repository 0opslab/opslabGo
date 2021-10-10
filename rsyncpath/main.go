package main

//
// @instruction
//		利用go实现的HTTP版的文件上传，上传接口以json方式返回
// @上传方式
// 		通用的http文件上传方式
// @实现原理
//		当前程序接受到上传请求处理并响应,同时利用go的协程同步到配置的其他客户端上
//
// @配置json
//	{
//    "addr":"0.0.0.0:9090",
//    "path":"c:/var/upload/wwww/",
//    "fileNameLength":11,
//    "rysncAddr":[
//        "http://localhost:9091/rsync",
//        "http://localhost:9092/rsync"
//    ]
//	}
// @说明
// 		普通上传方式的文件都会存储在配置文件指定的目录下，如果想在改目录下新建文件夹并存储
// 		到新建文件夹下的可以同http head字段path添加目录(目录名需要BASE64)

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	Autngo "github.com/0opslab/autngo"
	"github.com/fsnotify/fsnotify"
)

type ServerConfig struct {
	//监听地址和端口
	ADDR string `json:'ADDR'`
	//文件写入路径
	PATH string `json:'PATH'`
	//同步的地址
	RYSNCADDR []string `json:'RYSNCADDR'`
}

var APP = ServerConfig{}
var DIRPATH_INFP  []string
func main() {

	// confile := flag.String("conf", "", "the configuration file")
	// flag.Parse()
	// if *confile == "" {
	// 	fmt.Println("Please specify the configuration file")
	// 	return
	// }
	// file, _ := os.Open(*confile)
	// defer file.Close()
	// decoder := json.NewDecoder(file)

	// err := decoder.Decode(&APP)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	//@TODO-FORTEST
	var jsonstr = `{"addr":"0.0.0.0:9090","path":"C:/workspace/doc","fileNameLength":10,
		"rysncAddr":["http://localhost:9091/rsync","http://localhost:9092/rsync"]}`
	if err := json.Unmarshal([]byte(jsonstr), &APP); err != nil {
		panic("ErrorConfig")
	}
	DIRPATH_INFP = GetDirPathInfo(APP.PATH)
	watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal("NewWatcher failed: ", err)
  }
  defer watcher.Close()
  done := make(chan bool)
  go func() {
    defer close(done)
    for {
      select {
      case event, ok := <-watcher.Events:
        if !ok {
          return
        }
				//event.Op 时间类型
				//CREATE
				//REMOVE
				//WRITE
				//RENAME
				//CHMOD
        log.Printf("%s %s\n", event.Name, event.Op)
				//if event.Op&fsnotify.Write == fsnotify.Write {
				//只需要同步写入文件即可其实
				ss_tem := GetDirPathInfo(APP.PATH)
				change_list := Autngo.SliceHelper.Difference(DIRPATH_INFP, ss_tem)
				DIRPATH_INFP = ss_tem
				for _, v := range change_list {
					Rsyncfile(v)
				}
				//}
      case err, ok := <-watcher.Errors:
        if !ok {
          return
        }
        log.Println("error:", err)
      }
    }
  }()

  err = watcher.Add(APP.PATH)
  if err != nil {
    log.Fatal("Add failed:", err)
  }

	
	log.Println("Server is starting:" + APP.ADDR)
	log.Println("Server UploadPath:" + APP.PATH)
	log.Print("Server Rysnc Addr:" + strings.Replace(strings.Trim(fmt.Sprint(APP.RYSNCADDR), "[]"), " ", ",", -1))

	http.HandleFunc("/rsync", RsyncHandler)
	if err := http.ListenAndServe(APP.ADDR, nil); err != nil {
		fmt.Println("Server starting error")
	}
	
}


//获取目录信息
func GetDirPathInfo(dir_path string) []string{
	fileInfo := func(fileName string) string{
		isdir := 1
		mtime := int64(0)
		fileSize := int64(0)
		if(Autngo.FileHepler.IsFile(fileName)){
			isdir = 0
			mtime = Autngo.FileHepler.GetFileModTime(fileName)
			fileSize,_ = Autngo.FileHepler.FileSize(fileName)
		}
		fileName = trim_pathFile(fileName)
		result := "{\"isdir\":\"%v\",\"mtime\":\"%v\",\"fileName\":\"%v\",\"fileSize\":\"%v\"}"
		s:= fmt.Sprintf(result, isdir,mtime,fileName,fileSize)
		return s
	}

	//dirPath := "C:/workspace/useful-documents/doc/linux"
	var ss []string
	filepath.Walk(dir_path, func(filename string, fi os.FileInfo, err error) error {
		info := fileInfo(filename)
		ss = append(ss,info)
		return nil
	})
	return ss

}

func trim_pathFile(file_name string)string {
	re2, _ := regexp.Compile("\\\\{1,}")
	strs := re2.ReplaceAllString(file_name, "/")
	re3, _ := regexp.Compile("/{2,}")
	return re3.ReplaceAllString(strs, "/")
}
//同步文件
func Rsyncfile(fileInfo string){

	fmt.Printf("RsyncFile :%v\n", fileInfo)
	ffss := func (url string, dstPath string, files string) {
		bodyBuffer := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuffer)
		_, fileName := filepath.Split(files)
		fileWriter, _ := bodyWriter.CreateFormFile("rsyncfile", fileName)

		file, _ := os.Open(files)
		defer file.Close()

		io.Copy(fileWriter, file)

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		if req, err := http.NewRequest("POST", url, bodyBuffer); err == nil {
			req.Header.Set("Content-Type", contentType)
			req.Header.Set("Path", base64.StdEncoding.EncodeToString([]byte(dstPath)))
			if resp, errsp := http.DefaultClient.Do(req); errsp == nil {
				resp_body, _ := ioutil.ReadAll(resp.Body)
				log.Println(fmt.Sprintf("Clientrsyncfile %s %s ", resp.Status, string(resp_body)))
			} else {
				log.Println(fmt.Sprintf("Clientrsyncfile Error %s %s ", url, files))
			}

		} else {
			log.Println(fmt.Sprintf("Clientrsyncfile Error %s %s ", url, files))
		}

	}
	
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(fileInfo), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fileName := fmt.Sprintf("%v",mapResult["fileName"])
	upload_pathfile := strings.Replace(fileName, APP.PATH,"", 1)
	fmt.Println(fileName,"======>",upload_pathfile)


	for _, v := range APP.RYSNCADDR {
		//fmt.Printf("RsyncFile :%v  %v\n",v, fileInfo)
		go ffss(v, upload_pathfile, fileName)
	}
	
}



func getCurrentIP(r http.Request) (string) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		return r.RemoteAddr
	}
	return ip
}

func RsyncHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("rsyncfile")
	if err != nil {
		log.Println(fmt.Sprintf("%s rsyncfile %s %s ", getCurrentIP(*r), header.Filename, "FormParseError"))
		res := fmt.Sprintf("{'code':'error'}")
		w.Header().Add("Content-Type", "application/json;charset:utf-8;")
		fmt.Fprintf(w, res)
		return
	}
	defer file.Close()

	dstFile := APP.PATH + r.Header.Get("Path")
	if Autngo.FileHepler.FileIsExist(dstFile) {
		log.Println(fmt.Sprintf("%s rsyncfile %s %s ", getCurrentIP(*r), header.Filename, "FileExists"))
		res := fmt.Sprintf("{'code':'error'}")
		w.Header().Add("Content-Type", "application/json;charset:utf-8;")
		fmt.Fprintf(w, res)
		return
	}

	cur, err := os.Create(dstFile);
	if err != nil {
		log.Println(fmt.Sprintf("%s rsyncfile %s %s ", getCurrentIP(*r), header.Filename, "CreateError"))
		res := fmt.Sprintf("{'code':'error'}")
		w.Header().Add("Content-Type", "application/json;charset:utf-8;")
		fmt.Fprintf(w, res)
		return
	}
	defer cur.Close()

	res := fmt.Sprintf("{'code':'error'}")
	loginfo := ""
	_, erro := io.Copy(cur, file)
	if erro != nil {
		loginfo = fmt.Sprintf("%s rsyncfile %s  %s", getCurrentIP(*r), header.Filename, "WriteError")
	} else {
		loginfo = fmt.Sprintf("%s rsyncfile %s  %s", getCurrentIP(*r), header.Filename, "RysncSuccess")
		res = fmt.Sprintf("{'code':'success'}")

	}
	log.Println(loginfo)
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	fmt.Fprintf(w, res)
}




