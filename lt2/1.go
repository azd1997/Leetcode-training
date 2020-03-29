package lt2

// 两数相加

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路很简单，就是每位相加之后更新一个进位carry传到下一位的相加

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	output := &ListNode{} // 虚拟头结点
	cur := output         // 游标
	p1, p2 := l1, l2      // 游标
	carry := 0            // 进位
	sum := 0
	for p1 != nil && p2 != nil {
		sum = p1.Val + p2.Val + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
		p1, p2 = p1.Next, p2.Next
	}
	// 现在有三种情况：
	// 1. p1,p2同时为nil，要检查carry
	if p1 == nil && p2 == nil {
		if carry == 1 {
			cur.Next = &ListNode{Val: 1}
		}
	} else if p1 == nil { // 2. p2还有剩余
		cur.Next = p2
		cur.Next.Val += carry
	} else if p2 == nil { // 3. p1还有剩余
		cur.Next = p1
		cur.Next.Val += carry
	}

	return output.Next
}

// 上面解法是错的。因为在退出前面的迭代后，仍有可能发生进位，从而需要迭代下去
// 改成下面的写法：

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	output := &ListNode{} // 虚拟头结点
	cur := output         // 游标
	p1, p2 := l1, l2      // 游标
	sum := 0              // 每一轮两数对应位、进位之和。也作下一轮的进位
	for p1 != nil || p2 != nil {
		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		sum = sum / 10 // 进位
	}
	// 现在 p1,p2同时为nil，要检查sum
	if sum == 1 {
		cur.Next = &ListNode{Val: 1}
	}

	return output.Next
}
