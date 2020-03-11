package lt226

// 翻转二叉树

// 后序遍历，处理当前结点时，先求左右子树翻转后的新根，再更新当前子树

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)
	return root
}
