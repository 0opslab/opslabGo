package main

import (
	"fmt"
	"runtime"
	"sync"
)

//sync.Pool是一个可以存或取的临时对象集合
//sync.Pool可以安全被多个线程同时使用，保证线程安全
//注意、注意、注意，sync.Pool中保存的任何项都可能随时不做通知的释放掉，所以不适合用于像socket长连接或数据库连接池。
//sync.Pool主要用途是增加临时对象的重用率，减少GC负担。

func main(){
	p := &sync.Pool{
		//New()函数的作用是当从Pool中Get对象时，如果pool为空，则先先通过New创建一个，插入pool中，然后返回对象
		New:func() interface{}{
			return make([]int, 16)
		},
	}

	s := p.Get().([]int)
	s[0]=1
	s[1]=2
	fmt.Printf("%p ===> %v\n",&s,s)
	//将一个对象存入到pool中
	p.Put(s)

	//存pool中取出一个对象
	c := p.Get().([]int)
	c[2] = 3
	fmt.Printf("%p ===> %v\n",&c,c)

	//强制GC
	runtime.GC()
	d := p.Get().([]int)
	fmt.Printf("%p ===> %v\n",&d,d)
}
