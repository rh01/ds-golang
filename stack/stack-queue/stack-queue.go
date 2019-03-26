package stackqueue

import (
	"github.com/rh01/advance-go/stack/arrary-stack"
)

// 栈结构实现队列结构
type StackQueue struct {
	stackData arrarystack.Stack
	stackPop  arrarystack.Stack
}

func New(initSize int) *StackQueue {
	stackData := &arrarystack.Stack{Elements: make([]int, initSize), Size: 0}
	stackPop := &arrarystack.Stack{Elements: make([]int, initSize), Size: 0}

	return &StackQueue{
		stackPop:  *stackPop,
		stackData: *stackData,
	}
}

func (sq *StackQueue) Push(data int) {
	sq.stackData.Push(data)
}

func (sq *StackQueue) Poll() int {
	if sq.stackData.IsEmpty() && sq.stackPop.IsEmpty() {
		return -1
	} else if sq.stackPop.IsEmpty() {
		for !sq.stackData.IsEmpty() {
			data, _ := sq.stackData.Pop()
			sq.stackPop.Push(data)
		}
	}
	data, _ := sq.stackPop.Pop()
	return data
}

func (sq *StackQueue) Peek() int {
	if sq.stackData.IsEmpty() && sq.stackPop.IsEmpty() {
		return -1
	} else if sq.stackPop.IsEmpty() {
		for !sq.stackData.IsEmpty() {
			data, _ := sq.stackData.Pop()
			sq.stackPop.Push(data)
		}
	}
	return sq.stackPop.Peek()
}
