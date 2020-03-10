package lt543

// 二叉树的直径

// 给定一棵二叉树，你需要计算它的直径长度。
// 一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

//示例 :
// 给定二叉树

//  	 1
// 		/ \
// 	   2   3
// 	  / \
// 	 4   5
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

//注意：两结点之间的路径长度是以它们之间边的数目表示。

// 思考：
// 两点之间最长的距离其实必然穿过某个子树的根节点。
// 这个问题就可以变成：
// 遍历所有结点，找到左右子树深度之和最大的那个结点，
// 返回其左右子树相对（该结点）深度之和（假设该节点深度为0）

// 那么需要后序遍历
// 遍历每科子树时将左右子树的最大相对深度相加得到当前子树的“直径”
// 并将左右子树的最大相对深度的大者 + 1 返回

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxd := 0 // 最大直径
	postorder(root, &maxd)
	return maxd
}

// 返回当前子树（root）相对于其父节点的最大相对深度
func postorder(root *TreeNode, maxd *int) int {
	// 递归终止
	if root == nil {
		return 0
	}

	// 左子树
	leftDepth := postorder(root.Left, maxd)
	// 右子树
	rightDepth := postorder(root.Right, maxd)
	// 当前
	tmp := leftDepth + rightDepth // 当前子树的直径
	if tmp > *maxd {
		*maxd = tmp // 更新整棵树的直径
	}
	// 返回当前子树（root）相对于其父节点的最大相对深度
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
