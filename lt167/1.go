package lt167

// 两数之和II - 输入有序数组

// 既然数组是有序的，那么可以从数组两端向内双指针来找两数之和
// 而不需要使用哈希表等额外空间

// 1. 一遍线性 双指针
func twoSum(numbers []int, target int) []int {
	// 特殊
	n := len(numbers)
	if n < 2 {
		return []int{-1, -1}
	}

	// 双指针
	l, r := 0, n-1
	for l+1 <= r { // 至少有两个元素
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1} // 返回的下标从1开始
		} else if sum < target {
			l++
		} else { // >
			r--
		}
	}
	return []int{-1, -1} // 没找到
}

// 由于题目只要求找到一个解，并且数组有序，那么可以使用二分查找 O(nlogn)
func twoSum2(numbers []int, target int) []int {
	// 特殊
	n := len(numbers)
	if n < 2 {
		return []int{-1, -1}
	}

	// 外部线性遍历，内部二分查找
	for i := 0; i < n-1; i++ { // 至少要剩余一个元素，才能进行二分
		t := target - numbers[i]
		if j := bs(numbers, i+1, n-1, t); j != -1 {
			return []int{i + 1, j + 1}
		}
	}

	return []int{-1, -1} // 没找到
}

func bs(arr []int, start, end, target int) int {
	l, r, mid := start, end, 0
	for l <= r {
		mid = (r-l)/2 + l
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
