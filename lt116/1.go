package lt116

// 填充每个节点的下一个右侧节点指针

// Definition for a node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 明显是层序遍历，使用基于队列的迭代层序遍历实现
// 但是题目要求O(1)的额外空间

// 1. 使用队列类的辅助结构进行层序遍历 O(N)/O(N)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 首先设置root的next
	root.Next = nil
	queue := []*Node{root}
	var tmpQ []*Node
	num := 0
	for len(queue) != 0 {
		num = len(queue)
		tmpQ = make([]*Node, 0, 2*num) // 下层有当前层结点数的2倍
		// 考虑结点在队列中是树左边的结点则在队列数组的左边 的这样的存取顺序
		for i := 0; i < num; i++ {
			// 设置Next
			if i < num-1 {
				queue[i].Next = queue[i+1]
			} else {
				queue[i].Next = nil
			}
			// 添加下一层。注意顺序
			if queue[i].Left != nil {
				tmpQ = append(tmpQ, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmpQ = append(tmpQ, queue[i].Right)
			}
		}
		// 更新queue
		queue = tmpQ
	}
	return root
}

// 2. 不使用额外空间的解法
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/tian-chong-mei-ge-jie-dian-de-xia-yi-ge-you-ce-j-3/
// 有两种类型的Next需要设置，一种是同一父节点，另一种则是相邻父节点
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 游标
	leftmost := root

	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			// 第1类Next ： 相同父节点
			head.Left.Next = head.Right
			// 第2类Next : 相邻父节点
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			// head在同一层上后移
			head = head.Next
		}
		// leftmost移到下一层
		leftmost = leftmost.Left
	}
	return root
}
