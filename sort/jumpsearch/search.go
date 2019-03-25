package jumpsort

import "math"

// 假定有序列表中查询某个key，返回该key的索引
func JumpSearch(arrary []int, key int) int {
	jump := int(math.Floor(math.Sqrt(float64(len(arrary)))))
	minIndex := 0
	maxIndex := jump
	for arrary[maxIndex] <= key {
		minIndex += jump
		maxIndex = minIndex + jump
		if maxIndex >= len(arrary) {
			minIndex = minIndex - jump
			maxIndex = len(arrary)
			break
		}
	}

	for i := minIndex; i < maxIndex; i++ {
		if arrary[i] == key {
			return i
		}
	}
	return -1

}
