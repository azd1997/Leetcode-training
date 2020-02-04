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

// 1. 纯暴力 O(nn)/O(nn)
func kthSmallest1(matrix [][]int, k int) int {
	// 题目给出k有效，不用测边界条件
	n := len(matrix)

	// 倒入一维数组
	arr := make([]int, n*n)
	for i:=0; i<n*n; i++ {
		arr[i] = matrix[i/n][i%n]
	}

	// 数组排序
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


// 1. 使用最大堆(优先队列)
func kthSmallest(matrix [][]int, k int) int {

	// 准备好堆
	data := make([]int, k+1)	// 为啥是k+1? 因为打算存k个元素，当再推一个使长度加一就把堆顶pop
	maxheap := ltheap.NewIntHeap(data, func(i, j int) bool {
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