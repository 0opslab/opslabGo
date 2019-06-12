package monsoon

import (
	"fmt"
	"testing"
)

func TestStrToGBK(t *testing.T) {
	str := "hello 华夏10进制转36进制：";
	fmt.Println(StrToGBK(str))
	fmt.Println(str)
	fmt.Println(StrHex(str))
	hexSTT := "68656c6c6f20e58d8ee5a48f3130e8bf9be588b6e8bdac3336e8bf9be588b6efbc9a"
	str = HexStr(hexSTT)
	fmt.Println(str)

}