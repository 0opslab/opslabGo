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

	// fmt.Fprintln(w, "你发送的请求地址是：", r.URL.Path)
	// fmt.Fprintln(w, "你发送的请求地址后查询字符串是：", r.URL.RawQuery)
	// // 获取请求头里面所有信息
	// fmt.Fprintln(w, "请求头中所有信息：", r.Header)
	// // 请求头中xxx的信息
	// fmt.Fprintln(w, "请求头中Accept-Encoding信息：", r.Header["Accept-Encoding"])
	// fmt.Fprintln(w, "请求头中Accept-Encoding属性值是：", r.Header.Get("Accept-Encoding"))
	// // get没有请求体，所以要用post，获取请求体内容长度
	// // len := r.ContentLength
	// // 创建byte切片
	// // body := make([]byte, len)
	// // 将请求体中内容读到body中
	// // r.Body.Read(body)
	// // 在浏览器中显示请求体内容
	// // fmt.Fprintln(w, "请求体中内容：", string(body))

	// // 解析表单，在调用r.form之前必须执行该操作
	// // r.ParseForm()
	// // 获取请求参数
	// // 如果form表单的action属性地址中也有与form表单参数名相同的请求参数，那么参数值都可以得到
	// // 并且form表单中的参数值在URL参数值前面
	// // fmt.Fprintln(w, "请求参数有：", r.Form)
	// // fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)
	// // 通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值
	// fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	// fmt.Fprintln(w, "form表单中的username请求参数的值是：", r.PostFormValue("username"))

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
