package main

import "fmt"

//@Document 演示fmt包中print系列函数的使用

func main(){
	fmt.Print("a","b",1,2,"c","d")
	//在字符串末尾添加换行符
	//add a newline char '\n' at the end of the string
	fmt.Println("a","b",1,2,"c","d")

	//format string
	fmt.Printf("ab %d %d %d cd\n",1,2,3)


	s := fmt.Sprint("a","b",1,2,"c","d")
	fmt.Println(s)



}