package stack

import "testing"

// 实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作
func TestArraryStack(t *testing.T) {
	//Test
	aStack := New(10)
	aStack.Push(1)
	aStack.Push(2)
	aStack.Push(3)
	aStack.Push(4)
	aStack.Push(5)
	// if aStack.Len() != 5 {
	// 	t.Fail()
	// }
	if aStack.GetMin() != 1 {
		t.Fail()
	}

	// Test
	// aStack.Pop()
	// aStack.Pop()
	// if aStack.Len() != 3 {
	// 	t.Fail()
	// }

	// Test
	// if aStack.Peek() != 3 {
	// 	t.Fail()
	// }

	// Test

}
