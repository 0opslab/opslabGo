package main

import "fmt"

func main() {
	//map是go内置的关联数据类型,也就是所谓的字典

	//创建一个空的map
	m := make(map[string]int)

	//使用管用的map[key]=value语法类设置键值对
	m["k1"] = 7
	m["k2"] = 13

	//输出map的所有键值对
	fmt.Println("map", m)

	//输出指定的键值对
	fmt.Println("map['k1']", m["k1"])

	//输出map的长队
	fmt.Println("map len:", len(m))

	//使用內建的delete可以从map中移除一个键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//申明并创建一个map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//使用key输出map值
	for key := range n {
		fmt.Println("map key=", key, " ,value=", n[key])
	}

	//判的某个key值是否存在在map中
	key_value, ok := n["foo"]
	if ok {
		fmt.Println("map中有key=foo value=", key_value)
	} else {
		fmt.Println("map中没有key=foo的元素")
	}

}
