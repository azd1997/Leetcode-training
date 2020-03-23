package lt876

// 链表的中间节点

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思考：
// 1. 遍历并将元素倒到一个数组，然后通过数组访问中间元素
// 2. 先遍历一遍得到长度，再遍历至中点
// 3. 快慢指针，快指针到中点则慢指针指向中点

// 题中有说明，如果是偶数个节点，取右中点

// 1 -> 2 -> 3 -> 4 -> 5	// 中点3
// 1 -> 2 -> 3 -> 4 -> 5 -> 6	// 中点4
// 1 -> nil	// 中点1
// 1 -> 2 // 中点2

// 快慢指针
func middleNode(head *ListNode) *ListNode {
	// 题有说明链表非空，至少有一个节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
