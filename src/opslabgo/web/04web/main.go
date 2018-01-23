package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
	"io"
	"bufio"
	"log"
	"math/rand"
	"encoding/json"
	"time"
)

func load_file(filename string, list *[]string) {
	fi, err := os.Open(filename)
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
		*list = append(*list, str)
	}

	fmt.Println("init list. len", len(*list))
}

func apeomAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	random := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Fprintf(w, prose_list[random.Intn(len(prose_list))])
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var request = make(map[string]string)
	ip := r.Header.Get("Remote_addr")
	if (ip == "") {
		ip = r.RemoteAddr
	}
	request["remoteAdd"] = ip
	request["requestURI"] = r.RequestURI
	request["method"] = r.Method
	form, _ := json.Marshal(r.Form)
	request["form"] = string(form)
	header, _ := json.Marshal(r.Header)
	request["header"] = string(header)

	b, _ := json.Marshal(request)
	fmt.Fprintf(w, string(b))
}

var prose_list []string

func main() {
	prose_list = make([]string, 0)
	load_file("./data/data.json", &prose_list)

	//设置访问的路由
	http.HandleFunc("/", index)
	http.HandleFunc("/apeom.php", apeomAction)

	//设置监听的端口
	error := http.ListenAndServe(":9090", nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}
}