package main

import (
	"fmt"
	"math"
)

/*
题目：给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLIS(nums []int) int {

	dp := []int{}
	var maxNum int
	var res int

	for i := 0; i < len(nums); i++ {
		maxNum = 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > maxNum {
				maxNum = dp[j]
			}
		}
		dp = append(dp, maxNum+1)
		res = int(math.Max(float64(res), float64(maxNum+1)))
	}

	return res
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

// 总结：
