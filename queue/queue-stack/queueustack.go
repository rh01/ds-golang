package queuestack

import arraryqueue "github.com/rh01/advance-go/queue/arrary-queue"

// TODO: wrong implements.
type QueueStack struct {
	Queue arraryqueue.ArraryQueue
	Help  arraryqueue.ArraryQueue
}

func New(initSize int) *QueueStack {
	q := arraryqueue.ArraryQueue{
		Elements: make([]int, initSize),
		First:    0,
		Last:     0,
	}
	h := arraryqueue.ArraryQueue{
		Elements: make([]int, initSize),
		First:    0,
		Last:     0,
	}
	return &QueueStack{q,h}
}

func (qs *QueueStack) Push(data int) bool {
	return qs.Queue.Enqueue(data)
}

func (qs *QueueStack) Peek() int {
	if qs.Queue.IsEmpty() {
		return -1
	}
	for  qs.Queue.Len() != 1 {
		qs.Help.Enqueue(qs.Queue.Dequeue())
	}
	data := qs.Queue.Dequeue()
	qs.Help.Enqueue(data)
	swap(qs.Queue,qs.Help)
	return data
}


func (qs *QueueStack) Pop() int {
	if qs.Queue.IsEmpty() {
		return -1
	}
	for  qs.Queue.Len() > 1 {
		qs.Help.Enqueue(qs.Queue.Dequeue())
	}
	data := qs.Queue.Dequeue()
	swap(qs.Queue,qs.Help)
	return data
}

func swap(q1, q2 arraryqueue.ArraryQueue) {
	q1, q2 = q2, q1
}
