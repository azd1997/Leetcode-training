package lt142

// 环形链表II

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}


// 1. 先用哈希集合辅助来作
func detectCycle(head *ListNode) *ListNode {
	node := head
	visited := make(map[*ListNode]bool)
	for node!=nil {
		if visited[node] {return node}
		visited[node] = true
		node = node.Next
	}
	return nil
}


// 可能会想到和 环形链表 一样使用快慢指针， 但是快慢指针跑只能检测是否成环，
// 快慢指针相遇的地方不一定在环的入口