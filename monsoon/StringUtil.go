package monsoon

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"bytes"
)




// 字符串Base64编码
func Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// 字符串Base64解码
func Base64Decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

// Md5String 获取字符串md5值
func Md5String(s string) string {
	return Md5Byte([]byte(s))
}

// Sha1String 获取字符串sha1值
func Sha1String(s string) string {
	return Sha1Byte([]byte(s))
}

// Sha256String 获取字符串sha256值
func Sha256String(s string) string {
	return Sha256Byte([]byte(s))
}

// Sha512String 获取字符串sha512值
func Sha512String(s string) string {
	return Sha512Byte([]byte(s))
}


func Export(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		return ""
	}
	return buf.String()
}

// json化
func Json(v interface{}) string {
	return Export(v)
}