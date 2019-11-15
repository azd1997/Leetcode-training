package lt104



// 二叉树的最大深度

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}




/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// DFS 递归，自顶向下

	if root == nil {return 0}
	return nextLayer(root, 1)       // 1为初始深度
}

func nextLayer(node *TreeNode, depth int) int {
	// 当前节点已判定不为nil
	d1, d2 := depth, depth
	if node.Left != nil {
		d1 =  nextLayer(node.Left, d1+1)
	}
	if node.Right != nil {
		d2 = nextLayer(node.Right, d2+1)
	}

	if d1 > d2 {
		return d1
	} else {
		return d2
	}
}
