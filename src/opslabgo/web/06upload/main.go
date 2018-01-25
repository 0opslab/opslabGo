package main
import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strings"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"path/filepath"
)
//要使表达能够上传文件，首先第一不久是要添加form的enctype属性,enctype属性有如下山中情况
//application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
//multipart/form-data      不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
//text/plain      空格转换为 "+" 加号，但不对特殊字符编码。
func http_info(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	params := ""
	for k, v := range r.Form {
		params += "&"+k+"="+strings.Join(v, "")
	}
	fmt.Fprintf(w, params) //这个写入到w的是输出到客户端的
}
/**
 处理登录信息
 */
func upload(w http.ResponseWriter,r *http.Request){
	//获取请求方式
	fmt.Println("methdo:",r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("/local/workspace/opslabGo/data/web/upload.gtpl")
		t.Execute(w, token)
	} else {
		//设置maxMemory
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		file_name := filepath.Base(handler.Filename);
		f, err := os.OpenFile("/tmp/"+file_name, os.O_WRONLY|os.O_CREATE, 0666)  // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main(){
	http.HandleFunc("/",http_info)
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("Service:",err)
	}
}