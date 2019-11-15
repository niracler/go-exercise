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
}
