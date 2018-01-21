package main

import (
	"fmt"
	"math"
)

func main() {
	//演示结构体，方法的，接口的使用

	r := rect{width: 10, height: 20}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

//定义个几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

//定义个接口上的操作函数
func measure(g geometry) {
	fmt.Println("area:", g.area(), " perim:", g.perim())
}

//定义结构体
type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
