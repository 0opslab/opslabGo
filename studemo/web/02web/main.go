package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"os"
	"io"
	"bufio"
	"runtime/debug"
	//"encoding/json"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r)  //这些信息是输出到服务器端的打印信息
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	f, err := os.Open("c:/data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		fmt.Fprintf(w, strings.Replace(line, "\n", "", -1))
		/*
		str := strings.Replace(line, "'", "\"", -1)
		str = strings.Replace(line, "\n", "", -1)
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(str), &dat); err == nil {
			w.Header().Set("Content-Type","application/json;charset=utf-8")

			if err := json.NewEncoder(w).Encode(dat); err != nil {
				fmt.Println(err,"=>",str)
				panic(err)
			}
		} else {
			fmt.Println(err,"=>",line)
		} 
		*/
	}

}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}