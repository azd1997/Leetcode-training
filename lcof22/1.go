package lcof22

// 链表中倒数第k个元素


//Definition for singly-linked list.
type ListNode struct {
   Val int
   Next *ListNode
}



// 链表中倒数第k个结点。

// 注意点：k的有效性

//思路：
//- 遍历后将所有节点值输出到数组，利用数组下标取倒数第k个结点值
//- 一次遍历得到链表长度length，再一次遍历遍历length-k，到达目标节点
//- 一次遍历，先右移k次，再用双指针，p1从head出发，p2从当前出发，p2遍历到尾时，p1所指即为所求

//一次遍历解法迭代实现：


func getKthFromEnd(head *ListNode, k int) *ListNode {
    p1, p2 := head, head

    // 先定位到kthNode
    for i:=0; i<k; i++ {
        if p2 != nil {
            p2 = p2.Next
        } else {
            return nil      // 链表长度<k
        }
    }

    // 再双指针
    for p2 != nil {
        p2 = p2.Next
        p1 = p1.Next
    }

    return p1
}