package lcof25


//合并两个排序的链表，其实这个题思路很简单就是构建一个dummyHead，然后按双指针写法将两个链表头部较小者出列接到dummyHead后边。

//双指针解法：

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    p, p1, p2 := dummyHead, l1, l2
    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            p.Next = p1
            p = p.Next
            p1 = p1.Next    
        } else {
            p.Next = p2
            p = p.Next
            p2 = p2.Next
        }
    }

    if p1 != nil {
        p.Next = p1
    } else {
        p.Next = p2
    }

    return dummyHead.Next
}


// 递归解法
// 返回两个链表合并后的头结点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {return l2}    
    if l2 == nil {return l1}

    if l1.Val <= l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)        
        return l1
    } else {
        l2.Next = mergeTwoLists(l1, l2.Next)        
        return l2
    }
}