package monsoon

import (
	"bytes"
	"encoding/hex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

const num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o",
	25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z"}

func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// transform UTF-8 bytes to GBK bytes
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// transform GBK string to UTF-8 string and replace it, if transformed success, returned nil error, or died by error message
func StrToUtf8(str string) (string, error) {
	b, err := GbkToUtf8([]byte(str))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// transform UTF-8 string to GBK string and replace it, if transformed success, returned nil error, or died by error message
func StrToGBK(str string) (string, error) {
	b, err := Utf8ToGbk([]byte(str))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// 字符串转16进制
func StrHex(str string) (string) {
	return hex.EncodeToString([]byte(str))
}

// 十六进制转字符串
func HexStr(hexStr string) (string) {
	b, err := hex.DecodeString(hexStr)
	if (err != nil) {
		return ""
	}
	return string(b)
}
