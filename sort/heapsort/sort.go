package heapsort

func HeapSort(arrary []int) {

	if arrary == nil || len(arrary) < 2 {
		return
	}

	size := len(arrary)
	// heapArrary := make([]int, size)
	// 1. 插入堆
	for i := 0; i < size; i++ {
		heapInsert(arrary, i)
	}
	// swap(arr, 0, --size);
	size--
	swap(arrary, 0, size)
	//
	for size > 0 {
		heapify(arrary, 0, size)
		size--
		swap(arrary, 0, size)
	}

}

func heapInsert(arrary []int, index int) {
	for arrary[index] > arrary[(int(index-1)/2)] {
		swap(arrary, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func swap(arrary []int, a, b int) {
	arrary[a], arrary[b] = arrary[b], arrary[a]
}

func heapify(arrary []int, index, size int) {
	left := index*2 + 1 // left child node
	var largest = 0
	for left  < size {
		if left+1 < size && arrary[left+1] > arrary[left] {
			largest = left + 1
		} else {
			largest = left
		}

		if arrary[largest] <= arrary[index] {
			largest = index
		}

		if largest == index {
			break
		}
		swap(arrary, index, largest)
		index = largest
		left = index*2 + 1
	}

}
