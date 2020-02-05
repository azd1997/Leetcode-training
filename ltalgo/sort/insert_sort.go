package ltsort


// 插入排序
//插入排序的算法复杂度仍然是O(n2) ，但其工作原理稍有不同。
// 它总是保持一个位置靠前的 已排好的子表，
// 然后每一个新的数据项被 “插入” 到前边的子表里，排好的子表增加一项。
// 我们认为只含有一个数据项的列表是已经排好的。
// 每排后面一个数据（从 1 开始到 n-1），这 个的数据会和
// 已排好子表中的数据比较。
// 比较时，我们把之前已经排好的列表中比这个数据大的移到它的右边。
// 当子表数据小于当前数据，或者当前数据已经和子表的所有数据比较了时，
// 就可 以在此处插入当前数据项。

// NOTICE! 后序遍历和前序遍历都可以。这里采用前序遍历
// O(n2)/O(1)
func insertSort1(arr []int) []int {

	n := len(arr)
	curNum, j := 0, 0
	for i:=1; i<n; i++ {
		curNum = arr[i]	// 准备找位置插入的数
		j = i
		for ; j>0; j-- {	// 跟前面已排序好的arr[:i]逐一比较
			if curNum < arr[j-1] {
				arr[j] = arr[j-1]	// 将原arr[j-1]的数挪到arr[j]，这里其实是有点像冒泡，把curNum冒到已排序部分合适位置
			} else {	// 找到该放的位置了
				break
			}
		}
		arr[j-1] = curNum	// 把curNum放到前面找到的合适位置上
	}
	return arr
}


// 插入排序API
func InsertSort(arr []int) []int {
	return insertSort1(arr)
}
