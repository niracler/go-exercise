package main

import (
	"fmt"
)

/*
题目：Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]
url:
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	dfs(root, root.Val, sum, []int{root.Val}, &res)

	return res
}

func dfs(root *TreeNode, sum int, target int, path []int, res *[][]int) {

	if root.Left == nil && root.Right == nil {
		if sum == target {
			// fmt.Println(sum)
			// fmt.Println(path)
			*res = append(*res, path)
		}
	}

	if root.Left != nil {
		var tmp = make([]int, len(path)+1)
		copy(tmp, append(path, root.Left.Val))
		dfs(root.Left, sum+root.Left.Val, target, tmp, res)
	}

	if root.Right != nil {
		var tmp = make([]int, len(path)+1)
		copy(tmp, append(path, root.Right.Val))
		dfs(root.Right, sum+root.Right.Val, target, tmp, res)
	}
}

func main() {

	//   	5
	//     / \
	//    4   8
	//   /   / \
	//  11  13  4
	// /  \    / \
	//7    2  5   1

	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}

	fmt.Println(pathSum(root, 22))
}

// 总结：真想省空间的话，我认为不要用数组装起来？？？直接回溯的时候加？
