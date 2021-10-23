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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	Autngo "github.com/0opslab/autngo"
	"github.com/fsnotify/fsnotify"
)

type ServerConfig struct {
	//监听地址和端口
	ADDR string `json:'ADDR'`
	//文件写入路径
	PATH string `json:'PATH'`
	//模式 server(接受同步文件模式) client(监控指定目录并将变化同步至server模式主机下)
	MODE string `json:'MODE'`
	//同步的地址
	RYSNCADDR []string `json:'RYSNCADDR'`
}

var APP = ServerConfig{}
var DIRPATH_INFP  []string
var STATUS_SUCCESS = 10000
var STATUS_FAIL = 10001

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
	// var jsonstr = `{"addr":"0.0.0.0:9090","path":"C:/workspace/doc1","mode":"server",
	// 	"rysncAddr":["http://localhost:9091/rsync"]}`

	var jsonstr = `{"addr":"0.0.0.0:9091","path":"C:/workspace/doc","mode":"client",
		"rysncAddr":["http://localhost:9090/rsync"]}`
	if err := json.Unmarshal([]byte(jsonstr), &APP); err != nil {
		panic("ErrorConfig")
	}

	if APP.MODE == "client"{
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
					if event.Op&fsnotify.Write == fsnotify.Write {
						//log.Printf("%s %s\n", event.Name, event.Op)
						//只需要同步写入文件即可其实
						ss_tem := GetDirPathInfo(APP.PATH)
						change_list := Autngo.SliceHelper.Difference(ss_tem,DIRPATH_INFP)
						DIRPATH_INFP = ss_tem
						for _, v := range change_list {
							//go Rsyncfile(v)
							var mapResult map[string]interface{}
							err := json.Unmarshal([]byte(v), &mapResult)
							if err != nil {
								fmt.Println("JsonToMapDemo err: ", err)
							}
							fileName := fmt.Sprintf("%v",mapResult["fileName"])
							upload_pathfile := strings.Replace(fileName, APP.PATH,"", 1)
	
							for _, v := range APP.RYSNCADDR {
								go Rsyncfile(v, upload_pathfile, fileName)
							}
						}
					}
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
	}
	
	log.Println("Server is starting:",APP.ADDR,"\tServer-UploadPath:",APP.PATH)
	log.Print("Server Rysnc Addr:" + strings.Replace(strings.Trim(fmt.Sprint(APP.RYSNCADDR), "[]"), " ", ",", -1))
	http.HandleFunc("/rsync", RsyncHandler)
	if err := http.ListenAndServe(APP.ADDR, nil); err != nil {
		fmt.Println("Server starting error")
	}
	
}

//同步文件
func Rsyncfile(url string, dstPath string, files string){
		info := fmt.Sprintf("RsyncFile :%v path:%v",url,dstPath)
		log.Print(info)
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
			req.Header.Set("Path", Autngo.StringHelper.Base64Encode(dstPath))
			if resp, errsp := http.DefaultClient.Do(req); errsp == nil {
				resp_body, _ := ioutil.ReadAll(resp.Body)
				log.Println(fmt.Sprintf("Clientrsyncfile %s %s ", resp.Status, string(resp_body)))
			} else {
				log.Println(fmt.Sprintf("Clientrsyncfile Error %s %s %v", url, files,errsp))
			}
		} else {
			log.Println(fmt.Sprintf("Clientrsyncfile Error %s %s %v", url, files,err))
		}
}

//获取目录信息
func GetDirPathInfo(dir_path string) []string{
	fileInfo := func(fileName string) string{
		fileSize := int64(0)
		mtime := Autngo.FileHepler.GetFileModTime(fileName)
		fileSize,_ = Autngo.FileHepler.FileSize(fileName)
		fileName = Autngo.FileHepler.TrimPathFile(fileName)
		result := "{\"mtime\":\"%v\",\"fileName\":\"%v\",\"fileSize\":\"%v\"}"
		s:= fmt.Sprintf(result,mtime,fileName,fileSize)
		return s
	}
	var ss []string
	filepath.Walk(dir_path, func(filename string, fi os.FileInfo, err error) error {
		if(Autngo.FileHepler.IsFile(filename)){
			info := fileInfo(filename)
			ss = append(ss,info)
		}
		return nil
	})
	return ss
}


//处理同步文件
func RsyncHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("rsyncfile")
	if err != nil {
		Autngo.HttpHelper.HttpResponseCode(w, STATUS_FAIL,"FormParseError")
		return
	}
	defer file.Close()

	dstFile := Autngo.FileHepler.TrimPathFile(APP.PATH + Autngo.StringHelper.Base64Decode(r.Header.Get("Path")))
	log.Println(fmt.Sprintf("%s rsyncfile  %s", Autngo.HttpHelper.GetRemoteIP(r), dstFile))
	dirpath := filepath.Dir(dstFile)
	if !Autngo.FileHepler.FileIsExist(dirpath) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			Autngo.HttpHelper.HttpResponseCodeData(w,STATUS_FAIL,"CreateFile",dstFile)
			return
		}
	}
	cur, err := os.Create(dstFile);
	if err != nil {
		log.Println(fmt.Sprintf("%s rsyncfile %s %s ", Autngo.HttpHelper.GetRemoteIP(r), header.Filename, "CreateError"))
		Autngo.HttpHelper.HttpResponseCodeData(w,STATUS_FAIL,"CreateError",dstFile)
		return
	}
	defer cur.Close()

	res := STATUS_FAIL
	result :="error"
	_, erro := io.Copy(cur, file)
	if erro != nil {
		result = "WriteError"
	} else {
		result ="RysncSuccess"
		res = STATUS_SUCCESS
	}
	log.Println(fmt.Sprintf("%s rsyncfile %s  %s", Autngo.HttpHelper.GetRemoteIP(r), header.Filename, result))
	Autngo.HttpHelper.HttpResponseCode(w,res,result)
}




