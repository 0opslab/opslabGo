package util

import (
	"testing"
	"fmt"
)

func TestGet(t *testing.T) {
	html,_ := GetHtml("https://www.baidu.com/s?ie=utf-8&wd=GO%E4%B8%AD%E5%9B%BD")
	fmt.Println(html)
}
