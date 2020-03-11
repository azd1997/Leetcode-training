package traverse

// 前序遍历

// PreOrderRecurse 将树按前序遍历输出至数组(递归版本)
func PreOrderRecurse(root *TreeNode) []int {
	arr := make([]int, 0)
	preorder(root, &arr)
	return arr
}

func preorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)
	preorder(root.Left, arr)
	preorder(root.Right, arr)
}

// PreOrderIterate1 将树按前序遍历输出至数组(迭代版本)
func PreOrderIterate1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		// 出栈
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 处理当前
		arr = append(arr, cur.Val)

		// 压右
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		// 压左
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

	}

	return arr
}

// 迭代版本1中其实是最直接简单的迭代前序遍历
// 但是为了前序中序后序统一遍历框架，将迭代改写成如下：
// 使用一个游标结点 cur

// PreOrderIterate2 将树按前序遍历输出至数组(迭代版本)
func PreOrderIterate2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)

	stack := make([]*TreeNode, 0)
	cur := root // 游标结点

	for len(stack) != 0 || cur != nil { // cur != nil 其实就是递归过程中到了底部

		//  只要当前节点不为空，始终将左子孙压入栈。
		// 由于是前序遍历，因此需要在左子节点压栈之前处理当前结点（加入到arr数组）
		for cur != nil {
			arr = append(arr, cur.Val)
			// 为了之后能够找到该节点的右子树，需要暂存该结点
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 左边到底之后，开始考虑右子树
		// 如果栈已空，就不必再考虑
		//  弹出栈顶元素，将游标结点更新为栈顶元素的右结点
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}

	return arr
}
