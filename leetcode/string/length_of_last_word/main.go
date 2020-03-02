package main

import (
	"fmt"
	"strings"
)

/*
题目：Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word (last word means the last appearing word if we loop from left to right) in the string.

If the last word does not exist, return 0.

Note: A word is defined as a maximal substring consisting of non-space characters only.

Example:

Input: "Hello World"
Output: 5
url:https://leetcode.com/problems/length-of-last-word/
*/

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	sl := strings.Split(s, " ")
	return len(sl[len(sl)-1])
}

func main() {
	s := "hello world "
	fmt.Println(lengthOfLastWord(s))
}

// 总结： 关键是要消除后面的空格
