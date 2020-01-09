package lt234

// 回文链表

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-linked-list
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


// 思考：
// 1. 由于是单链表， 很直接的一种做法就是先遍历一遍，将数据存到数组中，再对数组用首尾双指针向内移动不断比较的方法。 O(n)/O(n)
// 2. 想要在O(1)空间完成，那么只能翻转一半链表。


// 1. 拷贝到数组中再双指针
//26/26 cases passed (8 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 25.41 % of golang submissions (7.1 MB)
func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	// 双指针
	l, r := 0, len(nums)-1
	for l<r {
		if nums[l] != nums[r] {return false}
		l++
		r--
	}

	return true
}


// 2. 反转后半部分链表
// NOTICE! 原输入链表被修改了！工程实现最好先把链表数据完整拷贝一份再操作
//26/26 cases passed (20 ms)
//Your runtime beats 14.02 % of golang submissions
//Your memory usage beats 76.23 % of golang submissions (6 MB)
func isPalindrome2(head *ListNode) bool {
	// 边界条件
	if head==nil || head.Next==nil {return true}

	// 快慢指针(快指针走两步慢指针走一步)，一次遍历找到中间节点		（找中间节点的另一种做法是先遍历到底，记录节点个数，再直接遍历一半到中间节点）
	// 情况一  1->2->3->4->5	// mid
	// 情况二  1->2->3->4
	// 由于fast指针每次走两步
	fast, low := &ListNode{-1, head}, &ListNode{-1, head}		// 需要设一个无意义头结点，否则后面不好将两种情况统一
	for fast != nil && fast.Next != nil {	// 情况一下fast会到"5"的下一个，low到"3"； 情况二下fast到达"4"，low则到达"2"
		low, fast = low.Next, fast.Next.Next
	}
	// 结束后low处于中间节点位置。 只需反转low后的节点就行（不含low）
	// 反转操作可以通过交换值，也可以更换Next
	// 这里选择修改后继结点指针
	cur := low.Next		// 从这里开始修改，注意在将cur.Next修改为前一个节点之前，先将cur.Next拷贝
	var pre, tmp *ListNode = nil, nil
	low.Next = nil		// 这是为了将前半部分链表打断。（对于情况一，前半部分最后会多一个节点，但这个节点不处理）; 前半部分从head开始； 不打断也没关系
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	// 停止后，cur变成nil，pre则是后半部分翻转后的真正第一个节点
	// 这里要很小心，不要使用cur.Next != nil 作为迭代终止条件，会出错

	// 比较两部分链表，以后半部分遍历完为停止条件
	front := head
	for pre != nil {
		if pre.Val != front.Val {return false}
		pre = pre.Next
		front = front.Next
	}

	return true
}

// 3. 反转前半部分链表，可以边找中间节点边反转前半部分
//26/26 cases passed (16 ms)
//Your runtime beats 74.87 % of golang submissions
//Your memory usage beats 61.48 % of golang submissions (6.1 MB)
func isPalindrome3(head *ListNode) bool {
	// 边界条件
	if head==nil || head.Next==nil {return true}

	// 快慢指针(快指针走两步慢指针走一步)，一次遍历找到中间节点		（找中间节点的另一种做法是先遍历到底，记录节点个数，再直接遍历一半到中间节点）
	// 情况一  1->2->3->4->5	// mid
	// 情况二  1->2->3->4
	// 由于fast指针每次走两步
	fast, low := head, head
	var pre, prepare *ListNode = head, nil	// pre表示


	for fast != nil && fast.Next != nil {	// 情况一下fast会到"5"，low到"3"； 情况二下fast到达"4"后一个，low则到达"3"
											// 两种情况合并就是将 low前面的部分反转。 但是在下边这么反转，会发现多反转一个
											// 不论是情况一还是情况二都多反转了一个，但是怎么保留下那个呢？（我们需要后半部分链表的头节点）
											// 方法是在翻转前半部分的时候 访问到low时翻转low前一个而不是low，这样就可以少翻转一次。
											// 见前面 pre, prepare *ListNode = head, nil
		fast = fast.Next.Next

		// 反转low沿途节点
		pre = low	// pre保存下low，low即将更新
		low = low.Next
		pre.Next = prepare		// 修改后继指针
		prepare = pre
	}
	// 操作结束后， pre就是前半部翻转后的链表的头结点； 而原链表后半部分通过以low(情况二)或者low.Next(情况一)开头
	if fast != nil {	// 说明是情况一
		low = low.Next
	}
	// 现在后半部分通过low开头访问

	for low != nil {	// pre != nil 也可以，两者长度一样
		if pre.Val != low.Val {return false}
		pre, low = pre.Next, low.Next
	}
	return true
}