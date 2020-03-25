package lt117

// 填充每个节点的下一个右侧节点指针II

// Definition for a node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 题目不再给 完美二叉树 ，而是 普通二叉树
// 其实解法和上一题几乎是一样的：
// 使用队列的层序遍历解法代码完全不用修改
// 常量空间解法则需要不停向右寻找非空结点作为Next
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
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/tian-chong-mei-ge-jie-dian-de-xia-yi-ge-you-ce-j-4/

// 全局游标
var leftmost, prev *Node

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftmost, prev = root, nil // 重置全局变量，避免提交时影响到其他测例
	cur := leftmost            // 游标

	for leftmost != nil {

		prev = nil
		cur = leftmost
		leftmost = nil

		for cur != nil {
			processchild(cur.Left)
			processchild(cur.Right)
			cur = cur.Next // 移动到本层的下一个节点
		}

	}
	return root
}

func processchild(child *Node) {
	if child != nil {
		if prev != nil {
			prev.Next = child // 如果本层之前存在节点，也就是prev!=nil，则将prev.Next指向当前child
		} else {
			leftmost = child // 说明这个child是本层第一个非空结点，也就是leftmost
		}

		prev = child
	}
}
