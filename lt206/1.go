package lt206

// 反转链表

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


type ListNode struct {
	Val int
	Next *ListNode
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// 1. 迭代做法
//27/27 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 66.97 % of golang submissions (2.6 MB)
func reverseList1(head *ListNode) *ListNode {
	var pre, tmp *ListNode = nil, nil
	cur := head
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre		// 最后pre是翻转后的头结点
}

// 递归做法需要隐式的调用栈，O(n)空间

// 2. 递归做法. 这个递归完全按照上面迭代的思路，是正向翻转来的
//27/27 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 10.55 % of golang submissions (2.9 MB)
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode = nil
	return reverse(pre, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {return pre}
	tmp := cur.Next
	cur.Next = pre
	return reverse(cur, tmp)
}

// 3. 递归 参考官方题解 从后向前翻转，【假设】后边一部分已经翻转好，将翻转好的部分的尾结点指向它左边的邻节点
// 1->2->3->4<-5<-6		// 假设4<-5<-6已经翻转好，则要做的只是将4.Next指向3
//27/27 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 14.22 % of golang submissions (2.9 MB)
func reverseList3(head *ListNode) *ListNode {
	if head==nil || head.Next==nil {return head}	// head真正翻转是从倒数第二个节点开始
	p := reverseList3(head.Next)	// 由于是从后往前翻转，所以是先递归下一次的，再执行当前的动作; p表示反转后的新头
	head.Next.Next = head	// 先将当前节点(head)的后继结点(head.Next)的后即指针(head.Next.Next)指向当前节点
	head.Next = nil			// 再将当前节点的后即指针(head.Next)置空。为什么要置空？中间的步骤都可以不置空，但是在翻转原头结点（也就是最后一次翻转时）时，必须把它的后继置空，否则无限循环
	return p
	// 整体而言，这样的递归，最终从末尾开始翻转，慢慢翻到原来的头节点。
}