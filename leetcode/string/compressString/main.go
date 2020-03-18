package main

import (
	"fmt"
)

/*
题目：字符串压缩
url:
*/

func compressString(S string) string {
	if S == "" {
		return ""
	}
	res := ""
	count := 0
	for i := 1; i < len(S); i++ {
		count++
		if S[i] != S[i-1] {
			res += string(S[i-1]) + fmt.Sprintf("%d", count)
			count = 0
		}
	}

	res += string(S[len(S)-1]) + fmt.Sprintf("%d", count+1)

	if len(res) >= len(S) {
		return S
	} else {
		return res
	}
}

func main() {
	fmt.Println(compressString("aabcccccaaa"))
}

// 总结：
