package monsoon

/**
	封装一些http相关的一些常用操作
 */


import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// JsQueryEscape escapes the string in javascript standard so it can be safely placed
// inside a URL query.
func JsQueryEscape(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}

// JsQueryUnescape does the inverse transformation of JsQueryEscape, converting
// %AB into the byte 0xAB and '+' into ' ' (space). It returns an error if
// any % is not followed by two hexadecimal digits.
func JsQueryUnescape(s string) (string, error) {
	return url.QueryUnescape(strings.Replace(s, "%20", "+", -1))
}


//Get HTTP RESPONSE BODY DATA
func HttpGet(url string) string{

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(data[:])
}

//DOWNLOAD FILE FORM URL
func HttpDownload(url string,file string) bool {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return true
}