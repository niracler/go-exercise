package main

import (
	"fmt"
	"math"
)

/*
题目：Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
url:https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/780/
*/

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	var start int
	var end int

	for i := 0; i < len(s); i++ {
		// 以I为中心向两遍散开
		s1 := getPalin(s, i, i)
		s2 := getPalin(s, i, i+1)
		tmp := int(math.Max(float64(s1), float64(s2)))
		if tmp > end-start {
			start = i - (tmp-1)/2
			end = i + tmp/2
		}
	}

	return s[start : end+1]
}

func getPalin(s string, l int, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}

// 总结：以I为中心向两遍散开
