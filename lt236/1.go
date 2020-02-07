package lt236

// 二叉树的最近公共祖先


/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}



// 先将下自己的思考：
// 利用中序遍历，栈顶节点的祖先是其栈顶下面的节点，这形成了祖先树
// 那么利用DFS中序遍历，维护两个栈，分别找到目标节点为止
// 把一个栈的节点存入哈希集合， 再倒序遍历另一个栈，看哪个最先在哈希集合中存在



// 1. DFS迭代 中序遍历 记录两个节点的祖先线
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pAncestors := findAncestors(root, p)
	qAncestors := findAncestors(root, q)

	// 比较那个祖先线长，其实就是深度
	// 要把深度更大的那个存到哈希表中去
	long, short := pAncestors, qAncestors
	if len(pAncestors) < len(qAncestors) {
		long, short = qAncestors, pAncestors
	}

	// 将long存哈希表。(对于严格的栈肯定是不停出栈的，但我们这可以直接遍历数组)
	set := make(map[*TreeNode]bool)
	for _, v := range long {
		set[v] = true
	}

	// 倒序遍历short
	for i:=len(short)-1; i>=0; i-- {
		if set[short[i]] {return short[i]}
	}
	// 不可能没有公共祖先，最后这里随便返回就行
	return root
}

func findAncestors(root, target *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 1)

	// 找target
	for {
		for root != nil {
			// 先压栈
			stack = append(stack, root)
			// 判断是否是目标节点
			if root == target {return stack}
		}

		// 弹出栈顶
		// 前面for循环没有return说明没找到target，需要弹出栈顶
		// 去右子树搜索
		root = stack[len(stack)-1]; stack = stack[:len(stack)-1]
		root = root.Right
	}
}


// 感觉思路是没问题的，但是实现出了问题
// 不在一道题拖太久
// 看题解区


// 2. 来自题解区 软微最蔡

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	// root == q || root == p 如果为真，那么无论左右子节点有没有找到，
	// 都可以返回root。 第二，由于题目限定条件，当一个分支找到两个节点后，
	// 另一个分支一定就找不到了，返回的一定是NULL，这就不需要额外弄个计数。

	// root==nil说明到叶节点为空，直接返回
	if root==nil || root==p || root==q {return root}

	// p,q肯定存在与树中，因此在左子树和右子树中不断寻找
	// 找到叶子时，如果左叶子为nil，那么右叶子必然是目标节点，返回
	// 都不为空，说明都为目标节点，那么root就是最近公共祖先
	//
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left==nil {
		return right
	} else if right==nil {
		return left
	} else {
		return root
	}
}




// 3. 回溯   参考官方题解
// 其实就是DFS，当找到target时回溯，其所有回溯的上层节点(父级节点)都标记为true
// 当回溯时发现某个节点左右子节点都为true标记，那么node就是最近公共祖先LCA
// O(N)/O(N)
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {

	// 遍历树
	var ans *TreeNode
	recurseTree(root, p, q, &ans)
	return ans
}

func recurseTree(cur, p, q *TreeNode, ans **TreeNode) bool {
	// 如果到分支尾了，返回false(没找到p,q任何一个)
	if cur == nil {return false}

	// 左子树遍历，如果左子树递归过程返回了true(说明找到了p,q之一)，则记为为1
	left := bool2int(recurseTree(cur.Left, p, q, ans))

	// 右子树遍历
	right := bool2int(recurseTree(cur.Right, p, q, ans))

	// 当前节点是p,q之一
	mid := bool2int(cur == p || cur == q)

	// 如果cur节点收集了两个true，那么cur就是LCA
	sum := mid + left + right
	if sum >= 2 {*ans = cur}

	// 如果至少有p,q之一，返回true
	return sum > 0
}

func bool2int(b bool) int {
	if b {return 1} else {return 0}
}


// 4. 使用父指针迭代
// 这个思路和我一开始想法很像，只不过我是想把沿路的父节点记录在栈中
// 而这个解法是用哈希表记录每一个节点的父节点
// O(N)/O(N)
func lowestCommonAncestor4(root, p, q *TreeNode) *TreeNode {
	// 1. 构造栈。 这个栈用来辅助DFS寻找两个目标节点，都找到时停止
	stack := make([]*TreeNode, 1)

	// 2. 构造哈希表，记录每一个节点的父节点
	parent := make(map[*TreeNode]*TreeNode)

	// 3. 将root存入  (base case)
	stack[0] = root
	fake := &TreeNode{}
	parent[root] = fake	// root节点没有父节点
	// 之所以要构造伪祖宗而不是用默认nil，是为了能够处理p,q之一为root的情况

	// 4. 迭代直到找到p,q
	var node *TreeNode
	for parent[p]==nil || parent[q]==nil {
		// 弹出栈顶
		node = stack[len(stack)-1]; stack = stack[:len(stack)-1]
		// 把node左子节点加入栈中
		if node.Left!=nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		// 把node右子节点加入栈中
		if node.Right!=nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
	}

	// 5. 用哈希集合收集所有p的祖先
	set := make(map[*TreeNode]bool)
	for p != nil {
		set[p] = true	// 节点自身也算祖宗
		p = parent[p]
	}

	// 6. 查看p的祖先中哪个是q的祖先，由于是从q倒推，所以第一个吻合的是LCA
	for !set[q] {
		q = parent[q]
	}
	return q

}