package lcof68i

// 二叉搜索树的最近公共祖先

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思考：
// 1. 首先，“二叉树的最近公共祖先”的解法肯定能解这道，但是没有利用二叉搜索树的性质
// 2. 利用二叉搜索树的性质。中序遍历的话，lca的值一定落在[p.Val, q.Val]之间
// 也就是说，有了这样的有序的性质，可以进行一定的搜索区间的缩减

// 1. 和二叉树的LCA一样的处理
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

// 2. 利用搜索树的数值排序性质
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 缩减了一些搜索空间
	rv, pv, qv := root.Val, p.Val, q.Val
	if rv > pv && rv > qv { // 比两目标都大，去左子树找
		return lowestCommonAncestor2(root.Left, p, q)
	} else if rv < pv && rv < qv {
		return lowestCommonAncestor2(root.Right, p, q)
	} else {
		return root // 找到LCA
	}
}
