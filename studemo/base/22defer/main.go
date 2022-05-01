package main

import (
	"fmt"
	"time"
)

//Golang这么时尚的语言是没有类似try..catch 这种异常处理机制，
//而是使用 panic 和 recover处理异常. 其实相当于python的raise。

//golang的异常处理组合 panic，defer，recover，
//跟java中的try catch finially是类似的。
//但是从语言的用户体验来说，不怎么好。 但考虑到golang的
//场景基本是系统高性能层面的，这种精准错误处理应该减少那种后遗症bug。

//defer 需要放在 panic 之前定义，另外recover只有在 defer 调用的函数中才有效。
//recover处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点.
//多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用
//panic (主动爆出异常) 与 recover （收集异常）
//recover 用来对panic的异常进行捕获. panic 用于向上传递异常，执行顺序是在 defer 之后。

func main() {
	f()
	fmt.Println("end")
}

func f() {
	defer func() {
		//必须要先声明defer，否则不能捕获到panic异常
		fmt.Println(".cc start")
		if err := recover(); err != nil {
			//这里的err其实就是panic传入的内容，"bug"
			fmt.Println(err)
		}
		fmt.Println(".cc end")
	}()
	for {
		fmt.Println("1")
		a := []string{"a", "b"}
		// 越界访问，肯定出现异常
		fmt.Println(a[3])
		// 上面已经出现异常了,所以肯定走不到这里了。
		panic("bug")
		//不会运行的.
		fmt.Println("4")
		time.Sleep(1 * time.Second)
	}
}
