package lt965

// 单值二叉树

// 这只是简单的遍历就行了...而且降不降重完全无所谓

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return preorder(root.Left, root.Val) && preorder(root.Right, root.Val)
}

func preorder(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}

	if root.Val != target {
		return false
	}

	return preorder(root.Left, target) && preorder(root.Right, target)
}
