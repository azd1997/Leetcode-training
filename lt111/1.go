package lt111

// 二叉树的最大深度

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. 解法1，写起来很简洁，但是我们要求最小深度，它却会遍历整颗树
//41/41 cases passed (4 ms)
//Your runtime beats 96.13 % of golang submissions
//Your memory usage beats 62.63 % of golang submissions (5.3 MB)
func minDepth(root *TreeNode) int {
	if root == nil {return 0}

	// 注意，求最小深度时 只有当两个节点都非空，才能使用 min来比较
	// 而一方为空的时候，却应该使用max
	if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else {
		return max(minDepth(root.Left), minDepth(root.Right)) + 1
	}

}

func min(a, b int) int {
	if a<=b {return a}
	return b
}

func max(a, b int) int {
	if a>=b {return a}
	return b
}

// 最佳的搜索应当是广度优先遍历，在二叉树中称作层级遍历，借助辅助队列或者数组，
// 在任一层级发现有节点已经是叶节点则立即返回

// 2. 广度优先搜索 迭代
func minDepth2(root *TreeNode) int {
	if root == nil {return 0}
	// 初始将根节点压入queue中。实际上只是作数组用，而并非按队列操作
	// queue中成对存放节点的左右子节点
	queue := [][2]*TreeNode{[2]*TreeNode{root, nil}}
	for len(queue) != 0 {	// 迭代至queue为空时，所有节点都遍历完

		// 每次取queue中两个值，它们互为亲兄弟关系
		// 若发现 一对亲兄弟都不为空，则可以继续遍历
		// 在一轮queue中，如果发现某个节点自己不为nil且下边左右子节点都为空，则停止整个循环，返回当前累加的最小深度
		for _, pair := range queue {

		}
	}
}