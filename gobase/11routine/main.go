package main

import (
	"fmt"
)

//go协程在执行上来说是轻量级的线程

func main() {
	//这种常规的调用方式，同时运行
	f("direct")

	//使用go来使一个函数到启用go协程去调用
	go f("go routine")

	//匿名的协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("this run")

	//按任意键中断
	var input string
	fmt.Scan(&input)
	fmt.Println("done")
}

//定义一个函数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
