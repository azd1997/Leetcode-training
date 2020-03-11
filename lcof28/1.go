package lcof28

// 对称的二叉树

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//  就是两个方向的前序遍历同时进行
// 对root.Left子树进行"中左右“前序遍历； 对root.Right子树进行”中右左“前序遍历
//  后序也可以做到，只是前序方便一些

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return preorder(root.Left, root.Right)
}

func preorder(lnode, rnode *TreeNode) bool {
	if lnode == nil {
		return rnode == nil
	}
	if rnode == nil {
		return lnode == nil
	}

	return lnode.Val == rnode.Val && preorder(lnode.Left, rnode.Right) && preorder(lnode.Right, rnode.Left)
}

// 使用迭代解法可以方便提前结束；递归结点需要额外设置信号
