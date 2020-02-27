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






// 普通快排 + 随机化基数优化 + 插入排序优化


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

/////////////////////////////////////////////////////////

// 双路快排

func quickSort2(arr []int) []int {
	_quickSort2(&arr, 0, len(arr)-1)
	return arr
}

func _quickSort2(arr *[]int, start, end int) {
	// 退出递归条件
	if start>=end {return}

	// 优化：数列长度较小时，使用插入排序
	if end-start<=15 {
		_insertSort(arr, start, end); return
	}

	// 对区域进行分区，找到一个分割点p
	p := _partition2(arr, start, end)

	_quickSort2(arr, start, p-1)
	_quickSort2(arr, p+1, end)
}

// 双路快排分区
func _partition2(nums *[]int, l, r int) int {
	// 随机一个元素(也就是)作为基数base
	randIdx := rand.Intn(r-l+1) + l
	(*nums)[randIdx], (*nums)[l] = (*nums)[l], (*nums)[randIdx]
	// 现在base = nums[l]

	////////////////////////////////////////////
	// 以下为修改过的代码

	// 满足 nums[l+1:i+1] <= nums[l] <= nums[j:r+1]
	i, j := l+1, r
	for {
		// i右移
		for i<=r && (*nums)[i]<(*nums)[l] {i++}
		// j左移
		for j>l && (*nums)[j]>(*nums)[l] {j--}
		// 是否已经遍历结束
		if i>j {break}
		// 交换i,j并继续移动
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		i, j = i+1, j-1
	}
	// 遍历结束后， i为从左向右看第一个 >=nums[l] 的元素
	// j 为i为从右向左看第一个 <=nums[l] 的元素
	// 因此 nums[l] 应该和 nums[j] 交换
	// 才能继续保证 base 左侧都 <=base， 右侧都 >= base
	(*nums)[l], (*nums)[j] = (*nums)[j], (*nums)[l]

	return j
}

/////////////////////////////////

func quickSort3(arr []int) []int {
	_quick3(&arr, 0, len(arr)-1)
	return arr
}


func _quick3(nums *[]int, l, r int) {
	// 递归终止(没法再partition)
	if l >= r {return}	// 区间为1时没必要再排序
	// 区间较小时使用插入排序
	if r - l <= 15 {
		_insertSort(nums, l, r)
		return
	}

	//////////////////////////////////

	// partition得到分界点
	p1, p2 := _partition3(nums, l, r)
	// 继续对左右进行递归处理
	_quick3(nums, l, p1-1)
	_quick3(nums, p2+1, r)

	//////////////////////////////////
}

// 三路快排分区
func _partition3(nums *[]int, l, r int) (int, int) {
	// 随机一个元素(也就是)作为基数base
	randIdx := rand.Intn(r-l+1) + l
	(*nums)[randIdx], (*nums)[l] = (*nums)[l], (*nums)[randIdx]
	// 现在base = nums[l]

	////////////////////////////////////////////
	// 以下为修改过的代码

	// 满足 nums[l+1:lt+1] < nums[lt+1:i-1] <= nums[gt:r+1]
	// 保证初始时 Left/Right/Mid 均为空
	lt, gt, i := l, r+1, l+1

	for i < gt {    // 如果i与gt相遇，遍历结束. 而且这个限制保证了 i 不会越界
		// 等于base时 i 右移
		if (*nums)[i] == (*nums)[l] {
			i++; continue
		}
		// 小于base i右移lt右移
		if (*nums)[i] < (*nums)[l] {
			(*nums)[i], (*nums)[lt+1] = (*nums)[l+1], (*nums)[i]
			i++; lt++; continue
		}
		// 大于base gt左移, 新交换过来的 i 仍需处理
		if (*nums)[i] > (*nums)[l] {
			(*nums)[i], (*nums)[gt-1] = (*nums)[gt-1], (*nums)[i]
			gt--; continue
		}
	}
	// 交换 l 与 lt
	(*nums)[l], (*nums)[lt] = (*nums)[lt], (*nums)[l]
	// 返回p1,p2. p1=lt, p2=gt-1
	return lt, gt-1
}




////////////////////////////////////

// 快速排序API
func QuickSort(no int) func([]int) []int {
	switch no {
	case 1:
		return quickSort1
	case 2:
		return quickSort2
	case 3:
		return quickSort3
	default:
		return nil
	}
}

const (
	QuickSortNormal = 1
	QuickSort2Way = 2
	QuickSort3Way = 3
)