package lt88

import "sort"

// 合并两个有序数组
// 假设nums1(长度m)空间足够，nums2(长度n)并入nums1

// 解法：
// 1. 最蠢最暴力解法：遍历nums2,将nums2[i]从前向后与nums1[j]比较，当nums2[i]>nums1[j]时插到其后
// 		（或者当nums2[i]<nums1[j]时插到其前，这种少一点数据搬迁工作量，而多了一些数据大小比较的工作量）
// 		时间复杂度O(n*m/2*m/2) 空间复杂度O(1)
// 		n是nums2遍历循环，第一个m是遍历nums1大小比较，第二个m是找到插入位置后插入操作
// 2. 在解法1基础上优化一些：用一个变量index标记nums2[i]插入到nums1的位置，nums2[i+1]就直接从这个位置起开始比较nums1的数据
//		时间复杂度比较难求，但显然比解法1高效许多
// 3. 在解法2基础上加上二分查找比较，比一个一个比较快多了，毕竟二分查找本身时间复杂度是O(logn)
// 4. 标记加二分查找使得找到插入位置变得高效，但是还有数据搬迁仍然非常耗时。在解法4中引入辅助数组a=[]int。
//		这个辅助数组有两种用法：
// 		4.1 如果不要求nums1指针不变，直接用a=[m+n]int存合并后的有序数组，然后再让nums1指向a。
// 			这种做法下如果要保证nums1指针不变，那么就将a数据拷贝给nums1
// 		4.2 如果要求nums1要求指针不变，除了上面的4.1后一种做法外，还以用这个辅助数组来存储nums1和nums2(这里a可以是一个m+n的，也可以是拆成两个)
// 			各数据的新下标位置并不实际搬移数据，等到a记录了所有数据的新位置，再计算搬移量来进行搬移
// 5. 还能不能更快呢？

// -------------------------------------------------------------------------------------
// 参考官方题解等，有以下方法：
// 1. 合并两个数组再排序

//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]

// 限于时间精力，除了朴素合并再排序外，只实现原地解法及时间复杂度较低的解法

// 合并两个数组再排序
// 时间复杂度O((m+n)log(m+n))，空间复杂度O(1)
// ！！！没有利用原数组有序这一特性，所以有提升空间
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 94.16 % of golang submissions (3.6 MB)
func Sol_1_1(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		//nums1 = append(nums1, nums2...)		// 原地
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下合并两数组再重新排序
	//nums1 = append(nums1, nums2...)
	for i:=m; i<n+m; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Sort(sort.IntSlice(nums1))		// 时间复杂度O((m+n)log(m+n))且不稳定
}

// 双指针从前向后
// 每次从nums1原有效数据的备份nums1_1数组和nums2的“头部”（已经填入nums1的不计）取出数据进行比较，小者先填入nums1。
// 如果相等，随便哪个先入，这里让nums1_1的“头部”值先入
// 时间复杂度O(m+n), 空间复杂度O(m)
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 75.32 % of golang submissions (3.6 MB)
func Sol_1_2(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		//nums1 = append(nums1, nums2...)		// 原地
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下
	nums1_1 := append([]int{}, nums1[:m]...)	// 将nums1前m个有效数据拷贝至nums1_1

	var p1, p2 int	// p1,p2分别代表nums1_1, nums2 "头部"数据的下标
	for i:=0; i<m+n; i++ {
		// p1,p2有个限制
		if p1 == m {	// 如果p1先到底，则将p2后续直接复制过来
			// p1, p2之前各有p1,p2个数据已经填到nums1,所以nums1现在有的应该是p1+p2个，那么剩下的就是nums1[p1+p2:]
			copy(nums1[p1+p2:], nums2[p2:])
			break
		}
		if p2 == n {
			copy(nums1[p1+p2:], nums1_1[p1:])
			break
		}

		if nums1_1[p1] <= nums2[p2] {
			nums1[i] = nums1_1[p1]
			p1++	// p1后移
		} else {
			nums1[i] = nums2[p2]
			p2++	// p2后移
		}

		// 这样写，那么最后p1p2一般会有其中一个先到末尾，就不能在继续加了，也就是有个p1<m,p2<n的限制。

	}
	// 最后nums1就得到排好序的合并数组了
}

// Sol_1_2中的循环语句也可以进行改写，如下所示，本质上两者是一个方法
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 69.81 % of golang submissions (3.6 MB)
func Sol_1_3(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		//nums1 = append(nums1, nums2...)		// 原地
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下
	nums1_1 := append([]int{}, nums1[:m]...)	// 将nums1前m个有效数据拷贝至nums1_1

	var p1, p2 int	// p1,p2分别代表nums1_1, nums2 "头部"数据的下标
	// p1,p2未到末尾之前，两个数组的“头部”数据都需要进行大小比较
	for p1<m && p2<n {
		if nums1_1[p1] <= nums2[p2] {
			nums1[p1+p2] = nums1_1[p1]		// 在把当前这个数据加到nums1之前，nums1已经有了p1+p2个数据，所以nums1即将填充的数据是nums[p1+p2]
			p1++	// p1后移
		} else {
			nums1[p1+p2] = nums2[p2]
			p2++	// p2后移
		}
	}

	// 退出循环之后，p1, p2可能有其中之一还没到末尾
	if p1 < m {	// 如果p1先到底，则将p2后续直接复制过来
		copy(nums1[p1+p2:], nums1_1[p1:])
	}
	if p2 < n {
		copy(nums1[p1+p2:], nums2[p2:])
	}
	// 最后nums1就得到排好序的合并数组了
}

