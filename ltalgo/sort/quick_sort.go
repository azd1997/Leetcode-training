package ltsort

import "math/rand"

// 快速排序
//快速排序由 C. A. R. Hoare 在1962年提出。
// 它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，
// 其中一部分的所有数据都比另外一部分的所有数据都要小，
// 然后再按此方法对这两部分数据分别进行快速排序，
// 整个排序过程可以递归进行，以此达到整个数据变成有序序列。

// 算法步骤
// 1.从数列中挑出一个元素，称为"基准"（pivot）。
// 2.重新排序数列，所有比基准值小的元素摆放在基准前面，
// 所有比基准值大的元素摆在基准后面（相同的数可以到任何一边）。
// 在这个分区结束之后，该基准就处于数列的中间位置。
// 这个称为分区（partition）操作。
// 3.递归地（recursively）把小于基准值元素的子数列
// 和大于基准值元素的子数列排序。

func quickSort1(arr []int) []int {
	_quickSort(&arr, 0, len(arr)-1)
	return arr
}

func _quickSort(arr *[]int, start, end int) {
	// 退出递归条件
	if start>end {return}

	// 优化：数列长度较小时，使用插入排序
	if end-start<=15 {
		_insertSort(arr, start, end); return
	}

	// 对区域进行分区，找到一个分割点p
	p := _partition(arr, start, end)

	_quickSort(arr, start, p-1)
	_quickSort(arr, p+1, end)
}

func _partition(arr *[]int, start, end int) int {
	// 随机产生一个基准下标
	pos := rand.Intn(end-start) + start

	// 交换第一个和基准的位置
	(*arr)[pos], (*arr)[start] = (*arr)[start], (*arr)[pos]

	v := (*arr)[start]	// 此时v存储的是随机选出来的基准值
	i, j := start+1, start
	for i<=end {
		if (*arr)[i] <= v {		// 当前数小于基准值，则将其与前一位交换
			(*arr)[j+1], (*arr)[i] = (*arr)[i], (*arr)[j+1]
			j++
		}
		i++
	}
	// 将j位置的值与其实基准值置换
	(*arr)[j], (*arr)[start] = (*arr)[start], (*arr)[j]

	return j	// 分区的位置，j前 < j后
}



// 快速排序一些可以优化的点:
// 1. 当数列近乎有序的时， 由于每次选取的都是第一个数，
// 所以造成数列分割的极其不等， 此时快排蜕化成  的算法，
// 此时只要随机选取基准点即可
// 2. 当数列中包含大量的重复元素的时候，这一版的代码也
// 会造成"分割不等“的问题， 此时需要将重复元素均匀的分散的自数列旁
// 3. 使用三路快排
//




// 快速排序API
func QuickSort(arr []int) []int {
	return quickSort1(arr)
}