package main

import "math"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.Min()
	minStack.Pop()
	minStack.Top()
	minStack.Min()
}

type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor /** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	t := s.minStack[len(s.minStack)-1]
	s.minStack = append(s.minStack, min(t, x))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Min() int {
	return s.minStack[len(s.minStack)-1]
}

func min(l, r int) int {
	if l > r {
		return r
	}
	return l
}
