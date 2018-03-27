package main

import "fmt"

func main() {
	//slice是go中一个关键的数据类型,是一个比数组更加强大的序列接口

	//创建一个长度为3，类型为string的slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	//像数组一样设置值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//使用len获取slice的长度
	fmt.Println("len:", len(s))

}