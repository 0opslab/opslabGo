package monsoon

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"bytes"
	"regexp"
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

// 字符串反转
func Reverse(s string) string{
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}


// json化
func Json(v interface{}) string {
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

// 检测字符串是否为空
// check string is empty
func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// 判断字符串是否不为空
// check string is not empty
func IsNotEmpty(s string) bool{
	return !IsEmpty(s)
}


// 判断字符串是否是空白字符串
// check string is whitespace,empty
func IsBlank(s string) bool{
	if len(s) == 0{
		return true
	}
	reg := regexp.MustCompile(`^\s+$`)
	actual := reg.MatchString(s)
	if actual {
		return true
	}
	return false
}

// 判断字符串是否为不为空白字符串
func IsNotBlank(s string) bool{
	return !IsBlank(s)
}

