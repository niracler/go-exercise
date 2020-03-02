package main

import (
	"fmt"
	"strconv"
)

/*
题目：Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3

url:https://leetcode.com/problems/binary-tree-paths/
*/

//Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	dfs(root, strconv.Itoa(root.Val), &res)

	return res
}

func dfs(root *TreeNode, path string, res *[]string) {

	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}

	if root.Left != nil {
		dfs(root.Left, path+"->"+strconv.Itoa(root.Left.Val), res)
	}

	if root.Right != nil {
		dfs(root.Right, path+"->"+strconv.Itoa(root.Right.Val), res)
	}

}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(binaryTreePaths(root))
}
