package lt658

import (
	"fmt"
	"math"
	"sort"
)

// 找到K个最接近的元素

// 思路：
// 先二分查找找到target的左边界idx（注意不要直接按模板，如果nums[l]!=target返回l）
// 再从idx向两边寻找比较。这时要使用双指针，来比较哪边的差值更小。差值一样的话左边排前面

func findClosestElements(arr []int, k int, x int) []int {
	// 特殊情况
	n := len(arr)
	if n == 0 || n < k {
		return nil
	}

	// 二分，找到target左边界或左邻
	idx := bsl(arr, x)
	fmt.Println("idx=", idx)
	if idx == n {
		return arr[n-k:]
	}
	if idx == 0 {
		return arr[:k]
	}
	// 现在以idx为界，分为两半，右侧以idx为起点，左侧以idx-1为起点， 向两边双指针
	res := make([]int, k)
	k1 := 0
	l, r := idx-1, idx
	for k1 < k {
		// 两端都有剩余元素
		if l >= 0 && r < n {
			if x-arr[l] <= arr[r]-x { // 注意=时取左边的（更小）
				res[k1] = arr[l]
				l--
			} else {
				res[k1] = arr[r]
				r++
			}
			k1++
			continue
		}
		// 一端没有
		if l >= 0 {
			res[k1] = arr[l]
			l--
			k1++
		}
		if r < n {
			res[k1] = arr[r]
			r++
			k1++
		}
	}

	// 要求res升序排列
	sort.Ints(res)

	return res
}

func bsl(arr []int, target int) int {
	l, r, mid := 0, len(arr), 0
	for l <= r {
		mid = (r-l)/2 + l
		if arr[mid] >= target {
			r = mid - 1 // 向左搜索
		} else {
			l = mid + 1 // 向右搜索
		}
	}
	// l 可能有哪些可能？
	// target比arr都小，则不断向左缩，最后arr[0]还是比target大，则r=-1，但是l=0. 也就是说target比arr都小，l=0
	// target比arr都大，不断向右，最后arr[n-1]还是比target小，则l=n
	// target存在于arr，则l=0~n-1
	// target在arr区间范围内但不等于任何一个数，最后一次区间可能是target左邻也可能是右邻。
	//    如果是左邻，那么l会增大到target右邻位置; 如果是右邻，则l就落在右邻位置

	// 在本题，其实都可以直接返回
	return l
}

/////////////////////////////////////////////////////

// 解法二： 根据差值直接排序
func findClosestElements2(arr []int, k int, x int) []int {
	// 特殊情况
	n := len(arr)
	if n == 0 || n < k {
		return nil
	}

	// 根据差值排序
	sort.Slice(arr, func(i, j int) bool {
		a, b := int(math.Abs(float64(arr[i]-x))), int(math.Abs(float64(arr[j]-x)))
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})

	// res升序
	res := arr[:k]
	sort.Ints(res)

	return res
}

// https://leetcode-cn.com/problems/find-k-closest-elements/solution/pai-chu-fa-shuang-zhi-zhen-er-fen-fa-python-dai-ma/
// 这篇题解提供了两种更好的解法
