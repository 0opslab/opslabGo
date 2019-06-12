package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//通过&i获取i的内存地址
	fmt.Println("pointer:", &i)
}

//定义一个普通的函数
func zeroval(ival int) {
	ival = 0
}

//定义一个函数，起参数为一个指向int的指针
func zeroptr(iptr *int) {
	*iptr = 0
}
