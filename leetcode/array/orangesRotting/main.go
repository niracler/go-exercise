package main

import (
	"fmt"
)

/*
题目：
url:
*/

func orangesRotting(grid [][]int) int {
	day := 0
	good := 0
	m := [][]int{}
	for i1, v1 := range grid {
		for i2, v2 := range v1 {
			if v2 == 2 {
				m = append(m, []int{i1, i2})
			}
			if v2 == 1 {
				good++
			}
		}
	}

	f := true
	for f && good > 0 {
		f = false
		count := len(m)
		for i := 0; i < count; i++ {
			if m[i][0]-1 >= 0 && grid[m[i][0]-1][m[i][1]] == 1 {
				f = true
				grid[m[i][0]-1][m[i][1]] = 2
				good--
				m = append(m, []int{m[i][0] - 1, m[i][1]})
			}
			if m[i][0]+1 < len(grid) && grid[m[i][0]+1][m[i][1]] == 1 {
				f = true
				grid[m[i][0]+1][m[i][1]] = 2
				good--
				m = append(m, []int{m[i][0] + 1, m[i][1]})
			}
			if m[i][1]-1 >= 0 && grid[m[i][0]][m[i][1]-1] == 1 {
				f = true
				grid[m[i][0]][m[i][1]-1] = 2
				good--
				m = append(m, []int{m[i][0], m[i][1] - 1})
			}
			if m[i][1]+1 < len(grid[0]) && grid[m[i][0]][m[i][1]+1] == 1 {
				f = true
				grid[m[i][0]][m[i][1]+1] = 2
				good--
				m = append(m, []int{m[i][0], m[i][1] + 1})
			}
		}

		m = m[count:]
		if f {
			day++
		}
	}

	if good == 0 {
		return day
	} else {
		return -1
	}
}

func main() {
	fmt.Println(orangesRotting([][]int{
		[]int{2, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 0, 1},
	}))
}

// 总结：非常愚蠢的方法
