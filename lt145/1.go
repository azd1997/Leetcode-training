package lt145



// 二叉树的后序遍历



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 1. 递归实现


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal1(root *TreeNode) []int {
	// 递归
	res := new([]int)
	if root == nil {return *res}

	postorder(root, res)

	return *res
}

// 后序遍历：左右中
func postorder(node *TreeNode, res *[]int) {
	if node.Left != nil {postorder(node.Left, res)}
	if node.Right != nil {postorder(node.Right, res)}
	*res = append(*res, node.Val)
}


// 2.DFS栈实现
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal2(root *TreeNode) []int {
	// 迭代 DFS
	res := make([]int, 0)
	if root == nil {return res}

	// DFS
	stack := []*TreeNode{root}
	visited := make(map[*TreeNode]bool)
	var cur *TreeNode
	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[cur] {
			res = append(res, cur.Val)
		}
		if !visited[cur] {
			// 后序遍历压栈顺序： 中右左
			stack = append(stack, cur)
			visited[cur] = true
			if cur.Right != nil {stack = append(stack, cur.Right)}
			if cur.Left != nil {stack = append(stack, cur.Left)}
		}
	}


	return res
}