// 其实Sol_1_3还有可优化的的地方。如果遍历到的两个“头部数据”相等的情况很多，可以一次将两个数据都填到nums1去
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 75.32 % of golang submissions (3.6 MB)
func Sol_1_4(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		//nums1 = append(nums1, nums2...)		// 原地
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下
	nums1_1 := append([]int{}, nums1[:m]...)	// 将nums1前m个有效数据拷贝至nums1_1

	var p1, p2 int	// p1,p2分别代表nums1_1, nums2 "头部"数据的下标
	// p1,p2未到末尾之前，两个数组的“头部”数据都需要进行大小比较
	for p1<m && p2<n {
		if nums1_1[p1] == nums2[p2] {
			nums1[p1+p2], nums1[p1+p2+1] = nums2[p2], nums2[p2]
			p1++
			p2++
		} else if nums1_1[p1] < nums2[p2] {
			nums1[p1+p2] = nums1_1[p1]
			p1++	// p1后移
		} else {
			nums1[p1+p2] = nums2[p2]
			p2++	// p2后移
		}
	}

	// 退出循环之后，p1, p2可能有其中之一还没到末尾
	if p1 < m {	// 如果p1先到底，则将p2后续直接复制过来
		copy(nums1[p1+p2:], nums1_1[p1:])
	}
	if p2 < n {
		copy(nums1[p1+p2:], nums2[p2:])
	}
	// 最后nums1就得到排好序的合并数组了
}

// 双指针法，从后向前。
// Sol_1_2~Sol_1_4都使用了一个额外的数组，但是nums1数组本身后边原本是零值的，不需要保存，如果从后向前填充，则不需要额外数组
// 时间O(n), 空间O(1)
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 94.16 % of golang submissions (3.6 MB)
func Sol_1_5(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下,从后向前遍历
	p1 := m-1	// p1,p2分别代表nums1, nums2 原本状态中有效数据最后一个的下标，每次当最后这个数据被填充到nums1末尾时，p1或p2前移
	p2 := n-1
	// p1,p2未到最左边之前，两个数组的“尾部”数据都需要进行大小比较
	for p1>=0 && p2>=0 {
		if nums1[p1] == nums2[p2] {
			nums1[p1+p2+1], nums1[p1+p2] = nums2[p2], nums2[p2]		// (n+m-1)-(m-p1-1)-(n-p2-1) = p1+p2+1 是当前比较的大者应该填充的位置
			p1--
			p2--
		} else if nums1[p1] > nums2[p2] {
			nums1[p1+p2+1] = nums1[p1]
			p1--	// p1前移
		} else {
			nums1[p1+p2+1] = nums2[p2]
			p2--	// p2前移
		}
	}

	// 退出循环之后，p1, p2可能有其中之一还没到末尾.
	// 如果是p1还没到nums1最左边（越过），则不用干什么。
	// 如果是p2还没到nums2最左边（越过），则将其拷贝给nums1
	if p2 >= 0 {
		copy(nums1[:p1+p2+2], nums2[:p2+1])		// 注意go切片[a:b]右半边b是不包含的
	}
	// 最后nums1就得到排好序的合并数组了
}

// 如果觉得前面Sol_1_5中nums1填入数据的最终下标有点绕，我们可以给它一个指针p,p从n+m-1开始递减
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 94.16 % of golang submissions (3.6 MB)
func Sol_1_6(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下,从后向前遍历
	p1 := m-1	// p1,p2分别代表nums1, nums2 原本状态中有效数据最后一个的下标，每次当最后这个数据被填充到nums1末尾时，p1或p2前移
	p2 := n-1
	p := m+n-1	// p指向两数组“尾部”数据大者将要填入的位置，递减
	// p1,p2未到最左边之前，两个数组的“尾部”数据都需要进行大小比较
	for p1>=0 && p2>=0 {
		if nums1[p1] == nums2[p2] {
			nums1[p], nums1[p-1] = nums2[p2], nums2[p2]		// (n+m-1)-(m-p1-1)-(n-p2-1) = p1+p2+1 是当前比较的大者应该填充的位置
			p1--
			p2--
			p = p-2
		} else if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--	// p1前移
			p--
		} else {
			nums1[p] = nums2[p2]
			p2--	// p2前移
			p--
		}
	}

	// 退出循环之后，p1, p2可能有其中之一还没到末尾.
	// 如果是p1还没到nums1最左边（越过），则不用干什么。
	// 如果是p2还没到nums2最左边（越过），则将其拷贝给nums1
	if p2 >= 0 {
		copy(nums1[:p+1], nums2[:p2+1])		// 注意go切片[a:b]右半边b是不包含的
	}
	// 最后nums1就得到排好序的合并数组了
}

// -------------------------------------------------------最后再来实现下我最初的解题想法

// index标记之前的插入位置，使用二分查找插入
func Sol_1_7(nums1, nums2 []int, m, n int) {

	// 不检查m,n是否是nums1,nums2长度，题已给定

	// 边界条件 m=0 || n=0
	if m == 0 {
		for i:=0; i<n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	// 一般情况下,使用我前面提到的思路
	// 遍历nums2,将nums2[i]和nums1[j]比较去插入，如果使用辅助数组来填充数，那么本质上和前面的的双指针法从前向后没区别


}