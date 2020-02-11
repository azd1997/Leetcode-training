package lt124

import "math"

// 二叉树中的最大路径和

// 这个路径和，表示的是从节点A到节点B的路径上所有节点值之和
// 然后要在整个树里面找最大的路径和

//// 而且这个路径是可以路过父级节点的，也就是说这个路径A->Parent->B
// 可以看做 Parent->A->Parent->B 的深度优先遍历


// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 首先明确找所谓的最大路径和，被转变为了求最大的左右两边单条路径和
// 对于节点 root而言，要求的是经过它的最大路径，意味着左边要找一条支路
// 使得这条支路的和最大，要注意的是，是支路，不是子树； 右边同样
// 然后左右支路路径和相加再加上当前节点值，就是经过当前节点的最大路径和

// 现在明确一下，从根开始深度优先遍历
// 到叶子节点时记录叶子节点的最大路径和(本身)
// 回溯阶段，父节点从左右中选一条大的支路，得到父节点所对应的最大支路路径和
// 同时计算父节点的两支路路径和+本身，更新最大路径和
// 不停回溯

// 1. 按照上面的思路，先直接写，不考虑优化.。返回值为最大路径和
// 假设有
func maxPathSum(root *TreeNode) int {
	if root == nil {return 0}

	// 计算所有节点的最大支路路径和
	maxBranchSums := make(map[*TreeNode]int)
	maxBranchSum(root, maxBranchSums)

	// DFS
	maxsum := math.MinInt32
	stack := []*TreeNode{root}
	visited := make(map[*TreeNode]bool)
	for len(stack)!=0 {
		node := stack[len(stack)-1]; stack = stack[:len(stack)-1]	// 取栈顶
		if !visited[node] {
			left, right := 0, 0
			if node.Left!=nil {
				left = maxBranchSums[node.Left]
				stack = append(stack, node.Left)
			}
			if node.Right!=nil {
				right = maxBranchSums[node.Right]
				stack = append(stack, node.Right)
			}
			maxsum = max(maxsum, node.Val + left + right)
			visited[node] = true
		}
	}

	return maxsum
}

// 求root及其子孙节点节点的最大支路路径和（（就是只到cur，而不越过cur到cur的另一侧））
// 返回所有节点对应的最大支路路径和，存入maxBranchSums表
func maxBranchSum(root *TreeNode, maxBranchSums map[*TreeNode]int) int {
	if root==nil {return 0}
	maxBranchSums[root] = max(root.Val, root.Val + max(maxBranchSum(root.Left, maxBranchSums), maxBranchSum(root.Right, maxBranchSums)))
	return maxBranchSums[root]
}


// TODO: 上面的代码有问题


func max(a,b int) int {if a>=b {return a} else {return b}}



// 2. 递归。 参考官方题解 O(n)/O(logn)
func maxPathSum2(root *TreeNode) int {
	maxsum := math.MinInt32
	maxGain(root, &maxsum)
	return maxsum
}

func maxGain(node *TreeNode, maxsum *int) int {
	if node==nil {return 0}

	// node左右子树的最大贡献(左右支路最大数值和)
	leftGain := max(maxGain(node.Left, maxsum), 0)	// 0表示不向下分支，只取本身
	rightGain := max(maxGain(node.Right, maxsum), 0)

	// 经过当前node且以当前node为路径最高点的路径的和
	curPathSum := node.Val + leftGain + rightGain

	// 更新maxsum
	*maxsum = max(*maxsum, curPathSum)

	// 将当前的节点最大分支和返回
	return node.Val + max(leftGain, rightGain)
}


