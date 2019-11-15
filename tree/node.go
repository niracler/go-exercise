package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) printNode() {
	fmt.Print(node.value, " ")
}

// 不加星号就是值传递了
func (node *treeNode) setValue(value int) {
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

//中序遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.printNode()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.right.left = createNode(646)
	root.right.left.setValue(555)
	root.right.left.printNode()
	fmt.Println()
	root.traverse()
}
