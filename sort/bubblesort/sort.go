package bubblesort

func BubbleSort(arrary []int) {
	size := len(arrary)
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arrary[j] > arrary[j+1] {
				arrary[j], arrary[j+1] = arrary[j+1], arrary[j]
			}
		}
	}
}
