package lt144



// 二叉树的前序遍历



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
func preorderTraversal1(root *TreeNode) []int {
	// 递归实现

	res := new([]int)
	preorder(root, res)

	return *res
}

// 注意函数参数是值传递，因此需要传切片指针
func preorder(node *TreeNode, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
		preorder(node.Left, res)
		preorder(node.Right, res)
	}
	return
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
func preorderTraversal2(root *TreeNode) []int {
	// 迭代实现。前序中序后序遍历都是使用DFS； 层级遍历使用BFS


	res := make([]int, 0)
	if root == nil {return res}        // 返回空数组
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visited := make(map[*TreeNode]bool)

	var cur *TreeNode
	for len(stack) != 0 {       // 只有非空的节点才加入栈
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[cur] {
			res = append(res, cur.Val)      // 前序遍历要先压根节点，再压右节点，再压左节点。
			visited[cur] = true             // 上一次提交忘加这句，等于visited白设，增添了一些重复访问
			if cur.Right != nil {stack = append(stack, cur.Right)}
			if cur.Left != nil {stack = append(stack, cur.Left)}
		}
	}

	return res
}

// 3.DFS栈实现
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal3(root *TreeNode) []int {
	// 迭代实现。前序中序后序遍历都是使用DFS； 层级遍历使用BFS


	res := make([]int, 0)
	if root == nil {return res}        // 返回空数组
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visited := make(map[*TreeNode]bool)

	var cur *TreeNode
	for len(stack) != 0 {       // 只有非空的节点才加入栈
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[cur] {
			res = append(res, cur.Val)      // 前序遍历回溯时再把值加入到返回数组（只有前序遍历才能够在回溯之前也就是刚访问的时候就把值加到返回数组）。 为了代码的通用与好理解，应该在写前中后序遍历的时候都在回溯时才把值加入到返回数组
		}
		if !visited[cur] {
			if cur.Right != nil {stack = append(stack, cur.Right)}
			if cur.Left != nil {stack = append(stack, cur.Left)}
			stack = append(stack, cur)  // 最后压栈中间结点
			visited[cur] = true
		}
	}

	return res
}
