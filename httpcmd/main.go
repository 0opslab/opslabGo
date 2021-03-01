package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)


var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

var akey string
var cmdlist map[string]interface{}
var shell = false

func main() {
	sysType := runtime.GOOS
	jsonConfigFile := ""

	if sysType == "linux" {
		jsonConfigFile = "/etc/httpcmd.conf"
		shell = true
	}
	if sysType == "windows" {
		jsonConfigFile = "C:/httpcmd.conf"
	}
	tempMap := make(map[string]interface{})
	content, err := ioutil.ReadFile(jsonConfigFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &tempMap)
	akey = tempMap["akey"].(string)
	logPath := tempMap["logpath"].(string)
	cmdlist = tempMap["cmdlist"].(map[string]interface{})

	//for k, v := range cmdlist {
	//	fmt.Println("============", k, "==================")
	//	//responseBody := RunCmd()
	//	//fmt.Println(responseBody)
	//	command := v.(string)
	//	responseBody := RunCmd(command, shell)
	//	fmt.Println(responseBody)
	//}

	//flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()
	defer logger.Init("LoggerExample", *verbose, true, lf).Close()



	http.HandleFunc("/httpcmd", httpCmd)
	//设置监听的端口
	errs := http.ListenAndServe(":10000", nil)
	if errs != nil {
		logger.Info("start error :", errs)
	}
	logger.Info("service is started")

}
func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func httpCmd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	item := r.Form.Get("item")


	timeStr:=time.Now().Format("2006010215")
	fmt.Println(timeStr)


	responseBody := "Invalid request"
	//通过简单的md5校验实现简单层面的安全控制
	if key == md5V(akey+":"+timeStr+":"+item) {
		if realCmd, ok := cmdlist[item]; ok {
			command := realCmd.(string)
			responseBody = RunCmd(command, shell)
		}
	}
	fmt.Fprintln(w, responseBody)

}

func RunCmd(cmd string, shell bool) string {
	logger.Info("cmd=> ",cmd)
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			logger.Info("InvalidCmd=> ",cmd)
			//panic("some error found")
		}
		return string(out)
	} else {
		out, err := exec.Command("cmd", "/C", cmd).Output()
		if err != nil {
			logger.Info("InvalidCmd=> ",cmd)
			//panic("some error found")
		}
		return string(out)
	}
}
