package main

import "fmt"

func main() {

	// 数组的例子
	// 数组就是内建容器
	// 数组跟切片不是同一个东西
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 1, 2, 3, 4, 56, 78, 9} // 这个才是让他自己数的写法，假如没有三个点， 就是切片
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// slice
	fmt.Println(arr3[:3])

	// 我们需要一个指定长度的slice
	s1 := make([]int, 0, 11)
	for i := 0; i < 50; i++ {
		s1 = append(s1, i)
		fmt.Printf("cap=%d, len=%d\n", cap(s1), len(s1))
	}
	fmt.Println(s1)

	// 拷贝, s2要先有相应的len才可以
	s2 := make([]int, 10, 500)
	copy(s2, s1[10:30])
	fmt.Println(s2)

	// slice之间的append
	s3 := append(s2[:3], s2[4:]...)
	fmt.Println(s3)
}
