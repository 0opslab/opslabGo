package main

import (
	"fmt"
	"sort"
)

func main() {
	//go的sort包实现了内置和用户自定义数据类型的排序功能
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 3}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)

	//自定义排序
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
