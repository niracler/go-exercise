package main

import (
	"fmt"
	"math"
)

/*
题目：Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
url:https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	dp := prices[0]
	for i := 1; i < len(prices); i++ {
		res = int(math.Max(float64(prices[i]-dp), float64(res)))
		dp = int(math.Min(float64(prices[i]), float64(dp)))
	}

	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// 总结：
