package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"runtime/debug"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", debug.Stack())
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r)  //这些信息是输出到服务器端的打印信息

	fmt.Println("path", r.URL)
	fmt.Println("path", r.URL.Path)

	fmt.Println(r.Form["id"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}