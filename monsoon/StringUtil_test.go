package monsoon

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	strs := "this is测试字符串";
	base64 := Base64Encode(strs)
	strs1 := Base64Decode(base64)

	fmt.Println(base64)
	fmt.Println(strs1)
}