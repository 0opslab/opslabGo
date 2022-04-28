package main

import (
	"fmt"
)

func main() {
	//函数是go的中心

	res := plus(1, 2)
	fmt.Println("1+2=", res)

	//多返回值的函数
	sum, plan := add(1, 2)
	fmt.Println("sum:", sum)
	fmt.Println("plan:", plan)

	//调用可变参数函数
	sum1, plan1 := sum_func(1, 2, 3, 4, 5)
	fmt.Println("sum:", sum1)
	fmt.Println("plan:", plan1)

	//调用可变参数函数
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum2, plan2 := sum_func(nums...)
	fmt.Println("sum:", sum2)
	fmt.Println("plan:", plan2)

	//调用递归函数
	fmt.Println(fact(7))
}

/**
 * 定义个函数接收俩个int类型的值，返回他们的和
 */
func plus(a int, b int) int {
	return a + b
}

/**
 * 定义返回多个值的函数
 */
func add(a int, b int) (int, int) {
	return a + b, a * b
}

/**
 * 定义个可变参数的函数
 */
func sum_func(nums ...int) (int, int) {
	total := 0
	plan := 1
	for _, num := range nums {
		total += num
		plan = plan * num
	}
	return total, plan
}

/**
 * 定义一个递归函数
 */
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}
