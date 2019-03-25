package radixsort

import "math"

const radix = 10

func RadixSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	radixSort(arr, 0, len(arr)-1, maxbits(arr))
}

func maxbits(arr []int) int {
	var max = math.MinInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	var res = 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
}

func radixSort(arr []int, begin, end, digit int) {
	// var i = 0
	var j = 0
	count := make([]int, radix)
	bucket := make([]int, end-begin+1)
	for d := 1; d <= digit; d++ {
		for i := 0; i < radix; i++ {
			count[i] = 0
		}
		for i := begin; i <= end; i++ {
			j = getDigit(arr[i], d)
			count[j]++
		}
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		for i := end; i >= begin; i-- {
			j = getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		j := 0
		for i := begin; i <= end; i++ {
			arr[i] = bucket[j]
			j++
		}
	}
}

func getDigit(x, d int) int {
	return (x / int(math.Pow10(d-1)) % 10)
}
