package search



// 2. 递归版本二分查找
func binarySearch2(nums []int, target int) bool {
    // 边界检查
    n := len(nums)
    if n == 0 {return false}
    if n == 1 {return nums[0] == target}
    if nums[0] > target || nums[n-1] < target {return false}

    // 递归
    return bs(nums, 0, n-1, target)
}

func bs(nums []int, l,r, target int) bool {
    if l == r {return nums[l] == target}

    mid := (r-l)/2 + l
    if nums[mid] == target {
        return true
    }
    if nums[mid] > target {
        return bs(nums, l, mid-1, target)
    } else {
	    return bs(nums, mid+1, r, target)
    }
}

