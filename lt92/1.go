package lt92

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代， n-1下找到m位置节点 ，单独将m到n范围的链表反转，记录新链表的头与尾，这样可以看做是一趟扫描

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 先找到位置m的前一个以及位置m
	dummy := &ListNode{Next: head}
	left := dummy
	for m > 1 {
		left = left.Next
		m--
	}

	// 区间链表的头
	subHead := left.Next

	fmt.Println("subHead: ", subHead)

	// 反转链表，反转n-m次
	var prev, tmp *ListNode = nil, nil
	cur := subHead
	for i := 0; i < n-m; i++ { // 题已给定1 ≤ m ≤ n ≤ 链表长度。
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	fmt.Println("prev: ", prev)
	fmt.Println("cur: ", cur)

	// pre就是区间部分链表反转后的头结点
	// 将区间链表重新连回去
	left.Next = prev
	subHead.Next = cur

	return dummy.Next
}


