package main

import (
	"fmt"
)

/*
题目：

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

url:https://leetcode.com/problems/min-stack/
*/

type MinStack struct {
	MyList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {
	return 0
}

func (this *MinStack) GetMin() int {
	return 1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Pop()
	param3 := obj.Top()
	param4 := obj.GetMin()

	fmt.Println(param3)
	fmt.Println(param4)
}

// 总结：
