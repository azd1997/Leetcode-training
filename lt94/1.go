package lt94



// 二叉树的中序遍历



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
func inorderTraversal1(root *TreeNode) []int {
	// 递归实现
	res := new([]int)
	if root == nil {return *res}
	inorder(root, res)

	return *res
}

// 中序遍历，先遍历左节点，再访问中间节点，最后右
func inorder(node *TreeNode, res *[]int) {
	if node.Left != nil {inorder(node.Left, res)}
	*res = append(*res, node.Val) // 当前node已经判定过非空
	if node.Right != nil {inorder(node.Right, res)}
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
func inorderTraversal2(root *TreeNode) []int {
	// 迭代实现 DFS
	res := make([]int, 0)
	if root == nil {return res}

	// DFS
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visited := make(map[*TreeNode]bool)
	var cur *TreeNode
	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[cur] {
			res = append(res, cur.Val)      // 访问过一次后，回溯到的时候再添加到数组
		}
		if !visited[cur] {
			// 压栈顺序：右中左
			if cur.Right != nil {stack = append(stack, cur.Right)}
			stack = append(stack, cur)
			visited[cur] = true
			if cur.Left != nil {stack = append(stack, cur.Left)}
		}
	}

	return res
}




