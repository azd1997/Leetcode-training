package lt75

// 颜色分类

// 题目提示计数排序，但我没学过计数排序
// 但是这道题其实很简单，因为只有三种颜色，只需要类似快排尤其是三路快排的那种做法。
// 这里：
// 凡是红色(0)交换到左边区间，凡是蓝色(2)交换到右边区间
// 一次遍历就可以排好序了。
//这只用到了三个指针（数组游标），符合题意

func sortColors(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	l, r, i := 0, n-1, 0 // l,r都指向区间右邻/左邻位置，而不是区间的右末位/左末位
	for i <= r {         // i==r时仍然要处理最后这个数
		switch nums[i] {
		case 0: // 红色
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		case 1: // 白
			i++
		case 2: // 蓝
			nums[r], nums[i] = nums[i], nums[r]
			r-- // 注意i不变，因为新交换过来的还没有处理过
		}
	}
}
