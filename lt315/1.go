package lt315


// 计算右侧小于当前元素的个数

// 1. 暴力法 O(n2)/O(1)	(空间复杂度不算返回数组)
// 反正这题暴力思路很简单，先实现一下
func countSmaller1(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	res := make([]int, n)
	for i:=n-1; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			if nums[j] < nums[i] {res[i]++}
		}
	}
	return res
}


// 2. 动态规划的思想优化暴力解法  或者称 记忆化暴力解
// 空间复杂度仍为O(1)，但时间复杂度最好O(n)，最差O(n2)
func countSmaller2(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	res := make([]int, n)
	res[n-1] = 0
	for i:=n-2; i>=0; i-- {
		// 从自身出发，向右寻找第一个小于等于自己的元素d
		firstNotLarger := i+1		// 右边第一个<=自己的元素的下标
		for firstNotLarger < n-1 {	// 一直遍历到n-2这个元素，如果这个元素还比nums[i]大，那么显然nums[i]就和n-1比了
			if nums[firstNotLarger] <= nums[i] {break}
			firstNotLarger++
		}

		if nums[i]==nums[firstNotLarger] {res[i] = res[firstNotLarger]; continue}
		if nums[i]>nums[firstNotLarger] {res[i] = res[firstNotLarger] + 1; continue}
	}
	return res
}

// 上面的思路是错的
// i<j .  dp[j]右侧比dp[j]大的数不一定比dp[i]大


// 下面参考题解区 Adam Wong和liweiwei1419等题解


// 3. 暴力模拟法 + 二分查找  O(nlogn)/O(n)
// 后序遍历时，维护一个已排序数组sortedNums, 每遍历一个元素，
// 就把该元素添加到sortedNums中(保持升序，便于计算结果，降序当然也是行得通的)
// 这样，每次计算nums[i]的右侧更小数数目时直接在sortedNums中二分查找
func countSmaller3(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	sortedNum, res := make([]int, 0), make([]int, n)
	pos := 0	// 应插入位置
	for i:=n-1; i>=0; i-- {
		pos = binarySearchAndInsert(&sortedNum, nums[i])
		res[i] = pos	// pos是升序的sortNums应插入位置，小于nums[i]的元素就有pos个
	}
	return res
}

// 在升序排列的数组中查找target的应放入位置
// 要注意的是，如果发现a=target，target要放在a之前
// 例如 arr=[1,3,5,6], target=4，arr变为[1,3,4,5,6]，返回2
func binarySearchAndInsert(arr *[]int, target int) int {
	n := len(*arr)
	if n==0 {
		*arr = append(*arr, target)
		return 0
	}

	l, r, mid := 0, n-1, 0
	for l<=r {
		mid = (l+r)/2

		// target应插入第一个位置的情况
		if mid==0 && (*arr)[mid] >= target {
			*arr = append([]int{target}, *arr...)
			return 0
		}
		// target追加到最后的情况
		if mid==n-1 && (*arr)[mid] < target {
			*arr = append(*arr, target)
			return n
		}
		// mid在中间(0, n-1]，并恰好是target应插入位置
		if mid>0 && (*arr)[mid-1] < target && (*arr)[mid] >= target {
			*arr = append((*arr)[:mid+1], (*arr)[mid:]...)
			(*arr)[mid] = target
			return mid
		}
		// target应插入到mid左侧(保留mid)
		if (*arr)[mid] >= target {
			r = mid; continue
		}
		// target应插入到mid右侧(不保留mid)
		if (*arr)[mid] < target {
			l = mid+1; continue
		}
	}
	return l
}


// 4. 暴力模拟法 + 二分查找 (二分查找过程代码优化版)
func countSmaller4(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	sortedNum, res := make([]int, 0), make([]int, n)
	pos := 0	// 应插入位置
	for i:=n-1; i>=0; i-- {
		pos = binarySearchAndInsert(&sortedNum, nums[i])
		res[i] = pos	// pos是升序的sortNums应插入位置，小于nums[i]的元素就有pos个
	}
	return res
}
func binarySearchAndInsert2(arr *[]int, target int) int {
	n := len(*arr)
	if n==0 {
		*arr = append(*arr, target)
		return 0
	}

	l, r, mid := 0, n, 0	// 注意r初始为n
	for l<r {
		mid = (l+r)/2

		if (*arr)[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 结束之后l=r就是应插入位置
	*arr = append((*arr)[:mid+1], (*arr)[mid:]...)	// 能够处理插入到头部和尾部的情况
	(*arr)[r] = target
	return r
}


// 5. 二叉搜索树


