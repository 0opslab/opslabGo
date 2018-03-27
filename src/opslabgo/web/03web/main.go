package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"os"
	"io"
	"bufio"
	//"encoding/json"
	"math/rand"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r)  //这些信息是输出到服务器端的打印信息
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	random := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Fprintf(w, prose_list[random.Intn(len(prose_list))])

}

var prose_list []string

func main() {
	prose_list = make([]string, 0)
	fi, err := os.Open("c:/data.json")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		var line = string(a)
		str := strings.Replace(line, "'", "\"", -1)
		str = strings.Replace(line, "\n", "", -1)
		prose_list = append(prose_list, str)
	}

	fmt.Println("init prose list. len", len(prose_list))
	for _, item := range prose_list {
		fmt.Println(item)
	}

	http.HandleFunc("/", sayhelloName) //设置访问的路由

	//设置监听的端口
	error := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", error)
	}
}