package lt106

// 从后序与中序遍历序列构造二叉树

// 对于后序遍历，头结点位于postorder末尾
// 然后根据这个头结点值 遍历inorder得到头结点位置，进而分出左右子树

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 { // inorder和postorder长度是一样的
		return &TreeNode{Val: inorder[0]}
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIdx := getIdx(inorder, root.Val)
	root.Left = buildTree(
		inorder[:rootIdx],
		postorder[:rootIdx], // 这种地方就画图，会清晰许多
	)
	root.Right = buildTree(
		inorder[rootIdx+1:],
		postorder[rootIdx:len(postorder)-1],
	)
	return root
}

func getIdx(inorder []int, val int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			return i
		}
	}
	return -1
}
