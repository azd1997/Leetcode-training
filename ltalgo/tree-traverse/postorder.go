package traverse

// 后序遍历

// PostOrderRecurse 将树按后序遍历输出至数组(递归版本)
func PostOrderRecurse(root *TreeNode) []int {
	arr := make([]int, 0)
	postorder(root, &arr)
	return arr
}

func postorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	postorder(root.Left, arr)
	postorder(root.Right, arr)
	*arr = append(*arr, root.Val)
}

// 后序遍历迭代版本与前序中序不同，前序中序都是先处理左子树再处理右子树
// 而后序遍历需要先处理完左右子树再处理当前。
// 因此需要一个标记 lastVisit
// 1. 如果 lastVisit = 当前节点的右子树，说明左右子树都处理完了
// 并且把lastVisit设置为当前结点，将当前节点cur（也是个游标）置为空
// 下一轮就可以访问栈顶元素
// 2. 否则接着考虑右子树， node = node.Right。 而且注意node暂时不能从栈中弹出

// PostOrderIterate 将树按后序遍历输出至数组(迭代版本)
func PostOrderIterate(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)

	stack := make([]*TreeNode, 0)
	cur := root       // 游标结点
	lastVisit := root // 游标结点，用来标记左右子树是否都访问完

	for len(stack) != 0 || cur != nil { // cur != nil 其实就是递归过程中到了底部

		//  只要当前节点不为空，始终将左子孙压入栈。
		// 由于是前序遍历，因此需要在左子节点压栈之前处理当前结点（加入到arr数组）
		for cur != nil {
			// 为了之后能够找到该节点的右子树，需要暂存该结点
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 查看当前栈顶元素 peek操作
		cur = stack[len(stack)-1]

		// 如果右子树也为空，或者右子树已经访问完，那么就直接输出当前结点
		if cur.Right == nil || lastVisit == cur.Right {
			// 处理当前结点
			arr = append(arr, cur.Val)
			// 将当前节点出栈，lastVisit移动到cur， cur 置空
			stack = stack[:len(stack)-1]
			lastVisit = cur
			cur = nil
		} else {
			// 否则的话，继续遍历右子树
			cur = cur.Right
		}
	}

	return arr
}
