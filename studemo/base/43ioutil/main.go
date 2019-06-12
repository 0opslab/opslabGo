package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	s := strings.NewReader("Hello World!")
	ra, _ := ioutil.ReadAll(s)
	fmt.Printf("%s", ra)

	ra1, _ := ioutil.ReadFile("C:\\Windows\\win.ini")
	fmt.Printf("%s", ra1)

	fn := "C:\\Test.txt"
	s1 := []byte("Hello World!")
	ioutil.WriteFile(fn, s1, os.ModeAppend)
	rf, _ := ioutil.ReadFile(fn)
	fmt.Printf("%s", rf)

	rd, err := ioutil.ReadDir("C:\\Windows")
	for _, fi := range rd {
		fmt.Println("")
		fmt.Println(fi.Name())
		fmt.Println(fi.IsDir())
		fmt.Println(fi.Size())
		fmt.Println(fi.ModTime())
		fmt.Println(fi.Mode())
	}
	fmt.Println("")
	fmt.Println(err)


	s2 := strings.NewReader("hello world!")
	//ReadCloser 接口组合了基本的 Read 和 Close 方法。NopCloser 将提供的
	// Reader r 用空操作 Close 方法包装后作为 ReadCloser 返回
	r := ioutil.NopCloser(s2)
	r.Close()
	p := make([]byte, 10)
	r.Read(p)
	fmt.Println(string(p))
}
