package main

import (
	"fmt"
	"math"
)

/*
题目：Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
url:https://leetcode.com/problems/triangle/
*/

func minimumTotal(triangle [][]int) int {
	dp := []int{triangle[0][0]}
	for i := 1; i < len(triangle); i++ {
		tmp := []int{triangle[i][0] + dp[0]}
		for j := 1; j < len(dp); j++ {
			tmp = append(tmp, triangle[i][j]+int(math.Min(float64(dp[j-1]), float64(dp[j]))))
		}
		dp = append(tmp, triangle[i][len(triangle[i])-1]+dp[len(dp)-1])
	}

	minTotal := dp[0]
	for _, num := range dp {
		if minTotal > num {
			minTotal = num
		}
	}

	return minTotal
}

func main() {
	fmt.Println(minimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}))
}

// 总结：
