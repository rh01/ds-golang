package queuestack

import "testing"

func TestQueueStack(t *testing.T) {
	qs := New(10)
	qs.Push(1)
	qs.Push(2)
	qs.Push(3)
	qs.Push(4)
	qs.Push(5)

	// test
	// if qs.Peek() != 1 {
	// 	t.Fail()
	// }

	// test
	if qs.Pop() != 5 {
		t.Fail()
	}

	if qs.Pop() != -1 {
		t.Fail()
	}

}
