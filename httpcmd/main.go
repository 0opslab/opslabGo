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


var akey string
var cmdlist map[string]interface{}
var shell = false
var logPath string

/**
 * 加载配置文件信息
 */
func loadConfig(){
	sysType := runtime.GOOS
	jsonConfigFile := ""

	if sysType == "linux" {
		jsonConfigFile = "/etc/httpcmd.conf"
		shell = true
	}
	if sysType == "windows" {
		jsonConfigFile = "./httpcmd.conf"
	}
	tempMap := make(map[string]interface{})
	content, err := ioutil.ReadFile(jsonConfigFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &tempMap)
	if err != nil {
		panic(err)
	}
	akey = tempMap["akey"].(string)
	cmdlist = tempMap["cmdlist"].(map[string]interface{})
	logPath = tempMap["logpath"].(string)
	logger.Info("httpCmdReload=> success")
}
func RunCmd(cmd string, shell bool) string {
	logger.Info("cmd=> ",cmd)
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			logger.Info("InvalidCmd=> ",cmd)
		}
		return string(out)
	} else {
		out, err := exec.Command("cmd", "/C", cmd).Output()
		if err != nil {
			logger.Info("InvalidCmd=> ",cmd)
		}
		return string(out)
	}
}

/**
*	返回所有的map键值
*/
func Map_Get_Keys(m map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}


func main() {
	
	loadConfig()

	//init logger
	var verbose = flag.Bool("verbose", false, "print info level logs to stdout")
	lf, err := os.OpenFile(logPath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()
	defer logger.Init("LoggerExample", *verbose, true, lf).Close()


	http.HandleFunc("/httpcmd", httpCmd)
	http.HandleFunc("/httpcmdList", httpCmdList)
	http.HandleFunc("/httpcmdReload", httpCmdReload)
	http.HandleFunc("/httpcmdInfo", httpCmdInfo)
	
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

	logger.info("httpCmdRun =>",item)
	responseBody := "Invalid request"
	//通过简单的md5校验实现简单层面的安全控制
	timeStr:=time.Now().Format("2006010215")
	if key == md5V(akey+":"+timeStr+":"+item) {
		if realCmd, ok := cmdlist[item]; ok {
			command := realCmd.(string)
			responseBody = RunCmd(command, shell)
		}
	}
	fmt.Fprintln(w, responseBody)
}

func httpCmdReload(w http.ResponseWriter, r *http.Request) {
	loadConfig()
	fmt.Fprintln(w, "success")
}

func httpCmdList(w http.ResponseWriter, r *http.Request) {
	list := Map_Get_Keys(cmdlist)
	jsonStr, err := json.Marshal(list)
	if err != nil {
			fmt.Println("ToJsonDemo err: ", err)
	}
	fmt.Fprintln(w, string(jsonStr))
}

func httpCmdInfo(w http.ResponseWriter, r *http.Request){
	responseBody := ""
	for k, v := range cmdlist {
		responseBody += k+v.(string)
	}
	fmt.Fprintln(w, responseBody)
}
