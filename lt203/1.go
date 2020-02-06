package lt203

// 移除链表元素
// 删除链表中 等于给定值val 的 所有 节点

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


// 这种题没有任何技巧了吧？ 遍历，然后删
// 只是链表删除节点有两种做法：
// a->b->c 假如说删除 b
// 可以将a.Next指向c
// 也可以将b.val = c.val, b.Next = c.Next
// 这里采用第二种
func removeElements1(head *ListNode, val int) *ListNode {

	// 由于下面 if newhead==nil {return nil} 可以处理，所以这里不需要了
	//if head==nil {return nil}

	// 这里要注意 head.Val = val 的情况
	// 需要定位到新head
	newhead := head
	for newhead!=nil && newhead.Val==val {
		newhead = newhead.Next
	}
	// 现在newhead.Val != val 或者 newhead==nil
	if newhead==nil {return nil}

	node := newhead
	pre := newhead
	for node != nil {
		if node.Val == val {
			if node.Next!=nil {
				node.Val = node.Next.Val
				node.Next = node.Next.Next
			} else {pre.Next = nil; break}// 最后一个节点=val，则将pre指向nil，直接跳出循环

			continue	// 这种情况下不能更新pre和node
		}
		pre = node
		node = node.Next
	}

	return newhead
}

// 这道题陷阱在于:
// 头部节点需要删除的情形
// 尾部节点需要删除的情形
// 边界条件(head==nil)的处理


// 显然，从上面的做法过程中，足足改了五次才改对，很容易出错
//
// 当需要删除头部节点时，通常使用 哨兵节点 sentinel
// 具体参考官方题解
func removeElements2(head *ListNode, val int) *ListNode {
	// 哨兵节点，作为伪头
	sentinel := &ListNode{Next:head}

	// pre, cur 双指针
	pre, cur := sentinel, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}
	return sentinel.Next
}


// 也可以改写成递归实现
func removeElements3(head *ListNode, val int) *ListNode {
	// 哨兵节点，作为伪头
	sentinel := &ListNode{Next:head}

	// pre, cur 双指针
	pre, cur := sentinel, head

	// 递归
	helper(pre, cur, val)

	return sentinel.Next
}

func helper(pre, cur *ListNode, val int) {
	if cur==nil {return}
	if cur.Val==val {
		pre.Next = cur.Next
		helper(pre, cur.Next, val)
	} else {
		helper(cur, cur.Next, val)
	}
}


// 更简洁的递归写法
func removeElements4(head *ListNode, val int) *ListNode {
	// 边界条件
	if head==nil {return nil}
	// 递去
	head.Next = removeElements4(head.Next, val)
	// 递归
	if head.Val==val {
		return head.Next
	} else {return head}
}