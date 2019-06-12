package main

import "fmt"

func main() {
	//简单的演示
	fmt.Println("Hello world")

	fmt.Println("go" + " style")
	fmt.Println("1+1=", 1 + 1)
	fmt.Println("7.0/3.0=", 7.0 / 3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)


	//变量定义
	//var 声明一个或多个变量
	var a string = "initstr"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	//推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	//以默认值出事一个变量
	var e int
	fmt.Println(e)

	//:= 语句是申明兵初始化变量的简写
	//相当于var f string ="short"
	f := "short"
	fmt.Println(f)

	//const 用于声明一个常量
	const s string = "constant str"


	//for循环
	//最常用的，带单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//初始化，条件，后续
	for j := 7; j < 10; j++ {
		fmt.Println(j)
	}

	//不带条件的for循环一直执行，知道内部使用break;或者return
	for {
		fmt.Println("loop")
		break
	}



	//if-else
	if 7 % 2 == 0 {
		fmt.Println("7能被2整除")
	} else {
		fmt.Println("7不能被2整除")
	}

	//一个稍微复杂的if-elss if- else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	ii := 2
	fmt.Print("write ", ii, " as ")
	switch ii {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	iii := 2
	fmt.Print("write ", iii, " as ")
	switch iii {
	case 1, 2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	}

}