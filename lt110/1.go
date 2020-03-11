package lt110

// 平衡二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 后序遍历就可以解决
func isBalanced(root *TreeNode) bool {
	return postorder(root) != -1
}

// 后序遍历，如果当前子树平衡，则返回当前子树的总高度
func postorder(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := postorder(root.Left)
	rightH := postorder(root.Right)

	// 如果左右子树都不是平衡的话，直接返回-1
	if leftH == -1 || rightH == -1 {
		return -1
	}

	if rightH-leftH > 1 || rightH-leftH < -1 {
		return -1 // 表示不平衡
	}
	return max(leftH, rightH) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
