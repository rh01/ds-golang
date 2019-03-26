package arrarystack

import "testing"

func TestArraryStack(t *testing.T) {
	//Test
	aStack := New(10)
	aStack.push(1)
	aStack.push(2)
	aStack.push(3)
	aStack.push(4)
	aStack.push(5)
	if aStack.len() != 5 {
		t.Fail()
	}

	// Test
	aStack.pop()
	aStack.pop()
	if aStack.len() != 3 {
		t.Fail()
	}

	// Test
	if aStack.peek() != 3 {
		t.Fail()
	}
}
