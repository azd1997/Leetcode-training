package traverse

// 中序遍历

// InOrderRecurse 将树按中序遍历输出至数组(递归版本)
func InOrderRecurse(root *TreeNode) []int {
	arr := make([]int, 0)
	inorder(root, &arr)
	return arr
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

// InOrderIterate1 将树按中序遍历输出至数组(迭代版本)
// func InOrderIterate1(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	arr := make([]int, 0)

// 	stack := []*TreeNode{root}

// 	for len(stack) != 0 {
// 		// 先不出栈
// 		cur := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		// 处理当前
// 		arr = append(arr, cur.Val)

// 		// 压右
// 		if cur.Right != nil {
// 			stack = append(stack, cur.Right)
// 		}

// 		// 压左
// 		if cur.Left != nil {
// 			stack = append(stack, cur.Left)
// 		}

// 	}

// 	return arr
// }

// 迭代版本1中其实是最直接简单的迭代前序遍历
// 但是为了前序中序后序统一遍历框架，将迭代改写成如下：
// 使用一个游标结点 cur

// InOrderIterate2 将树按中序遍历输出至数组(迭代版本)
func InOrderIterate2(root *TreeNode) []int {
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
			// 为了之后能够找到该节点的右子树，需要暂存该结点
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 左边到底之后，开始考虑右子树
		// 如果栈已空，就不必再考虑
		//  弹出栈顶元素，将游标结点更新为栈顶元素的右结点
		if len(stack) != 0 {
			// 左出栈
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 处理出栈的结点（左->中）
			// 永远先考虑左子树，直到左子树为空，才访问根节点
			arr = append(arr, cur.Val)
			// 再将右入栈
			cur = cur.Right
		}
	}

	return arr
}

// 中序遍历迭代版本在结点出栈时处理
