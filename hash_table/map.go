package main

import "fmt"

func main() {

	// 直接赋值地创建map
	m1 := map[string]string{
		"name": "niracler",
		"id":   "1235",
		"age":  "132",
	}
	fmt.Println(m1)

	// 使用 make 创建 map m2 == empty map
	m2 := make(map[string]int)
	fmt.Println(m2)

	// 使用 var 定义 m3 == nil
	var m3 map[string]int
	fmt.Println(m3)

	// map 的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 当我们拿一个不存在的 key 的时候， 会获得一个 zero value
	v, ok := m2["dtytfyu"]
	fmt.Println("空值:", v, ok)
	m2["dtytfyu"] = 2345678
	v, ok = m2["dtytfyu"]
	fmt.Println("不是空值:", v, ok)

	// 删除元素
	delete(m1, "age")
	fmt.Println(m1)
}
