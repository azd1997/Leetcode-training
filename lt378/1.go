package lt378

import (
	"container/heap"
	ltheap "github.com/azd1997/Leetcode-training/ltcontainer/heap"
	"sort"
)

// 有序矩阵第K小元素

// 这道题是有序矩阵查找元素的变体
// 现在首先利用这种有序矩阵的特性

// 暴力一点可以用nlogn的二分
// 更机巧一些可以从左下角开始往上搜寻

// 回到这题，
// 纯暴力解法：将矩阵元素倒到一维数组进行快速排序（升序），
// 再用count标记从最大开始数直到count=k，就得到了第k小元素
// 变体（差不多也是暴力）：遍历矩阵元素，丢到大小为k的大顶堆中，
// 如果元素大于堆顶直接丢弃，小于则加入
// 第三种思路则是利用好题目矩阵的特性，尽可能不使用额外空间进行寻找

// 1. 纯暴力 O(n2lgn)/O(n2)
func kthSmallest1(matrix [][]int, k int) int {
	// 题目给出k有效，不用测边界条件
	n := len(matrix)

	// 倒入一维数组
	arr := make([]int, n*n)
	for i:=0; i<n*n; i++ {
		arr[i] = matrix[i/n][i%n]
	}

	// 数组排序 O(n2*lgn2) = O(n2lgn)
	sort.Ints(arr)

	// 数第K小数
	// count := 1
	// for i:=1; i<len(arr); i++ {
	// 	if arr[i]!=arr[i-1] {
	// 		count++
	// 		if count==k {return arr[i]}
	// 	}
	// }


	return arr[k-1]
}


// 2. 使用最大堆(优先队列)
// O(n2 * lnk)/O(k)
func kthSmallest(matrix [][]int, k int) int {

	// 准备好堆
	data := make([]int, 0, k+1)	// 为啥是k+1? 因为打算存k个元素，当再推一个使长度加一就把堆顶pop
	maxheap := ltheap.NewIntHeap(&data, func(i, j int) bool {
		return data[i]<data[j]
	})
	heap.Init(maxheap)

	// 将数据压入堆中，若堆长度超K则弹出堆顶最大值
	n := len(matrix)
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			heap.Push(maxheap, matrix[i][j])
			if maxheap.Len()==k+1 {heap.Pop(maxheap)}
		}
	}
	return maxheap.Seek()
}

// 3. 二分查找
// O(nlgnlgX) / O(lgnlgX)	  X为矩阵最大值与最小值之差
func kthSmallest3(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]	// 左上和右下，分别是最小值和最大值

	for left < right {

		// 统计整个矩阵中 < mid 的数的数量
		mid := left + (right - left) / 2
		count := 0
		for i:=0; i<n; i++ {
			count += findNotLargers(matrix[i], 0, n-1, mid)
		}

		// 根据该数量与k的比较调整范围，改变mid，重新尝试
		if count < k {
			left = mid + 1
		} else {right = mid}
	}

	// 最后count会等于k，left=right，返回此时的left
	return left
}

// 利用二分搜索在arr(升序排列)中寻找不大于target的元素总数， O(lgn)
func findNotLargers(arr []int, l, r, target int) int {
	if l>r {return 0}
	mid := (l+r)/2
	if arr[mid]<=target {
		return (mid-l+1) + findNotLargers(arr, mid+1, r, target)
	} else {
		return findNotLargers(arr, l, mid-1, target)
	}
}



// 4. 对解法3二分进行优化 // 二分查找 + 利用矩阵特性 O(nlgX)
// 二分本质上没有缩减搜索空间
// 前面说过可以利用自左下角一路向右上方移动来达到搜索目的
// 同样可以改造用来寻找整个矩阵中小于等于target的数目
func kthSmallest4(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]	// 左上和右下，分别是最小值和最大值

	for left < right {

		// 统计整个矩阵中 < mid 的数的数量
		mid := left + (right - left) / 2

		count := findNotLargersInMatrix(matrix, n, mid)
		// 根据该数量与k的比较调整范围，改变mid，重新尝试
		if count < k {
			left = mid + 1
		} else {right = mid}
	}

	// 最后count会等于k，left=right，返回此时的left
	return left
}

// 统计矩阵中 <= target 的元素总数。 O(n)
func findNotLargersInMatrix(matrix [][]int, n, target int) int {
	count := 0
	i, j := n-1, 0		// i,j为元素坐标，从左下角开始
	for i>=0 && j<n {	// 不超出矩阵范围
		if matrix[i][j] <= target {
			count += i+1	// 当前元素所在列的上部包括当前元素，总共i+1个
			j++		// 当前元素右移
		} else {
			i--		// 当前坐标上移
		}
	}
	return count
}
