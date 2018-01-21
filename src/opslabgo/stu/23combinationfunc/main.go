package main

import "strings"
import "fmt"

//go中不支持不支持泛型，
//当你的程序或者数据类型需要时，
//通常是通过组合的方式来提供操作函数。

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(index(strs, "pear"))
	fmt.Println(include(strs, "grape"))
	fmt.Println(any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(all(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(mapItem(strs, strings.ToUpper))
}

//返回目标字符串t出现的第一个索引的位置,没有返回-1
func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//在目标字符串t在切片中返回true
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//如果这些切片中的字符串有一个满足条件 f 则返回true。
func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//如果切片中的所有字符串都满足条件 f 则返回 true。
func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//返回一个包含所有切片中满足条件 f 的字符串的新切片。
func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//返回一个对原始切片中所有字符串执行函数 f 后的新切片。
func mapItem(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
