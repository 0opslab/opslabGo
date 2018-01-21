package main

import (
	"fmt"
)

func main() {
	//range迭代各种各样的数据结构。

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		fmt.Println("inedx", i, " value=", num)
	}

	maps := map[string]string{"a": "a.str", "b": "b.str"}
	for key, value := range maps {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for i, c := range "this is str" {
		fmt.Println(i, ":", c)
	}
}
