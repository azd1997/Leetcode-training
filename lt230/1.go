package lt230

// 二叉搜索树中第K小的元素

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


// 思路：
// 1. 遍历树(无论是前中后还是层序遍历)，装入数组，排序。 O(nlogn)/O(n)
// 2. 遍历树，将节点装入大小为k的最大堆中(也可以是其他类型和尺寸的堆，看具体做法)，最后堆顶为第k小  O(nlgk)/O(k)
// 3. 堆本身就是基于二叉搜索树的，因此直接对该二叉搜索树进行堆化。


// 参考官方题解
// 遍历树：DFS(前中后序遍历)、BFS(层序遍历)
// 题目要求的是第K小
// 那么利用二叉搜索树的定义和特性(node.Left < node < node.Right)
// 直到其中序遍历是升序的，那么只需要进行中序遍历到第k个节点就好了



// 1. 递归DFS实现中序遍历 O(n)/O(n) 遍历了整棵树
func kthSmallest(root *TreeNode, k int) int {
	arr := inorder(root, new([]int))
	return (*arr)[k-1]	// 这就是第k小元素
}

func inorder(root *TreeNode, arr *[]int) *[]int {
	if root==nil {return arr}	// 叶子节点.left/right

	// 中序遍历
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)

	return arr
}

// 2. 迭代(DFS)实现中序遍历
// 前面递归解法遍历了整棵树，但事实上我们只需要遍历前K个节点就好
// 使用辅助栈进行迭代，可以使我们控制找到答案后停止
// 时间 O(H+k)/O(H+k) H为树高度，当树平衡，H=logN, 树不平衡时，所有节点都在左子树，H=N
func kthSmallest2(root *TreeNode, k int) int {
	// 1. 边界处理(被迭代过程处理了)

	// 2. 辅助栈
	stack := make([]*TreeNode, 1)

	for {
		// 从当前节点出发，一路到该节点左支路所有节点压入栈中，到达最小值所在节点。
		// 结束该for root!=nil {} 时， root = nil， 要把栈顶取出，栈顶为最小值节点
		for root != nil {
			stack = append(stack, root)		// 先中后左
			root = root.Left
		}
		// 取出最小值节点，也就是第1小节点
		root = stack[len(stack)-1]; stack = stack[:len(stack)-1]
		k--
		if k==0 {return root.Val}	// 到达第k小节点，返回

		// 还没到第k小节点，需要遍历右节点(压栈过程中是中左中左的顺序压入)
		// 所以栈顶出栈时都要再访问其右节点，然后又会以该右节点作为根节点重新继续循环
		root = root.Right
	}
}


// 3. 其实递归解法也可以控制在第K小节点停止

func kthSmallest3(root *TreeNode, k int) int {
	// 不喜欢指针的话，可以在函数外定义一个变量 count，使两个函数都能访问
	// 也可以把inorder3定义到当前函数内部
	ans := 0
	inorder3(root, &k, &ans)
	return ans	// 这就是第k小元素
}

func inorder3(root *TreeNode, k, ans *int) {
	if root==nil {return}	// 叶子节点.left/right

	// 中序遍历
	inorder3(root.Left, k, ans)
	*k--	// 中间节点
	if *k==0 {
		*ans = root.Val
		return		// 找到目标则返回
	}
	inorder3(root.Right, k, ans)
}