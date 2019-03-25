package quicksort

func QuickSort(array []int, L, R int) {
	if L < R {
		var p = partition(array, L, R)
		QuickSort(array, L, p[0]-1)
		QuickSort(array, p[1]+1, R)
	}
}

// 借鉴荷兰国旗问题优化快排
func partition(arr []int, L, R int) []int {
	less := L - 1
	more := R
	current := L 
	for current < more {
		if arr[current] < arr[R] {
			less++
			swap(arr, current, less)
			current++
		} else if arr[current] > arr[R] {
			more--
			swap(arr, current, more)
		} else {
			current++
		}
	}
	swap(arr, more, R)
	return []int{less + 1, more}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
