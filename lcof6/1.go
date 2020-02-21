package lcof6

// 从头到尾打印链表

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}


// 思考：
// 1. 从前向后遍历链表，存至到数组，再将数组反转
// 2. 翻转链表并访问
// 3. 利用翻转链表的方式去递归倒序的访问，并不真的翻转(这种递归是想要的答案)
// 4. 先遍历一遍链表，得到长度，再遍历一遍，倒序填入数组(稳妥，不易出错)
// 5. 一次遍历，每次得到的结果都存入数组(res = append([]int{cur}, res...))，但是这种情况下需要频繁的拼接数组
// 6. 遍历链表存到栈中，再从栈取出

// 1. 两次遍历
func reversePrint1(head *ListNode) []int {
	if head == nil {return nil}

	length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length++
	}

	res := make([]int, length)
	idx := length - 1
	cur = head
	for cur != nil {
		res[idx] = cur.Val
		cur = cur.Next
		idx--
	}

	return res
}


// 2. 递归
// 子函数取为: 将以node为头的链表的最后一个节点追加到res数组，同时返回res的新状态
func reversePrint2(head *ListNode) []int {
	res := make([]int, 0)
	return *reverse(head, &res)
}

func reverse(node *ListNode, res *[]int) *[]int {
	if node == nil {
		return res		// 到最后空的时候把res的状态返回(此时res还是一个空数组)
	}

	// 递去 ->
	// 递归 <-

	// res为链表后边处理得到res(倒序存着值)
	res = reverse(node.Next, res)
	// 把当前节点的值存入
	*res = append(*res, node.Val)
	return res
}

// 3. 既然递归可以，循环+栈 自然也可以
