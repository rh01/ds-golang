package arrarystack

import "testing"

func TestArraryStack(t *testing.T) {
	//Test
	aStack := New(10)
	aStack.Push(1)
	aStack.Push(2)
	aStack.Push(3)
	aStack.Push(4)
	aStack.Push(5)
	if aStack.Len() != 5 {
		t.Fail()
	}

	// Test
	aStack.Pop()
	aStack.Pop()
	if aStack.Len() != 3 {
		t.Fail()
	}

	// Test
	if aStack.Peek() != 3 {
		t.Fail()
	}
}
