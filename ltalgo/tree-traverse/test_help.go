package traverse

// GenBinaryTree 生成二叉树，返回树根
func GenBinaryTree(data []int) *TreeNode {
	// 根据数组可以生成多种多样的树
	// 方便的做法是采用二分法，来构建树
	// 每次取当前结点作为中间结点
	// 左区间为左子树，右区间为右子树

	n := len(data)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: data[0]}
	}

	mid := n / 2
	root := &TreeNode{Val: data[mid]}
	root.Left = GenBinaryTree(data[:mid])
	root.Right = GenBinaryTree(data[mid+1:])
	return root
}
