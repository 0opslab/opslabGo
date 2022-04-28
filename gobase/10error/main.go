package main

import (
	"errors"
	"fmt"
)

//Go 语言使用一个独立的·明确的返回值来传递错误信息的。这与使用异常的
//Java 和 Ruby 以及在 C 语言中经常见到的超重的单返回值/错误值相比，
//Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用

func main() {
	for _, value := range []int{7, 42} {
		if r, e := f1(value); e != nil {
			fmt.Println("f1 exec failed", e)
		} else {
			fmt.Println("f1 exec success", r)
		}
	}

	for _, value := range []int{7, 42} {
		if r, e := f2(value); e != nil {
			fmt.Println("f1 exec failed", e)
		} else {
			fmt.Println("f1 exec success", r)
		}
	}
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	//返回值nil代表没有错误
	return arg + 3, nil
}

//通过统一错误信息
type argError struct {
	arg  int
	prob string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with"}
	}
	return arg + 3, nil
}
