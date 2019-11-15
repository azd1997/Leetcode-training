package lt102



// 二叉树的层次遍历



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

// 1. BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	// BFS
	res := make([][]int, 0)
	if root == nil  {return res}
	// BFS start
	queue := []*TreeNode{root}
	var level []int


	// 层级遍历，那么就是要第一次访问的时候就加到返回数组，并标记访问过了.
	// 并且要想办法一次性把一层的全取出来加到返回数组。不然不知道cur对应返回数组的哪一行
	// 那就是用数组收集每一层了。这里可以直接用queue的这个数组
	// 而且也不需要visited，因为不需要回溯（DFS才用到回溯思想）
	// 当然这样的写法并不算标准的BFS，但没关系，实现目的就好
	for len(queue) != 0 {

		l := len(queue)     //记录一下这一层的节点数

		// 顺序迭代该层
		level = []int{}
		for _, node := range queue {        // 不为空的才会加到queue，所以这里不用检查node为不为空
			level = append(level, node.Val)
			if node.Left != nil {queue = append(queue, node.Left)}
			if node.Right != nil {queue = append(queue, node.Right)}
		}
		res = append(res, level)
		queue = queue[l:]   // 把当前层的节点全部丢掉
	}

	return res
}




