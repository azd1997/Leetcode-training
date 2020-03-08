package lcof24

// 反转链表

// 迭代做法：从左向右交换指针方向

func reverseList(head *Node) *Node {
    // 边界情况
    if head == nil || head.Next == nil {return head}    

    // 反转链表
    var pre, tmp *Node
    cur := head
    for cur != nil {
        tmp = cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    return pre
}



// 反转链表，按照链表的递归遍历，在回溯时修改指针指向

func reverseList2(head *Node) *Node {
    // 边界情况
    if head == nil || head.Next == nil {return head}
    // 不仅仅是特殊情况处理，也是递归时到链表末尾的边界

    // 递去
    ret := reverseList(head.Next)
    
    // 递归时转向，并把当前节点返回    
    head.Next.Next = head   // 指针转向
    head.Next = nil

    // 返回新的头结点
    return ret
}


