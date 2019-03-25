package netherlandsflag

// 荷兰国旗问题
func partition(arr []int, L, R int, num int) []int {
	less := L - 1
	more := R + 1
	current := L
	for current < more {
		if arr[current] < num {
			less++
			swap(arr, less, current)
			current++
		} else if arr[current] > num {
			more--
			swap(arr, current, more)
		} else {
			current++
		}
	}
	return []int{less - 1, more + 1}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
