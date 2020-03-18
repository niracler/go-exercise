package main

import (
	"fmt"
	"math"
)

/*
题目：Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
url:https://leetcode.com/problems/maximum-product-subarray/
*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := [][]int{
		[]int{nums[0], nums[0]},
	}

	for i := 1; i < len(nums); i++ {
		maxNum := int(math.Max(float64(nums[i]*dp[i-1][0]), float64(nums[i]*dp[i-1][1])))
		minNum := int(math.Min(float64(nums[i]*dp[i-1][0]), float64(nums[i]*dp[i-1][1])))
		dp = append(dp, []int{
			int(math.Max(float64(nums[i]), float64(maxNum))),
			int(math.Min(float64(nums[i]), float64(minNum))),
		})
	}

	res := dp[0][0]

	for i := 1; i < len(dp); i++ {
		if res < dp[i][0] {
			res = dp[i][0]
		}
	}

	return res
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

// 总结：
