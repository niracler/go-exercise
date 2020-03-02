package tree

import "fmt"

type Node struct {
	Value       interface{}
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 不加星号就是值传递了
func (node *Node) SetValue(value interface{}) {
	node.Value = value
}

func CreateNode(value interface{}) *Node {
	return &Node{Value: value}
}

//中序遍历
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

//在遍历的过程中做点什么
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
