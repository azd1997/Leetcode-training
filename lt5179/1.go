package lt5179

// 单周赛180 t3
// 将二叉搜索树变平衡

// 动态修复BST的平衡，尝试过，但写不出来。
// 像红黑树、2-3树之类的要维护平衡都是每一步都按规则来保证子树高度差在一定范围，然后可以通过旋转等一些操作来维护
// 但是这道题，给的BST你不知道子树的左右高度差有多大，这就很难搞，反正现在不会

// 另一个粗暴的思路是先中序遍历得到数组，再生成一颗平衡的二叉搜索树（二分区间构造树）

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//直接中序遍历，构造新树
func balanceBST(root *TreeNode) *TreeNode {
	// 1. 中序遍历
	vals := make([]int, 0)
	inorder(root, &vals)

	// 2.重新构造BST
	newRoot := help(vals)
	return newRoot
}

func inorder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)
}

func help(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	if len(vals) == 1 {
		return &TreeNode{Val: vals[0]}
	}

	l, r := 0, len(vals)-1
	mid := (l + r) / 2
	root := &TreeNode{Val: vals[mid]}
	root.Left = help(vals[l:mid])
	root.Right = help(vals[mid+1 : r+1])
	return root
}
