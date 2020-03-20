package lt105

// 从前序与中序遍历序列构造二叉树

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 递归停止
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	} // 这两个数组始终等长

	// 当前区间的子树树根的值为 preorder[0]，再线性遍历找到树根在inorder中的位置
	// 据此可以将区间划分为左右子树的区间(先分inorder，然后可确定preorder)，并可递归下去

	root := &TreeNode{Val: preorder[0]}
	rootIdx := findIdx(inorder, root.Val)
	root.Left = buildTree(preorder[1:rootIdx+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	return root
}

// 线性遍历找目标值的下标
func findIdx(nums []int, target int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
