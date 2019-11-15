package lt101

import "fmt"

// 对称二叉树
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}




// 1.递归

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric1(root *TreeNode) bool {
	return compare(root, root)
}

// 递归做法
func compare(left, right *TreeNode) bool {
	// 比较两个节点值是否一致，一致则递归比较其左右两子节点

	if left == nil && right == nil {    // 到叶子节点
		return true
	}
	if left == nil || right == nil {
		return false
	}


	return left.Val == right.Val && compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

// 2.迭代DFS（失败版）
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric2(root *TreeNode) bool {

	// 迭代解法也就是DFS栈解法，一个中左右，一个中右左。（只要保证左右访问顺序是相反的就行，由于前序最好实现，比较起来也是最快的）

	// 坑爹的是测例数组中的null对应的节点不是nil

	if root == nil {return true}
	if root.Left == nil && root.Right == nil {return true}
	if root.Left == nil || root.Right == nil {return false}

	// DFS
	stack1, stack2 := []*TreeNode{root.Left}, []*TreeNode{root.Right}
	visited1, visited2 := make(map[*TreeNode]bool), make(map[*TreeNode]bool)
	var cur1, cur2 *TreeNode

	for len(stack1) != 0 && len(stack2) != 0 {
		cur1, cur2 = stack1[len(stack1)-1], stack2[len(stack2)-1]
		fmt.Printf("stack1=%v, stack2=%v\n", stack1, stack2)
		stack1, stack2 = stack1[:len(stack1)-1], stack2[:len(stack2)-1]


		if visited1[cur1] || visited2[cur2] {     //其实按照逻辑来讲，只要是能够压入栈的，两个栈的访问次序是一致的，所以这里是与还是或都无所谓
			continue
		}

		if cur1.Val != cur2.Val {return false}
		visited1[cur1], visited2[cur2] = true, true

		// 不同时为nil说明不相等、立刻return false
		if (cur1.Left != nil && cur2.Right == nil) || (cur1.Left == nil && cur2.Right != nil) ||
			(cur1.Right != nil && cur2.Left == nil) || (cur1.Right == nil && cur2.Left != nil) {
			return false
		}

		// 同时不为nil需要压入栈中
		if cur1.Left != nil && cur2.Right != nil {
			stack1 = append(stack1, cur1.Left)
			stack2 = append(stack2, cur2.Right)
		}
		if cur1.Right != nil && cur2.Left != nil {
			stack1 = append(stack1, cur1.Left)
			stack2 = append(stack2, cur2.Right)
		}

		// 同时为nil，不用添加到stack,什么都不做
	}

	if len(stack1) != len(stack2) {return false}


	return true
}

// 失败测例： [1,3,3,null,3,3]
// 输出：stack1=[0xc00000a160], stack2=[0xc00000a180]
//stack1=[<nil>], stack2=[<nil>]
// Line 34: panic: runtime error: invalid memory address or nil pointer dereference		（Line34是  if cur1.Val != cur2.Val {return false} ）
//预期答案：true

// 极其纳闷：我后面的代码写得是必须不为nil才可能压入栈，怎么还能空指针调用

// 3.一种正确的迭代DFS写法

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric3(root *TreeNode) bool {

	if root == nil {return true}
	if root.Left == nil && root.Right == nil {return true}
	if root.Left == nil || root.Right == nil {return false}

	// DFS
	stack := []*TreeNode{root.Left, root.Right}
	var cur1, cur2 *TreeNode

	for len(stack) != 0 {
		cur1, cur2 = stack[len(stack)-2], stack[len(stack)-1]
		//fmt.Printf("stack1=%v, stack2=%v\n", stack1, stack2)
		stack = stack[:len(stack)-2]

		if cur1 == nil && cur2 == nil {continue}
		if cur1 == nil || cur2 == nil {return false}
		if cur1.Val != cur2.Val {return false}
		stack = append(stack, cur1.Left, cur2.Right, cur1.Right, cur2.Left)
	}

	return true
}
