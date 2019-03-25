package heapsort

func HeapSort(arrary []int) {

	if arrary == nil || len(arrary) < 2 {
		return
	}

	size := len(arrary)
	// 1. 插入堆
	for i := 0; i < size; i++ {
		heapInsert(arrary, i)
	}

	// 2. 调整堆(大顶堆)
	size--
	swap(arrary, 0, size)
	for size > 0 {
		// 调整堆的主逻辑
		heapify(arrary, 0, size)
		size--
		swap(arrary, 0, size)
	}

}

// insert element to arrary according to relation of arrary index of heap structure 
func heapInsert(arrary []int, index int) {
	for arrary[index] > arrary[(int(index-1)/2)] {
		swap(arrary, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// swap two arrary elements given index
func swap(arrary []int, a, b int) {
	arrary[a], arrary[b] = arrary[b], arrary[a]
}

func heapify(arrary []int, index, size int) {
	left := index*2 + 1 // left child node
	var largest = 0
	for left < size {
		// 确定最大的index，相对index来说
		if left+1 < size && arrary[left+1] > arrary[left] {
			largest = left + 1
		} else {
			largest = left
		}

		if arrary[largest] <= arrary[index] {
			break
		}

		swap(arrary, index, largest)
		index = largest
		left = index*2 + 1
	}

}
