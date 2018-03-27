package main

import (
	"net/http"
	"log"
	"fmt"
	"opslabgo/gosearch/sites"
	"encoding/json"
)
func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("k")
	count := r.Form.Get("count")
	fmt.Println(key,count)
	Lst,err := sites.Baidu_search(key,10);
	if err == nil {
		res1B, _ := json.Marshal(Lst)
		fmt.Println(string(res1B))
	}
}


func main(){
	http.HandleFunc("/search",search)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("Service:",err)
	}
}
