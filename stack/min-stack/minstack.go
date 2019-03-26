package stack

import (
	arrarystack "github.com/rh01/advance-go/stack/arrary-stack"
)

// 实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作
type Stack struct {
	stackMin  arrarystack.Stack
	stackData arrarystack.Stack
}

func New(initSize int) *Stack {
	stackMin := &arrarystack.Stack{Elements: make([]int, initSize), Size: 0}
	stackData := &arrarystack.Stack{Elements: make([]int, initSize), Size: 0}

	return &Stack{
		stackMin:  *stackMin,
		stackData: *stackData,
	}
}

func (s *Stack) Push(data int) {
	// var MIN int = math.MaxInt64

	if s.stackMin.IsEmpty() {
		s.stackMin.Push(data)
	} else if data < s.GetMin() {
		s.stackMin.Push(data)
		// MIN = data
	}
	s.stackData.Push(data)
	// return true
}

func (s *Stack) Pop() int {
	if s.stackData.IsEmpty() {
		return -1
	}
	data, _ := s.stackData.Pop()
	if data == s.GetMin() {
		s.stackMin.Pop()
	}
	return data

}

func (s *Stack) GetMin() int {
	if s.stackMin.IsEmpty() {
		return -1
	}
	return s.stackMin.Peek()
}
