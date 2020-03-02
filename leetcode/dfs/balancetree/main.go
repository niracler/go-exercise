package main

/*
题目：Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
url:https://leetcode.com/problems/balanced-binary-tree/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bool isBalanced(TreeNode *root, int &height)
// {
// 	if(!root) return true;
// 	int left=0,right=0;
// 	if(!isBalanced(root->left,left)) return false;
// 	if(!isBalanced(root->right,right)) return false;
// 	if(abs(left-right) > 1) return false;
// 	height = max(left,right) + 1;
// 	return true;
// }

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	depthDiff := maxDepth(root.Left) - maxDepth(root.Right)
	if depthDiff < 0 {
		depthDiff *= -1
	}
	return depthDiff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}

// 总结：我的天， 这要遍历好多次吧？？？
