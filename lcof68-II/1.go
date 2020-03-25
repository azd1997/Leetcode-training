package lcof68ii

// 二叉树的最近公共祖先

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最近公共祖先必然是 自下而上 第一个子树中同时包含两个target 的结点
// 考虑一个辅助的递归函数，每次当其子树中搜索到其中一个target时，就往上返回true
// 当自下而上第一个 Left返回了true，right也返回了true的时候，就找到了 最近公共祖先 LCA
// 这个返回的true/false可以用*TreeNode的nil与非nil来表示
// 这里考虑true时返回内含的某个target(p或q)，false返回nil

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} // 找到底都没找到target
	if root == p || root == q { // 这里这么写： 因为是后序遍历，所以在下面找到的target一定会先返回
		return root
	} // 找到target之一

	// 当前结点非空，则继续找当前节点的子树（后序遍历）
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	// 处理当前。
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// 2. 使用辅助函数的写法，啰嗦了一些

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	lca, _ := help(root, p, q)
	return lca
}

// 当找到LCA，则返回的结点不为空，而是LCA.
// 当LCA!=nil时，布尔值无意义
func help(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	} // 找到底都没找到target

	// 当前结点非空，则继续找当前节点的子树（后序遍历）
	count := 0 // count=2就是得到两个true
	// 左
	lca1, left := help(root.Left, p, q)
	if lca1 != nil {
		return lca1, true
	}
	if left == true {
		count++
	}
	// 右
	lca2, right := help(root.Right, p, q)
	if lca2 != nil {
		return lca2, true
	}
	if right == true {
		count++
	}
	// 中
	if root == p || root == q {
		count++
	}
	if count == 2 {
		return root, true
	}
	if count == 1 {
		return nil, true
	}
	return nil, false
}

// 上面这个解法是正确的，但是很明显可以看出，它在找到LCA时不能及时返回，只能不断回去重新到达根节点处才真正返回给调用者
