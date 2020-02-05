package ltsort

// 冒泡排序
// 多次遍历，相邻者两两比较，若逆序则交换。
// 每一轮遍历都把最大值冒泡到后部的正确位置
// 直至最后某一轮遍历全部升序没有逆序则停止
// O(n^2)/O(1)
func bubbleSort1(arr []int) []int {
	n := len(arr)
	for i:= n-1; i>=0; i-- {
		for j:=0; j<i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}


// 冒泡排序优化
// 如果整个排序过程中没有交换，提前结束。(也就是前面说的某一轮遍历发现没有逆序)
func bubbleSort2(arr []int) []int {
	n := len(arr)
	existExchange := false
	for i:= n-1; i>=0; i-- {
		for j:=0; j<i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				existExchange = true
			}
		}
		if !existExchange {break}	// 如果该轮遍历不存在交换，那么提前结束
	}
	return arr
}


// 冒泡排序API
// 冒泡排序的优点在于：对于逆序度较低的序列，能够提前通过检测交换发现来决定是否提前结束
func BubbleSort(arr []int) []int {
	return bubbleSort2(arr)
}