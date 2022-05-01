package main

import "fmt"

func main() {
	a, b, c := "", 0, false

	fmt.Sscan("123456 7 true ", &a, &b, &c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	fmt.Sscan("1234567 true ","%d%t", &a, &c)
	fmt.Println(a)
	fmt.Println(c)

	fmt.Sscanf("1234567true ","%4s%d%t", &a,&b,&c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
