package ltsort

// 归并排序
//归并排序是一种递归算法，它持续地将一个列表分成两半。
// 如果列表是空的或者 只有一个元素，那么根据定义，
// 它就被排序好了（最基本的情况）。如果列表里的元素超过一个，
// 我们就把列表拆分，然后分别对两个部分调用递归排序。
// 一旦这两个部分被排序好了，然后就可以对这两部分数列进行归并了。
// 归并是这样一个过程：把两个排序好了的列表结合在一起组合成一个单一
// 的有序的新列表。有自顶向下（递归法）和自底向上的两种实现方法。
//


// ===============自顶向下(递归) 实现归并排序 ================
func mergeSort1(arr []int) []int {

	n := len(arr)

	// 自顶向下，将区间不断分解
	_mergeSort(&arr, 0, n-1)

	return arr
}

func _mergeSort(arr *[]int, start, end int) {
	// 退出条件
	if start > end {return}

	// 如果区间长度<15，数列基本有序的概率较大，插入排序比较适合
	if end-start<=15 {
		_insertSort(arr, start, end); return
	}

	// 二分 递归
	mid := start + (end - start) / 2
	_mergeSort(arr, start, mid)
	_mergeSort(arr, mid+1, end)
	// 优化 合并结果
	if (*arr)[mid] > (*arr)[mid+1] {
		_merge(arr, start, mid, end)
	}
}

func _insertSort(arr *[]int, start, end int) {
	curNum, j := 0, 0
	for i:=start+1; i<=end; i++ {
		curNum, j = (*arr)[i], i
		for ; j>=start+1; j-- {
			if curNum < (*arr)[j-1] {
				(*arr)[j] = (*arr)[j-1]
			} else {break}
		}
		(*arr)[j-1] = curNum
	}
}

// 合并两个有序数列 A : arr[start:mid+1] 和 B : arr[mid+1:end+1]
func _merge(arr *[]int, start, mid, end int) {
	// 复制一份数据
	copyArr := make([]int, end-start+1)
	copy(copyArr, (*arr)[start:end+1])

	l := start	// A的指针
	k := mid+1  // B的指针
	pos := start	// 待修改位置

	// 合并两个有序序列
	for pos <= end {
		// 处理pos位置数据如何设置
		if l>mid {
			(*arr)[pos] = copyArr[k-start]
			k++
		} else if k > end {
			(*arr)[pos] = copyArr[l-start]
			l++
		} else if copyArr[l-start] <= copyArr[k-start] {
			l++
		} else {
			k++
		}
		// pos后移
		pos++
	}
}


// ===============自底向上(非递归) 实现归并排序 ================

func mergeSort2(arr []int) []int {

	n := len(arr)

	// 表示归并的大小
	size := 1

	// 自底向上
	for size <= n {
		for i:=0; i<n-size; i+=size+size {
			// end这么设置是为了处理最后一个分块长度不足的情况
			_merge2(&arr, i, i+size-1, min(i+size+size-1, n-1))
		}
		size += size
	}

	return arr
}

func min(a, b int) int {
	if a<=b {return a} else {return b}
}

// 合并两个有序数列 A : arr[start:mid+1] 和 B : arr[mid+1:end+1]
func _merge2(arr *[]int, start, mid, end int) {
	// 复制一份数据
	copyArr := make([]int, end-start+1)
	copy(copyArr, (*arr)[start:end+1])

	l := start	// A的指针
	k := mid+1  // B的指针
	pos := start	// 待修改位置

	// 合并两个有序序列
	for pos <= end {
		// 处理pos位置数据如何设置
		if l>mid {
			(*arr)[pos] = copyArr[k-start]
			k++
		} else if k > end {
			(*arr)[pos] = copyArr[l-start]
			l++
		} else if copyArr[l-start] <= copyArr[k-start] {
			(*arr)[pos] = copyArr[l-start]
			l++
		} else {
			(*arr)[pos] = copyArr[k-start]
			k++
		}
		// pos后移
		pos++
	}
}










// 归并排序API
func MergeSort(arr []int) []int {
	return mergeSort1(arr)
}