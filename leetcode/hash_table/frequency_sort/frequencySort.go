package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
题目：
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
Example 2:

Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
Example 3:

Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.

url:https://leetcode.com/problems/sort-characters-by-frequency/
*/

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func frequencySort(s string) string {

	// 循环，计数
	countMap := make(map[rune]int)
	for _, v := range s {
		countMap[v]++
	}

	// 排序
	pl := make(PairList, len(countMap))
	i := 0
	for k, v := range countMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	result := ""
	for i, v := range pl {
		fmt.Println(i, v)
		result += strings.Repeat(string(v.Key), v.Value)
	}

	return result
}

func main() {
	fmt.Println(frequencySort("Aa   bb"))
}

// 总结：一开始的大致想法是， 先计数后排序
// If there are spaces in the test case, an error will be reported
