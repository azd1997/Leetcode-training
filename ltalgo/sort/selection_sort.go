package ltsort


// 选择排序
// 选择排序提高了冒泡排序的性能，
// 它每遍历一次列表只交换一次数据，即进行一次遍历时找到最大的项，
// 完成遍历后，再把它换到正确的位置。
// 和冒泡排序一样，第一次遍历后，最大的数据项就已归位，
// 第二次遍历使次大项归位。
// 这个过程持续进行，一共需要 n-1 次遍历来排好 n 个数 据，
// 因为最后一个数据必须在第 n-1 次遍历之后才能归位。
//
// NOTICE! 上面的解释是每次将待排序部分的最大值归位。
// 同样的道理，每次将最小值归位，也是一样的。下面代码用最小值
// O(n2)/O(1)
func selectionSort1(arr []int) []int {

	n := len(arr)
	minIdx := 0		// 每一轮遍历中的最小值的下标
	for i:=0; i<n; i++ {
		minIdx = i
		for j:=i+1; j<n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}


// 选择排序API
func SelectionSort(arr []int) []int {
	return selectionSort1(arr)
}