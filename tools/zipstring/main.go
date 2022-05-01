package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io/ioutil"
)

/*
golang 用falte zip 压缩、解压字符串
*/

func FlateEncode(input string) (result []byte, err error) {
	var buf bytes.Buffer
	w, err := flate.NewWriter(&buf, -1)
	w.Write([]byte(input))
	w.Close()
	result = buf.Bytes()
	return
}

func FlateDecode(input []byte) (result []byte, err error) {
	result, err = ioutil.ReadAll(flate.NewReader(bytes.NewReader(input)))
	return
}

func main() {
	s := "Hello, 世界Hello, 世界Hello, 世界Hello, 世界Hello, 世界Hello, 世界"

	s1, err := FlateEncode(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s, "原本长度：", len(s), "压缩后的长度:", len(s1))

	// flate
	enflated, err := FlateDecode(s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enflated))

}
