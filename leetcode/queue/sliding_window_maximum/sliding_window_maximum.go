package main

import (
	"fmt"
)

/*
题目：
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

Follow up:
Could you solve it in linear time?
url:
*/

func maxSlidingWindow1(nums []int, k int) []int {
	res := []int{}

	if k == 0 || len(nums) < k {
		return res
	}

	for i := 0; i < len(nums)-k+1; i++ {

		// 遍历找出当前窗口的最大值
		maxNum := nums[i]
		for j := 0; j < k; j++ {
			if nums[i+j] > maxNum {
				maxNum = nums[i+j]
			}
		}
		res = append(res, maxNum)
	}

	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	res := []int{}

	if k == 0 || len(nums) < k {
		return res
	}

	//找出第一个最大值
	maxNum := nums[0]
	for i := 0; i < k; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	res = append(res, maxNum)

	// 开始滑动窗口
	for i := k; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]

		} else {
			if maxNum == nums[i-k] {
				maxNum = nums[i]
				for j := i - k + 1; j < i+1; j++ {
					if nums[j] > maxNum {
						maxNum = nums[j]
					}
				}
			}
		}

		res = append(res, maxNum)
	}

	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 || len(nums) < k {
		return nil
	}
	var res []int
	var indexQueue []int // 保存最大值的下标的队列

	// 对数组进行遍历，即滑动窗口
	for i := 0; i < len(nums); i++ {
		//队列中超出窗口的下标值从队列头出队列
		if i >= k && i-indexQueue[0] >= k {
			indexQueue = indexQueue[1:]
		}

		// 比进队列的当前下标的值要小的数都要从队列尾出队列
		for j := len(indexQueue) - 1; j >= 0; j-- {
			if nums[i] > nums[indexQueue[j]] {
				indexQueue = indexQueue[:j]
			}
		}
		indexQueue = append(indexQueue, i)

		// 保存结果
		if i >= k-1 {
			res = append(res, nums[indexQueue[0]])
		}
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow1([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

// 总结：
// 关键是保证maxque[0] 是当前滑动窗口的最大值的下标
// 1. 队列中超出窗口的下标值从队列头出队列
// 2. 比进队列的当前下标的值要小的数都要从队列尾出队列
