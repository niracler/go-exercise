package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int) // 保存一个字母上一次出现的下标的哈希表
	start := 0
	maxLength := 0

	// 遍历每一个字母
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch] //拿出上一个下标
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	res := lengthOfLongestSubstring("616161")
	fmt.Println(res)
}

/*
分析：对于每一个字母x

- lastOccurred[x] 不存在， 或者 < start -> 无需操作
- lastOccurred[x] >= start -> 更新 start
- 更新 lastOccurred[x] ， 更新 maxLength

*/
