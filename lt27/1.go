package lt27

// 移除元素
// 原地移除数组中与给定值相等的元素，返回移除后数组的长度
// 只能使用O(1)额外空间
// 元素顺序可以改变
// 不需要考虑数组中超过新长度以后的元素

// nums=[3,2,2,3], val=3
// 返回新长度2

// 1. 快慢指针
//113/113 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.1 MB)
func removeElement1(nums []int, val int) int {

	// 思路： 快慢指针，快指针向后遍历，慢指针遇val则停，交换快慢指针所指元素后慢指针方可后移
	// 注意： 这样做，结果数组中原本元素顺序很可能被打乱，但满足题意

	// 1. 处理边界情况
	if len(nums) == 0 {return 0}

	// 2. 快慢指针移动
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != val {
			//nums[i], nums[j] = nums[j], nums[i]
			nums[j] = nums[i]
			j++
		}
	}

	// 3. 返回新数组长度
	return j
}

// 2. 双指针-优化 删除元素较少时
// 遇val则与末位值交换

func removeElement2(nums []int, val int) int {

	// 1. 处理边界情况
	if len(nums) == 0 {return 0}

	// 2. 双指针移动. j 代表数组长度，也相当于指针从后往前移动
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}

	// 3. 返回新数组长度
	return j
}