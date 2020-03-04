package main

import (
	"fmt"
	"sort"
)

/*
题目：
url:
*/

func merge(A []int, m int, B []int, n int) {
	for i := 0; i < n; i++ {
		A[m+i] = B[i]
	}

	sort.Ints(A)
}

func main() {
	A := []int{1, 2, 3, 0, 0, 0}
	B := []int{2, 5, 6}

	merge(A, 3, B, 3)

	fmt.Println(A)
}

// 总结：拼接后排序