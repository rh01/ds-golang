package randomquicksort

import "math/rand"

func RandomQuickSort(array []int, L, R int) {
	if L < R {
		swap(array, int(rand.New(rand.NewSource(1024)).Float64())*(R-L+1), R)
		var p = partition(array, L, R)
		RandomQuickSort(array, L, p[0]-1)
		RandomQuickSort(array, p[1]+1, R)
	}
}

// 借鉴荷兰国旗问题优化快排
func partition(arr []int, L, R int) []int {
	less := L - 1
	more := R
	current := L
	// random := rand.New(rand.NewSource(1024))
	// index := random.Intn(R-L) + L + 1
	// swap(arr, index, R)
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
