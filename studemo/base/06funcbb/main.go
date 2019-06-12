package main

import (
	"fmt"
)

//定义一个累加器函数,由于返回的是一个函数所有函数相关的变量n，因此n也整体保留因此参数闭包
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()

	fmt.Println(f(1))
	fmt.Println(f(1))
}
