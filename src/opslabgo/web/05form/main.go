package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strings"
)

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
func login(w http.ResponseWriter,r *http.Request){
	//获取请求方式
	fmt.Println("methdo:",r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/local/workspace/opslabGo/data/web/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "username:"+r.Form["username"][0]+" password:"+r.Form["password"][0])
	}
}

func main(){
	http.HandleFunc("/",http_info)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("Service:",err)
	}
}