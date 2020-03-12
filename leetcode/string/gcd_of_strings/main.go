package main

import (
	"fmt"
)

/*
题目：对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。



示例 1：

输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
示例 2：

输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"
示例 3：

输入：str1 = "LEET", str2 = "CODE"
输出：""


提示：

1 <= str1.length <= 1000
1 <= str2.length <= 1000
str1[i] 和 str2[i] 为大写英文字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func gcdOfStrings(str1 string, str2 string) string {
	gcdNum := gcd(len(str1), len(str2))

	for i, ch := range str1 {
		if ch != rune(str1[i%gcdNum]) {
			return ""
		}
	}

	for i, ch := range str2 {
		if ch != rune(str1[i%gcdNum]) {
			return ""
		}
	}

	return str1[:gcdNum]
}

func gcd(temp1 int, temp2 int) int {
	var gcdnum int
	for i := 1; i <= temp1 && i <= temp2; i++ {
		if temp1%i == 0 && temp2%i == 0 {
			gcdnum = i
		}
	}
	return gcdnum
}

func main() {
	str1 := "ABABAB"
	str2 := "ABAB"

	fmt.Println(gcdOfStrings(str1, str2))
}

// 总结：
