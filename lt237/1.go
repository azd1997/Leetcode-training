package lt237

// 删除链表节点

//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
//
//现有一个链表 -- head = [4,5,1,9]，它可以表示为:
//
//
//
// 
//
//示例 1:
//
//输入: head = [4,5,1,9], node = 5
//输出: [4,1,9]
//解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//示例 2:
//
//输入: head = [4,5,1,9], node = 1
//输出: [4,5,9]
//解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
// 
//
//说明:
//
//链表至少包含两个节点。
//链表中所有节点的值都是唯一的。
//给定的节点为非末尾节点并且一定是链表中的一个有效节点。
//不要从你的函数中返回任何结果。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
// 根据题给的说明，所提交的函数几乎不用处理任何特殊情况......


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


// 由于是单链表，先要删除节点通常要知道前驱结点，再修改前驱结点next，但是题目没给head节点指针，这样就没法遍历node以前部分
// 另一个删除办法是： 修改node的后继结点和值，替换成下一个。这也正是这道题的解法
func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}