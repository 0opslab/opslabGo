package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 构造函数
func result(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

/*
 统一的响应方法
*/
func respone(w http.ResponseWriter, data Response) {
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	raw, _ := json.Marshal(data)
	result := string(raw)
	fmt.Fprint(w, result)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("%s", debug.Stack())
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r) //这些信息是输出到服务器端的打印信息

	fmt.Println("path", r.URL)
	fmt.Println("path", r.URL.Path)

	fmt.Println(r.Form["id"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	res := result(200, "success", nil)
	respone(w, *res)
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
