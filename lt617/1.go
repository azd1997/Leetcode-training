package lt617

// 合并二叉树

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 对两棵树同时进行前序遍历即可

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	} // 处理了三种情况：都为空以及一方为空

	// 处理中间节点
	t := &TreeNode{Val: t1.Val + t2.Val}

	// 进入左右子树
	t.Left = mergeTrees(t1.Left, t2.Left)
	t.Right = mergeTrees(t1.Right, t2.Right)
	return t
}
