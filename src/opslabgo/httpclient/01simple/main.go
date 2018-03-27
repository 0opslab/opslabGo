package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"bytes"
)

func main() {
	//Get、Head、Post、PostForm配合使用实现HTTP请求：

	resp, err := http.Get("http://0opslab.com/")
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("GET status:", resp.StatusCode, " cotent:", body)

	bodyBuf := &bytes.Buffer{}
	resp2, err := http.Post("http://0opslab.com/", "image/jpeg", bodyBuf)
	checkErr(err)
	defer resp2.Body.Close()
	body2, err := ioutil.ReadAll(resp2.Body)
	fmt.Println("GET status:", resp2.StatusCode, " cotent:", body2)

	resp3, err := http.PostForm("http://0opslab.com/", url.Values{"key": {"Value"}, "id": {"123"}})
	checkErr(err)
	defer resp3.Body.Close()
	body3, err := ioutil.ReadAll(resp3.Body)
	fmt.Println("GET status:", resp3.StatusCode, " cotent:", body3)

	req, err := http.NewRequest("GET", "http://0opslab.com/", nil)
	req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	// ...
	client := &http.Client{//
	}
	resp4, err := client.Do(req)
	defer resp4.Body.Close()
	body4, err := ioutil.ReadAll(resp4.Body)
	fmt.Println("GET status:", resp4.StatusCode, " cotent:", body4)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}