package main

import (
	"fmt"
	"go-exercise/tree"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
}

// 算是重构一下别人的树
type myTreeNode struct {
	node *tree.Node
}

// 自己加的后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.Right.Left = tree.CreateNode(646)
	root.Right.Left.SetValue(555)
	root.Right.Left.Print()
	fmt.Println()
	root.Traverse()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
