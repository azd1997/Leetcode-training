package lt654

import (
	"fmt"
	"math"
)

// 最大二叉树

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 思考：
// 核心就是不断找最大值然后递归构建树
// 效率的差异体现在找最大值

// 低效的办法是，每次对区间进行线性遍历
// 高效的办法是线段树

// 1. 线性遍历求最大值 + 递归构建树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	return help(nums, 0, n-1)
}

func help(nums []int, l, r int) *TreeNode {
	// 递归终止条件
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}

	// 否则的话，构建当前结点及其左右孩子
	maxIdx := maxOfSlice(nums, l, r)
	cur := &TreeNode{Val: nums[maxIdx]}
	cur.Left = help(nums, l, maxIdx-1)
	cur.Right = help(nums, maxIdx+1, r)
	return cur
}

// 线性遍历求最大值，返回最大值的数组下标
func maxOfSlice(arr []int, l, r int) int {
	max := math.MinInt32
	maxIdx := -1
	for i := l; i <= r; i++ {
		if arr[i] > max {
			max = arr[i]
			maxIdx = i
		}
	}
	return maxIdx
}

// 2. 线段树求最大值 + 递归构建树
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	// 构建线段树
	tree := buildSegmentTree(nums)

	fmt.Println("构建线段树成功")

	return help2(nums, tree, 0, n-1)
}

func help2(nums []int, tree [][2]int, l, r int) *TreeNode {
	// 递归终止条件
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}

	// 否则的话，构建当前结点及其左右孩子
	max := querySegmentTree(nums, tree, 0, 0, len(nums)-1, l, r)
	cur := &TreeNode{Val: max[1]}
	cur.Left = help2(nums, tree, l, max[0]-1)
	cur.Right = help2(nums, tree, max[0]+1, r)
	return cur
}

// 线段树，非常适合对区间进行最大值最小值平均值等情况的快速计算。
// 这里直接以tree数组来表示。
// 这里的线段树存的是最大值的下标和最大值 [2]int
func buildSegmentTree(arr []int) [][2]int {
	n := len(arr)
	if n == 0 {
		return nil
	}

	tree := make([][2]int, 4*n)
	build(arr, tree, 0, 0, n-1)
	return tree
}

func build(arr []int, tree [][2]int, node, start, end int) {
	if start == end {
		tree[node] = [2]int{start, arr[start]}
		return
	}

	mid := (start + end) / 2
	leftnode, rightnode := 2*node+1, 2*node+2
	build(arr, tree, leftnode, start, mid)
	build(arr, tree, rightnode, mid+1, end)
	if tree[leftnode][1] > tree[rightnode][1] {
		tree[node] = tree[leftnode]
	} else {
		tree[node] = tree[rightnode]
	}
}

// 查询区间[l,r]，返回区间的最大值下标和最大值
func querySegmentTree(arr []int, tree [][2]int, node, start, end, l, r int) [2]int {
	fmt.Println(node, start, end)

	if l > end || r < start { // 当前区间在查询区间外
		return [2]int{-1, math.MinInt32}
	}
	if start >= l && end <= r { // 当前区间在查询区间内部
		return tree[node]
	}

	mid := (start + end) / 2
	leftnode, rightnode := 2*node+1, 2*node+2
	maxLeft := querySegmentTree(arr, tree, leftnode, start, mid, l, r)
	maxRight := querySegmentTree(arr, tree, rightnode, mid+1, end, l, r)
	if maxLeft[1] >= maxRight[1] {
		return maxLeft
	}
	return maxRight

}

////////////////////////////////////////

// 3. 单调栈或单调队列求最大值 + 递归构建树
// 单调队列或者单调栈在滑动窗口求最大值最小值这一类问题上也很好用，实现O(n)的总体时间复杂度
// 以后再写
