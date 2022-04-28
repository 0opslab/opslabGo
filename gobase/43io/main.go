package main

import(
	"io"
	"fmt"
	"os"
	"strings"
)

//@Document 演示io包中的函数
func main(){

	//os.Stdout实现了writer接口
	//WriteString将字符串s写入到w中
	io.WriteString(os.Stdout,"Hello world!")

	//ReadAtLest从r中读取数据到buf中，要求至少读取min个字节


	r := strings.NewReader("hello world")
	b := make([]byte,32)

	n,err := io.ReadAtLeast(r,b,10)
	fmt.Println("%s\n %d,%v",b,n,err)


	//readFull的功能和ReadAtLest一样，只不过min=len(buf)
	//其中要求最少读取的字节数目是len(buf)，当r中数据少于len(buf)时便会报错
	n1,err := io.ReadFull(r,b)
	fmt.Println("%s\n %d,%v",b,n1,err)

	// CopyN 从 src 中复制 n 个字节的数据到 dst 中
	// 它返回复制的字节数 written 和复制过程中遇到的任何错误
	// 只有当 written = n 时，err 才返回 nil
	// 如果 dst 实现了 ReadFrom 方法，则调用 ReadFrom 来执行复制操作
    n2, err := io.CopyN(os.Stdout, r, 20)
    fmt.Printf("\n%d, %v", n2, err)

    // Copy 从 src 中复制数据到 dst 中，直到所有数据复制完毕
	r3 := strings.NewReader("hello world")
    n3, err := io.Copy(os.Stdout, r3)
    fmt.Printf("\n%d, %v", n3, err)
}