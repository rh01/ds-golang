package mergesort

func merge(array []int, leftIndex, midIndex, rightIndex int) {
	// 初始化两个
	var temp1Array, temp2Array []int

	for i := leftIndex; i <= midIndex; i++ {
		temp1Array = append(temp1Array, array[i])
	}

	for i := midIndex + 1; i <= rightIndex; i++ {
		temp2Array = append(temp2Array, array[i])
	}

	var temp1Index = 0
	var temp2Index = 0
	var arrayIndex = leftIndex // 注意！！
	for temp1Index != len(temp1Array) && temp2Index != len(temp2Array) {
		if temp1Array[temp1Index] <=  temp2Array[temp2Index]  {
			array[arrayIndex] = temp1Array[temp1Index]
			temp1Index++
		}else{
			array[arrayIndex] = temp2Array[temp2Index]
			temp2Index++
		}
		arrayIndex ++
	}

	for temp1Index < len(temp1Array) {
		array[arrayIndex] = temp1Array[temp1Index]
		arrayIndex++
		temp1Index++
	}

	for temp2Index < len(temp2Array) {
		array[arrayIndex] = temp2Array[temp2Index]
		arrayIndex++
		temp2Index++
	}

}

func MergeSort(arr []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	midIndex := int((leftIndex + rightIndex) / 2)
	MergeSort(arr, leftIndex, midIndex)
	MergeSort(arr, midIndex+1, rightIndex)
	merge(arr, leftIndex, midIndex, rightIndex)
}
