package lt993

// 二叉树的堂兄弟结点

// 二叉树的节点数介于 2 到 100 之间。
// 每个节点的值都是唯一的、范围为 1 到 100 的整数。

// 其实就是要检查两项：深度是否相同、不能在一个父节点下

// 还是基本的遍历，这里选择后序遍历

// 遍历一遍树，为每个节点加上深度信息

// 这样深度优先遍历，自然也是可以的，但是需要的辅助变量较多，容易写错

// 层序遍历（广度优先）很适合这道题
//

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	depth, xd, yd := 0, 0, 0
	for len(queue) != 0 {
		newQ := make([]*TreeNode, 0)

		for _, cur := range queue {
			// 检查目标是否在该节点的两个孩子上
			if cur.Left != nil && cur.Right != nil &&
				((cur.Left.Val == x && cur.Right.Val == y) ||
					(cur.Left.Val == y && cur.Right.Val == x)) {
				return false
			}

			// 检查cur是否是目标之一
			if cur.Val == x {
				xd = depth
			}
			if cur.Val == y {
				yd = depth
			}

			// 检查cur的孩子，入队
			if cur.Left != nil {
				newQ = append(newQ, cur.Left)
			}
			if cur.Right != nil {
				newQ = append(newQ, cur.Right)
			}
		}

		// 更新队列
		queue = newQ
		// 最后将深度加1
		depth++
	}
	return xd == yd
}
