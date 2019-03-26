package arraryqueue

import "testing"

func TestArraryQueue(t *testing.T) {
	aq := New(10)
	aq.Enqueue(1)
	aq.Enqueue(2)
	aq.Enqueue(3)
	aq.Enqueue(4)
	aq.Enqueue(5)
	if aq.Len() != 5 {
		t.Fail()
	}

	if aq.Peek() != 1 {
		t.Fail()
	}

	aq.Dequeue()
	aq.Dequeue()
	aq.Dequeue()
	if aq.Len() != 2 {
		t.Fail()
	}

	if aq.Peek() != 4 {
		t.Fail()
	}

}
