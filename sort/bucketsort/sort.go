package bucketsort

import "math"

func BucketSort(array []int) {
	// 1. 检查当前 array 大小
	size := len(array)
	if size == 0 || size < 2 {
		return
	}

	// 2. 构造一个桶
	var max = math.MinInt64
	for i := 0; i < size; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	bucketArrary := make([]int, max+1)
	for i := 0; i < size; i++ {
		bucketArrary[array[i]]++
	}

	// 3. 桶排序主逻辑
	i := 0
	for j := 0; j < len(bucketArrary); j++ {
		for ; bucketArrary[j] > 0; bucketArrary[j]-- {
			array[i] = j
			i++
		}
	}

}
