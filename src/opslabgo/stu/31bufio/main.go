package main

import (
	"strings"
	"bufio"
	"fmt"
	"encoding/hex"
)

func main() {
	str := "this is test string"
	s := strings.NewReader(str)
	Str2Hex(str)

	buf_reader := bufio.NewReader(s)

	fmt.Println(buf_reader.Buffered())

	b, _ := buf_reader.Peek(0)
	fmt.Println("读取数据Peek:", b, " ->", string(b))

	c, _ := buf_reader.Peek(5)
	fmt.Println("读取数据Peek:", c, " ->", string(c))

	d := make([]byte, 5)
	n, err := buf_reader.Read(d)
	if err != nil {
		fmt.Println("读取错误 ...", err)
	}
	fmt.Println("读取长度:", n, " 读取结果:", string(b[:n]))

	e := make([]byte, 5)
	f, err := buf_reader.Read(e)
	if err != nil {
		fmt.Println("读取错误 ...", err)
	}
	fmt.Println("读取长度:", n, " 读取结果:", string(e[:f]))

	src := []byte("Hello GO中国")
	encodedStr := hex.EncodeToString(src)
	fmt.Println(encodedStr)

	test, _ := hex.DecodeString(encodedStr)
	fmt.Println(string(test))

	br := bufio.NewReader(strings.NewReader("ABC DEF GHI JKL"))

	for {
		w, err := br.ReadBytes(' ')
		if err != nil {
			return
		}
		fmt.Printf("%q\n", w)
	}

}

func Str2Hex(str string) {
	var sa = make([]string, 0)
	for _, v := range str {
		sa = append(sa, fmt.Sprintf("%0X", v))
	}
	fmt.Println(str + " ===> " + strings.Join(sa, " "))
}