package lcof52


// 两个链表的第一个公共结点

// 常规思路的话，是使用哈希表/集合记录一条链表的所有节点，再遍历另一条链表，
//第一个在哈希集合中遇到的结点就是第一个公共结点。

// 想要优化空间，则需要想办法在遍历两条链表的时候能够在公共结点相遇，
// 这样能找到公共结点甚至是第一个公共结点。

// 但是两条链表长度并不一定一样


// 图示：
// 链表A前半段              a1 -> a2 ->
//  公共段                                              c1 -> c2 -> c3
// 链表B前半段   b1 -> b2 -> b3 ->
// 
// 假设双指针(p1, p2)，每个指针都走一遍两条链表，路径如下：
// p1路径： a1 -> a2 -> c1 -> c2 -> c3  =>  b1 -> b2 -> b3 ->  c1 -> c2 -> c3
// p2路径： b1 -> b2 -> b3 ->  c1 -> c2 -> c3  =>  a1 -> a2 -> c1 -> c2 -> c3
// 如果两条链表没有交集，那么这样走一遍后是不会相遇的
// 像上面图示中，p1,p2会在"c1"处相遇，这就是相交点

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // 边界    
    if headA == nil || headB == nil {return nil}
    // 双指针
    p1, p2 := headA, headB
    for {
        // 如果双方同时到尾，说明走完了
        if p1 == nil && p2 == nil {return nil}

        // 换一条“跑道”继续移动
        if p1 == nil {p1 = headB}    
        if p2 == nil {p2 = headA}
        
        // 找到第一个交点
        if p1 == p2 {return p1}  

        // 否则向后移动
        p1, p2 = p1.Next, p2.Next  
    }
    // 返回nil
    return nil
}