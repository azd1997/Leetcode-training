package lcof55_I

// 二叉树的深度

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 求二叉树深度，DFS或BFS都行，这里使用BFS(层次遍历)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) != 0 {
		depth++
		newQ := make([]*TreeNode, 0)

		for _, cur := range queue {
			// cur不需要处理
			// 只检查其孩子是否为空
			if cur.Left != nil {
				newQ = append(newQ, cur.Left)
			}
			if cur.Right != nil {
				newQ = append(newQ, cur.Right)
			}
		}

		queue = newQ
	}
	return depth
}
