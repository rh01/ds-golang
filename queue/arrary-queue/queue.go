package arraryqueue

type ArraryQueue struct {
	Elements []int
	First    int
	Last     int
} // size 表示当前队列的大小，出队列在数组的最后一个元素

func New(initSize int) *ArraryQueue {
	if initSize < 0 {
		return nil
	}
	return &ArraryQueue{
		Elements: make([]int, initSize),
		First:    0,
		Last:     0, // 队尾
	}
}

// 查看
func (aq *ArraryQueue) Peek() int {
	if aq.First < 0 {
		return -1
	}
	return aq.Elements[aq.First]
}

// 入队
func (aq *ArraryQueue) Enqueue(data int) bool {
	if aq.Last == len(aq.Elements)-1 {
		return false
	}
	aq.Elements[aq.Last] = data
	aq.Last++
	return true
}

// 出队
func (aq *ArraryQueue) Dequeue() int {
	if aq.First == aq.Last && aq.First == len(aq.Elements) {
		return -1
	}
	data := aq.Elements[aq.First]
	aq.First++
	return data
}

func (aq *ArraryQueue) Len() int {
	return aq.Last - aq.First
}


func (aq *ArraryQueue) IsEmpty() bool {
	return aq.Last - aq.First == 0
}