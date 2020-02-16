package lcci1712

// 程序员面试金典 17.12

// BiNodes 将二叉搜索树转换为单链表

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// 解法：
// 中序遍历，用一个pre记录上一次处理的节点，处理当前节点node时，
// 将pre.right指向node，将node.left置空，pre再后移至node上，而后node进入右子树



// 1. 递归
func convertBiNode1(root *TreeNode) *TreeNode {
	dummyHead := &TreeNode{}
	inOrder(root, dummyHead)
	return dummyHead.Right
}

// 中序遍历 返回新的根节点
func inOrder(node, pre *TreeNode) *TreeNode {
	if node == nil {return pre}		// 返回上一个节点作为新根
	// 可以想象以下遍历到整棵树左下方时，那个叶子节点(最小值节点)被记录为pre，其再下就是nil
	// 因此将pre返回作为新根

	// 遍历左子树
	pre = inOrder(node.Left, pre)
	// 处理当前
	node.Left = nil
	pre.Right = node
	pre = node	// 更新pre(上移)
	// 遍历右子树
	pre = inOrder(node.Right, pre)

	return pre
}

// 2. 循环 + 栈
// 暂时有问题，
// 输入
//[4,2,5,1,3,null,6,0]
//输出
//[0,null,1,null,2,null,4,null,5,null,6]
//预期结果
//[0,null,1,null,2,null,3,null,4,null,5,null,6]
func convertBiNode2(root *TreeNode) *TreeNode {
	dummyHead := &TreeNode{}
	pre := dummyHead

	// 中序遍历
	node := root
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) != 0 {
		if node != nil {	// 一路向左到达最小节点处
			stack = append(stack, node)
			node = node.Left
		} else {
			node := stack[len(stack)-1]; stack = stack[:len(stack)-1]

			// 处理当前
			node.Left = nil
			pre.Right = node
			pre = node

			// 中序遍历，进入右子树
			node = node.Right
		}
	}

	return dummyHead.Right
}