package tree

import (
	"math"
)

// 线段树

// 基于数组存储
// 这里以整数、区间求和为例，

// 数组长度：以满二叉树形式，取一个较大值，足够存储满二叉树

// MaxLen 最大长度，保证tree能完全存储一颗满二叉树（其实再小些，完全二叉树也行）
// 通常 tree数组的长度初始化为 arr 数组长度的 4 倍
const MaxLen = 100

// SegmentTree 基于数组存储的线段树
type SegmentTree struct {
	arr  []int // 原数组
	tree []int // 线段树
}

// BuildTree 根据给定数组生成线段树
func BuildTree(arr []int) *SegmentTree {
	n := len(arr)
	if n == 0 {
		return nil
	}

	st := &SegmentTree{}
	st.arr = arr
	st.tree = make([]int, MaxLen)

	st.build(0, 0, n-1) // 从0这个结点开始（0是树顶）
	return st
}

// node 指当前节点的数组下标
func (st *SegmentTree) build(node, start, end int) {
	if start == end {
		st.tree[node] = st.arr[start] // 直接给当前节点赋值
		return
	}

	mid := (start + end) / 2
	leftnode, rightnode := 2*node+1, 2*node+2
	// 构建左孩子
	st.build(leftnode, start, mid)
	// 构建右孩子
	st.build(rightnode, mid+1, end)
	// 赋值当前结点
	st.tree[node] = st.tree[leftnode] + st.tree[rightnode]
}

// Query 查询
// 根据查询的区间上下界去寻找
func (st *SegmentTree) Query(l, r int) int {
	// 1. 检查区间有效性
	if l > r || l < 0 || r >= len(st.arr) {
		return math.MinInt32
	}

	return st.query(0, 0, len(st.arr)-1, l, r)
}

func (st *SegmentTree) query(node, start, end, l, r int) int {
	if l > end || r < start { // 没有重合
		return 0
	}
	if start >= l && end <= r { // 当前节点处在所求区间的内部，则直接返回当前节点的值就好
		return st.tree[node]
	}

	// 先获取左边部分的总和、右边部分的总和，再相加，返回
	mid := (start + end) / 2
	leftnode, rightnode := 2*node+1, 2*node+2
	sumLeft := st.query(leftnode, start, mid, l, r)
	sumRight := st.query(rightnode, mid+1, end, l, r)
	return sumLeft + sumRight
}

// Update 更新
// 根据idx寻找要修改的路径
func (st *SegmentTree) Update(idx, val int) {
	// 1. 修改arr
	if idx < 0 || idx >= len(st.arr) {
		return // idx无效
	}
	st.arr[idx] = val

	// 2. 更新tree
	st.update(0, 0, len(st.arr)-1, idx, val)
}

func (st *SegmentTree) update(node, start, end, idx, val int) {
	// 递归终止
	if start == end {
		st.tree[node] = val
		return
	}

	// 1. 先将下方的路径修改
	mid := (start + end) / 2
	leftnode, rightnode := 2*node+1, 2*node+2
	// 看idx是在node的左子树还是右子树
	if idx >= start && idx <= mid { // idx有效且在左子树
		st.update(leftnode, start, mid, idx, val)
	} else if idx <= end && idx > mid { // idx有效且在右子树
		st.update(rightnode, mid+1, end, idx, val)
	}
	// 2. 再修改当前节点
	st.tree[node] = st.tree[leftnode] + st.tree[rightnode]
}
