package main

import (
	"fmt"
)

func main() {
	//创建一个新的结构体元素
	person1 := person{"bob", 20}
	fmt.Println("person1", person1)

	//明确指定对应关系
	person2 := person{name: "alice", age: 25}
	fmt.Println("person2", person2)

	//省略部分字段初始化
	person3 := person{name: "tomcat"}
	fmt.Println("person3", person3)

	//访问结构体的属性
	fmt.Println("person's name:", person1.name)

	//通过指针来访问
	sp := &person3
	fmt.Println("person3's name:", sp.name)

	//改变结构的值
	sp.name = "tomcat7"
	fmt.Println("person3's name:", sp.name)

	//机构体及方法调用
	r := rect{width: 10, height: 15}
	fmt.Println("r area:", r.area())
	fmt.Println("r perim:", r.perim())

	rp := &r
	fmt.Println("r area:", rp.area())
	fmt.Println("r perim:", rp.perim())
}

//定义一个结构体
type person struct {
	name string
	age  int
}

//定义个结构体，并且附件一些方法
type rect struct {
	width  int
	height int
}

//定义方法
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width * 2 * r.height
}
