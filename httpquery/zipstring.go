package main

import (
	"bytes"
	"compress/flate"
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

func FlateEncodeByte(input []byte) (result []byte, err error) {
	var buf bytes.Buffer
	w, err := flate.NewWriter(&buf, -1)
	w.Write(input)
	w.Close()
	result = buf.Bytes()
	return
}

func FlateDecode(input []byte) (result []byte, err error) {
	return ioutil.ReadAll(flate.NewReader(bytes.NewReader(input)))
}
