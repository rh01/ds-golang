package arraryqueue

import "testing"

func TestArraryQueue(t *testing.T) {
	aq := New(10)
	aq.enqueue(1)
	aq.enqueue(2)
	aq.enqueue(3)
	aq.enqueue(4)
	aq.enqueue(5)
	if aq.len() != 5 {
		t.Fail()
	}

	if aq.peek() != 1 {
		t.Fail()
	}

	aq.dequeue()
	aq.dequeue()
	aq.dequeue()
	if aq.len() != 2 {
		t.Fail()
	}

	if aq.peek() != 4 {
		t.Fail()
	}

}
