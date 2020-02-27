package ltsort


// 希尔排序
//希尔排序有时又叫做 “缩小间隔排序”，它以插入排序为基础，
// 将原来要排序的列表划分为一些子列表，再对每一个子列表执行插入排序，
// 从而实现对插入排序性能的改进。划分子列的特定方法是希尔排序的关键。
// 我们并不是将原始列表分成含有连续元素的子列，而是确定一个划分列表的
// 增量 “i”，这个i更准确地说，是划分的间隔。然后把每间隔为i的所有元素
// 选出来组成子列表，然后对每个子序列进行插入排序，最后当 i=1 时，对
// 整体进行一次直接插入排序。

// NOTICE! 后序遍历和前序遍历都可以。这里采用前序遍历
// O(n2)/O(1)
func shellSort1(arr []int) []int {

	n := len(arr)

	gap := n / 2	// 间隔初设为n/2，也就是说把arr一开始分成n/2个长度为2的数列进行插入排序
	for gap>0 {
		for i:=0; i<gap; i++ {
			gapInsertSort(&arr, i, gap)
		}
		gap /= 2
	}
	// 这里时间复杂度的计算：
	// 首先外层gap更新这里是 log_2_{n}	记作logn
	// 内层循环： gap*(n/gap)^2 / 2 = (n^2/gap) / 2
	// 也就是： n + 2n + ... + (n/4)n =

	return arr
}

// 令m = n/gap
// 时间复杂度 约为 m^2/2
func gapInsertSort(arr *[]int, start, gap int) {
	n := len(*arr)
	curNum, j := 0, 0
	for i:=start+gap; i<n; i+=gap {
		j = i
		curNum = (*arr)[i]
		for ; j>=start+gap; j-- {
			if curNum < (*arr)[j-1] {
				(*arr)[j] = (*arr)[j-1]
			} else {break}
		}
		(*arr)[j] = curNum
	}
}


// 希尔排序API
func ShellSort(arr []int) []int {
	return shellSort1(arr)
}

