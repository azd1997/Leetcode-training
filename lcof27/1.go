package lcof27

//  二叉树的镜像

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 处理当前
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}
