package lt1343

// 分裂二叉树的最大乘积

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 根据题意，显然要先求整棵树的sum1，再求某颗子树的sum2，剩下的和为sum1-sum2
// product = sum2 * (sum1 - sum2)
// 求树节点和深度优先遍历
//
// 程序主体就直接深度优先搜索，碰到叶子节点就返回

//var MOD = int(math.Pow10(9)) + 7
var MOD = int(1e9) + 7
func maxProduct(root *TreeNode) int {
	totalSum := dfsSum(root)
	maxProduct := 0
	dfsTree(root, &maxProduct, totalSum)
	return maxProduct % MOD
}

func dfsSum(root *TreeNode) int {
	if root==nil {return 0}
	sum := root.Val
	return (sum + dfsSum(root.Left) + dfsSum(root.Right)) % MOD
}

// dfsTree在dfsSum的基础上加入了最大乘积的比较，其返回值也是节点值之和
func dfsTree(root *TreeNode, maxProduct *int, totalSum int) int {
	if root==nil {return 0}
	sum := root.Val
	sum += dfsTree(root.Left, maxProduct, totalSum)
	sum += dfsTree(root.Right, maxProduct, totalSum)
	tmp := sum * (totalSum - sum)
	if *maxProduct < tmp {*maxProduct = tmp}
	return sum
}

func max3(a,b,c int) int {
	sum := a
	if b>sum {sum = b}
	if c>sum {sum = c}
	return sum
}