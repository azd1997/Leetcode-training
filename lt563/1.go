package lt563

// 二叉树的坡度

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 选择后序遍历(前序中序其实也都可以)，携带两类变量：坡度、子树结点之和

func findTilt(root *TreeNode) int {
	tilt := 0
	postorder(root, &tilt)
	return tilt
}

// 返回子树结点之和
func postorder(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	} // 直接返回

	l := postorder(root.Left, tilt)
	r := postorder(root.Right, tilt)

	// 计算当前
	tmp := l - r
	if tmp < 0 {
		tmp = -tmp
	}
	*tilt += tmp
	return l + r + root.Val
}
