package arraryqueue

type ArraryQueue struct {
	elements []int
	first    int
	last     int
} // size 表示当前队列的大小，出队列在数组的最后一个元素

func New(initSize int) *ArraryQueue {
	if initSize < 0 {
		return nil
	}
	return &ArraryQueue{
		elements: make([]int, initSize),
		first:    0,
		last:     0, // 队尾
	}
}

// 查看
func (aq *ArraryQueue) peek() int {
	if aq.first < 0 {
		return -1
	}
	return aq.elements[aq.first]
}

// 入队
func (aq *ArraryQueue) enqueue(data int) bool {
	if aq.last == len(aq.elements)-1 {
		return false
	}
	aq.elements[aq.last] = data
	aq.last++
	return true
}

// 出队
func (aq *ArraryQueue) dequeue() (int, bool) {
	if aq.first == aq.last && aq.first == len(aq.elements) {
		return -1, false
	}
	data := aq.elements[aq.first]
	aq.first++
	return data, true
}

func (aq *ArraryQueue) len() int {
	return aq.last - aq.first
}
