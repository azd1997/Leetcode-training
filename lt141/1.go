package lt141

// 环形链表

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/linked-list-cycle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//进阶：
//
//你能用 O(1)（即，常量）内存解决此问题吗？

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}


// 思考：
// 	1. 非常直观的一个做法是：遍历过程中记录节点是否出现过，出现过则循环了。用哈希集和优化之后 O(n)/O(n)
//  2. 如果要O(1)空间的话，由于成环者，使用快慢指针，最终一定会相遇




// 1. 哈希集合解法
//17/17 cases passed (12 ms)
//Your runtime beats 15.9 % of golang submissions
//Your memory usage beats 14.31 % of golang submissions (5.2 MB)
func hasCycle(head *ListNode) bool {
	p1 := head
	set := make(map[*ListNode]bool)
	for p1 != nil {
		if set[p1] {
			return true
		} else {
			set[p1] = true
		}
		p1 = p1.Next
	}
	return false // p1 == nil
}


// 2. 快慢指针
//17/17 cases passed (8 ms)
//Your runtime beats 83.53 % of golang submissions
//Your memory usage beats 95.89 % of golang submissions (3.8 MB)
func hasCycle2(head *ListNode) bool {
	low, fast := head, head
	for low != nil && fast != nil && fast.Next != nil {		// 不可以只用fast!=nil或者low!=nil，可能会使fast空指针引用
		low, fast = low.Next, fast.Next.Next
		if low==fast {return true}
	}
	return false
}