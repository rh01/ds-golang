package insertsort

func InsertSort(arr []int) {
	size := len(arr)
	for i := 1; i < size; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
