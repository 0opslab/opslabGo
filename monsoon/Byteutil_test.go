package monsoon

import (
	"fmt"
	"testing"
)

func TestMd5Byte(t *testing.T) {
	data := []byte("TestString测试")
	fmt.Println(Md5Byte(data))
}