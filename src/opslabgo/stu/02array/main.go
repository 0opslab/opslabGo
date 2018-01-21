package main

import "fmt"

func main(){
	//在go中数组是一个固定长度的数列

	//创建一个数组a来存放5个int值
	var a [5]int
	fmt.Println("emp:",a)

	//是index来访问个设定数组中值
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])

	//获取数组的长度
	fmt.Println("len:",len(a))

	//创建并赋值指定的值给数组
		b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	//定义多维数组
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i * j
		}
	}
	fmt.Println("2d array:",twoD)

}