package lt160

// 相交链表
// 编写一个程序，找到两个单链表相交的起始节点。

//注意：
//
//如果两个链表没有交点，返回 null.
//在返回结果后，两个链表仍须保持原有的结构。
//可假定整个链表结构中没有循环。
//程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
//	1. 好像只能使用两层循环(或者说双指针)去遍历寻找是否有相同节点(注意是相同，意味着节点的内存地址要是同一个)
// 但这个思路基本上得算暴力解， O(N*M)
// 测试后发现测例其实还是强调节点的值相同....
// 并且还要检查这个相同节点后边的所有值是否都相等，这样才叫相交
// 但是这个思路来讲，题目的第一个示例明显通不过。这TM...

// 看题解
// 题目的意思是两个链表，到了某一个相交节点之后，相交节点及其之后的部分完全一致
// 而之前的节点哪怕数值一样，地址也是不一样的。
// 官方题解给出了暴力解、哈希表法(暴力解查询优化)、双指针法三种
// 都实现一遍吧



// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}



// 1. 我的暴力解，错解
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	p1, p2 := headA, headB
//	for p1 != nil {
//		for p2 != nil {
//			if p1.Val == p2.Val {return p1}
//			p2 = p2.Next
//		}
//		p1 = p1.Next
//	}
//	return nil	//找不到
//}

// 1. 暴力解
// 官方题解的暴力解思路与我完全一致，
// 但是bug在于：我已开始设置判等条件为p1==p2，结果错了
// 后来我改成值相等p1.Val==p2.Val也错了
// 这就推翻了必须要求节点地址相同和仅要求值相等两种情况
// 所以题目要的是节点值相等且next指针相等，但没要求节点地址一致
// 毕竟测例是通过两个数组生成的，其中节点地址不同，除非生成时先找出数组后边的相交部分
// 按照这个判定条件再试试
// 也没通过。TMD!








// 哈希表优化暴力解
// 看到题解区的这个解法我是真的蒙了，因为思路完全就一样的
//45/45 cases passed (52 ms)
//Your runtime beats 31.91 % of golang submissions
//Your memory usage beats 79.44 % of golang submissions (7.3 MB)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p1,p2 := headA,headB
	// 先将headA链表存入哈希集合
	set := make(map[*ListNode]bool)
	for p1!=nil {
		set[p1] = true
		p1 = p1.Next
	}
	for p2!=nil {
		if set[p2] {return p2}
		p2 = p2.Next
	}

	return nil
}

// 这么看来，我的暴力解是有问题的，但暂时没想明白



// 既然暴力解没做出来，哈希表解也是无效的
// 现在来试试最后的双指针法
// 这是利用相交链表的一个特性：相交部分完全一致 => 相交部分长度一致
// 有的做法是把长链表先走(减去长的部分)
// 有的做法是将短链表补长
// 还有一种是长短链表都走一遍，对齐长度

// 第三种思路解法如下
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	p1,p2 := headA,headB
	// p1,p2步长一致
	// p1从headA->headA尾->headB开始
	// p2从headB->headB尾->headA开始
	// 这么做的意义是消除两个链表的长度差
	for p1 != p2 {
		if p1 == nil {    // 如果第一次遍历到链表尾部，就指向另一个链表的头部，继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是 nil ==  nil
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
