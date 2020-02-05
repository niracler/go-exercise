package main

import (
	"fmt"
	"math"
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
	s  []int // 普通的栈
	ms []int // 栈顶保证是当前栈内最小值的栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)

	var minNum float64
	if len(this.ms) == 0 {
		minNum = float64(x)
	} else {
		minNum = math.Min(float64(this.ms[len(this.ms)-1]), float64(x))
	}

	this.ms = append(this.ms, int(minNum))
}

func (this *MinStack) Pop() {
	if len(this.ms) != 0 {
		this.ms = this.ms[:len(this.ms)-1]
		this.s = this.s[:len(this.s)-1]
	}
}

func (this *MinStack) Top() int {
	var res int
	if len(this.s) != 0 {
		res = this.s[len(this.s)-1]
	}
	return res
}

func (this *MinStack) GetMin() int {
	var res int
	if len(this.ms) != 0 {
		res = this.ms[len(this.ms)-1]
	}
	return res
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
	obj.Push(4)
	obj.Push(8)
	fmt.Println(obj.ms)
	fmt.Println(obj.s)
	obj.Pop()
	fmt.Println(obj.ms)
	fmt.Println(obj.s)
	param3 := obj.Top()
	param4 := obj.GetMin()

	fmt.Println(param3)
	fmt.Println(param4)
}

// 总结： 这道题的关键是维护这两个内部的栈
// s  []int // 普通的栈
// ms []int // 栈顶保证是当前栈内最小值的栈
