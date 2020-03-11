package lcof55_II

// 平衡二叉树

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return postorder(root) != -1
}

// 返回子树的最大深度
func postorder(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := postorder(root.Left)
	r := postorder(root.Right)
	if l == -1 || r == -1 {
		return -1
	}

	if l-r > 1 || l-r < -1 {
		return -1
	} // -1表示不平衡
	if l > r {
		return l + 1
	}
	return r + 1
}
