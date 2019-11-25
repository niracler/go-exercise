package main

import "fmt"

/*
题目：Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]

url:https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type point_level struct {
	pnt *TreeNode
	lvl int
}

func levelOrderBottom(root *TreeNode) [][]int {

	var queue []point_level
	var current point_level
	ongoing_level := -1
	values := [][]int{}

	queue = append(queue, point_level{pnt: root, lvl: 0})

	for 0 != len(queue) {
		current, queue = queue[0], queue[1:]
		pnt := current.pnt
		if nil != current.pnt {
			lft := pnt.Left
			rgt := pnt.Right
			val := pnt.Val
			lvl := current.lvl + 1
			if ongoing_level < current.lvl {
				ongoing_level = current.lvl
				values = append([][]int{{}}, values...)
			}
			values[0] = append(values[0], val)
			queue = append(queue, point_level{pnt: lft, lvl: lvl}, point_level{pnt: rgt, lvl: lvl})
		}
	}

	return values
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(levelOrderBottom(root))
}

// 总结：关键是，你怎么知道它是哪一层呢？？？
