package lt148


// 排序链表

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


// 归并排序 非递归 自底向上迭代
// O(nlogn)/O(1)
func sortList(head *ListNode) *ListNode {
	var h, h1, h2, pre, res *ListNode
	h = head

	// 求链表长度
	n := 0
	for h!=nil {
		h = h.Next
		n++
	}

	//
	intv :=  1
	res = &ListNode{}
	res.Next = head
	for intv < n {
		pre = res
		h = res.Next
		for h!=nil {
			//
			i := intv
			h1 = h
			for i>0 && h!=nil {
				h = h.Next
				i--
			}
			if i>0 {break}
			i = intv
			h2 = h
			for i>0 && h!=nil {
				h = h.Next
				i--
			}

			//
			c1, c2 := intv, intv - i
			for c1>0 && c2>0 {
				if h1.Val < h2.Val {
					pre.Next = h1
					h1 = h1.Next
					c1--
				} else {
					pre.Next = h2
					h2 = h2.Next
					c2--
				}
				pre = pre.Next
			}
			if c1==0 {
				pre.Next = h2
			} else {pre.Next = h1}

			for c1>0 || c2>0 {
				pre = pre.Next
				c1--; c2--
			}
			pre.Next = h
		}
		intv *= 2
	}
	return res.Next
}
