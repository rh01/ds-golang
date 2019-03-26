package stackqueue

import "testing"

func TestStackQueue(t *testing.T) {
	aStack := New(10)
	aStack.Push(1)
	aStack.Push(2)
	aStack.Push(3)
	aStack.Push(4)
	aStack.Push(5)
	// test
	if aStack.Poll() != 1 {
		t.Fail()
	}
	// test
	if aStack.Peek() != 2 {
		t.Fail()
	}

}
