package selectsort

func SelectedSort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
