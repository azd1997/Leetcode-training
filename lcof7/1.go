package lcof7


// 重建二叉树


// Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}



func buildTree(preorder []int, inorder []int) *TreeNode {

	n := len(inorder)

	// 中序遍历表
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	// 递归
	return buildRoot(m, preorder, inorder, 0, n-1, 0, n-1)
}

//  构建当前的子树的根节点
func buildRoot(m map[int]int, pre, in []int, l1, r1, l2, r2 int) *TreeNode {
	// 没有子树了，返回nil节点
	if l1>r1 || l2>r2 {return nil}
	// 否则找到当前子树的根节点（pre的左边界）
	root := &TreeNode{Val:pre[l1]}
	// 找到根节点在中序遍历数组中的下标
	rootInorderIdx := m[pre[l1]]

	// 接下来就是找到pre/in的接下来的left和right子树

	// 递归其左右子节点
	// 在中序遍历数组直接是将rootInorderIndex作为分割线，分成两个子区间
	// 前序遍历则需要将需要将区间置为[l1+1, l1+rootInorder-l2]
	root.Left = buildRoot(m, pre, in, l1 + 1, l1 + rootInorderIdx - l2, l2, rootInorderIdx-1)
	root.Right = buildRoot(m, pre, in, l1 + 1 + rootInorderIdx - l2,r1, rootInorderIdx+1, r2)
	return root
}

