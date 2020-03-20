package lt112

// 路径总和

// 这道题自顶向下或者自底向上都可以做，要么是从叶子到根，检查是否右路径和为sum，要么是从根往下
// 但是显而易见的是，自顶向下会好实现很多。这里自顶向下递归实现

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归，自顶向下
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	} // 到叶子结点之下时需检查sum是否恰好减为0
	// 处理当前
	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// 漂亮简洁，但是未通过。
// 原因是 存在测例 tree=[] sum=0
// 改写为如下

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return help(root, sum)
}

func help(root *TreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	} // 到叶子结点之下时需检查sum是否恰好减为0
	// 处理当前
	sum -= root.Val
	return help(root.Left, sum) || help(root.Right, sum)
}

// 还是不通过，不通过测例：tree=[1,2], sum=1
// 说明root自己不算叶子节点
// 题目中说叶子节点指没有孩子的结点...

// 正解：

func hasPathSum3(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return help3(root, sum)
}

func help3(root *TreeNode, sum int) bool {
	if root.Left == nil && root.Right == nil { // 叶子结点
		return sum-root.Val == 0 // 检查sum是否减为0
	}
	// 处理当前
	sum -= root.Val
	leftok, rightok := false, false
	if root.Left != nil {
		leftok = help3(root.Left, sum)
	}
	if root.Right != nil {
		rightok = help3(root.Right, sum)
	}
	return leftok || rightok
}

// 测例
// tree=[5,4,8,11,null,13,4,7,2,null,null,null,1],sum=22
// tree=[], sum=0
// tree=[1,2], sum=1
// tree=[1], sum=1
