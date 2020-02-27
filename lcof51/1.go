package lcof51

import "fmt"

// 数组中的逆序对


// 基于归并排序思想：在归并排序的过程中统计逆序对数 O(nlogn)/O(n)
func reversePairs(nums []int) int {
	// 边界情况
	n := len(nums)
	if n < 2 {return 0}

	// 准备一个数组，给每次合并使用
	numsCopy := make([]int, n)

	return _mergeSort(&nums, &numsCopy, 0, n-1)
}

// 归并排序辅助函数。返回这一轮发现的逆序对数
func _mergeSort(nums, numsCopy *[]int, l, r int) int {
	if l >= r {return 0}

	// 区间二分
	mid := (r-l) / 2 + l
	a := _mergeSort(nums, numsCopy, l, mid)
	b := _mergeSort(nums, numsCopy, mid+1, r)

	// 合并。 合并过程中也会发现逆序对
	c := 0
	if (*nums)[mid] > (*nums)[mid+1] {      // 否则两个区间直接组合就是有序的，不存在逆序对
		c = _merge(nums, numsCopy, l, mid, r)
	}

	return a+b+c    // 返回逆序对总数
}

// merge阶段，返回merge发现的逆序对数
func _merge(nums, numsCopy *[]int, l, mid, r int) int {
	// merge阶段实际是对两个子区间采用双指针的方式来O(n)排序合并

	// NOTICE: 左右区间内部已经是升序状态，不存在逆序对
	// 这个merge阶段是为了求合并时发现的逆序对数

	// 备份当前左右区间
	copy((*numsCopy)[:r-l+1], (*nums)[l:r+1])

	total := 0      // 当前左右区间合并时统计得到的逆序对数

	// 这种情况下从前向后和从后向前移动没区别，这里从前向后
	p := l
	p1, p2 := l, mid+1
	for ; p<=r; p++ {    // p走到r就结束

		// p1先走到底，则后边依次填充右区间剩下的部分，逆序对不需要增加
		if p1 > mid {
			(*nums)[p] = (*numsCopy)[p2-l]
			p2++
		} else if p2 > r {
			(*nums)[p] = (*numsCopy)[p1-l]
			p1++
			// 都没走到底的比较(ps: 针对大量重复元素，还可以利用二分查找优化，快速越过重复地带，这里按下不表)
		} else if (*numsCopy)[p1-l] > (*numsCopy)[p2-l] {
			// 这样排序依然是稳定的，且满足题目要求： (前 > 后) 称为一个逆序对
			(*nums)[p] = (*numsCopy)[p2-l]		// TODO： 右区间每出一个元素，都意味着左区间剩下的还没出列的元素与其构成逆序对。并且当右区间用完后，左区间剩下的元素无需再叠加逆序对数，因为已被统计掉了
			total += mid+1-p1     // 发现一个逆序对
			// 调试用
			for i:=p1; i<=mid; i++ {
				fmt.Printf("total=%d, 新增逆序对 <%d, %d>\n", total, (*numsCopy)[i-l], (*numsCopy)[p2-l])
			}
			p2++
		} else {
			(*nums)[p] = (*numsCopy)[p1-l]
			p1++
		}
	}

	return total
}
