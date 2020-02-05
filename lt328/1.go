package lt328

// 奇偶链表


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

func oddEvenList(head *ListNode) *ListNode {
	// 边界条件
	if head == nil {return head}
	if head.Next == nil {return head}
	if head.Next.Next == nil {return head}

	odd, even := &ListNode{}, &ListNode{}
	oddCur, evenCur := odd, even
	cur := head
	isOdd := true
	// 迭代到倒数第三个节点，迭代结束后cur为倒数第二个节点。isOdd标记倒数第二个节点的奇偶
	for cur.Next.Next != nil {
		tmp := cur.Next

		if isOdd {
			oddCur.Next = cur
			oddCur = oddCur.Next
			isOdd = false
		} else {
			evenCur.Next = cur
			evenCur = evenCur.Next
			isOdd = true
		}

		cur = tmp
	}
	// 现在cur，isOdd都是最后一个节点的

	// 对于倒数第二个，如果该节点是odd链表的那么需要将其节点Next指向even
	// 如果是even的，需要将其Next置为nil，再将odd的末尾的Next指向even
	if isOdd {	// 倒数第二个节点是odd尾结点
		tmp := cur.Next
		cur.Next = even.Next	// 指向even的第一个数据节点
		oddCur.Next = cur
		evenCur.Next = tmp
	} else {	// 倒数第二个节点是even尾结点
		tmp := cur.Next
		cur.Next = nil
		evenCur.Next = cur
		tmp.Next = even.Next	// 指向even的第一个数据节点
		oddCur.Next = tmp
	}
	return odd.Next
}


// 看到一种更叼的写法：
// 这代码好太多

func oddEvenList2(head *ListNode) *ListNode {
	// 边界条件
	if head == nil {return head}
	if head.Next == nil {return head}

	odd, even := head, head.Next	// odd,even双指针
	evenHead := even

	// 这个迭代截止条件用来：确保所有节点都被分别加入对应的odd或even
	for even!=nil && even.Next!=nil {
		// 更新odd
		odd.Next = even.Next
		odd = odd.Next

		// 更新even
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}