package main

import (
	"container/list"
	"fmt"
)

func main() {
	ll := list.New()
	//尾部添加原始
	//first为保存list中的句柄
	first := ll.PushBack("first")
	//second 保存句柄
	second := ll.PushBack("second")

	// 在fist之后添加high
	ll.InsertAfter("high", first)

	//首部添加
	ll.PushFront("0000")

	ll.Remove(second)

	for i := ll.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
