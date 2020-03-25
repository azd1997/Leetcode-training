package lt437

// 路径总和III

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 这道题标的easy，但我觉得是三道路径总和中最难的...
//
// 需要做一些后缀和或者说前缀和的处理
// 还需要进行回溯
// 也就是说路径总和III需要用到 递归+回溯+前缀和
// 而前两道只是无脑递归

// 如果两个数的前缀和相同，则这两个数之间的数的和为0
// 如果在 结点A 和 结点B 的前缀总和相差target，则从A到B之间的元素和为target
//
// 因此对于本题来讲，当抵达结点B时，得到当前前缀和cursum，
// 再寻找root->B路径上是否存在前缀和为cursum - target的结点
// 每次找到路径上符合条件的结点后，累加到res返回结果变量上，再递归进入到左右子树
//
// 问题来了：如何记录前缀和
// 当然可以根据当前树重新构建一颗树，树的每个节点包含前缀和和值两个值域
// 这里可以将前缀和记录在结点上，覆盖原本的Val，
// 但是这样需要在走完一条之路后将Val还原
// 也就是需要回溯
// 下面其实不是这么做而是用哈希表记录的数量以及递归栈传递
//
// 另外一个问题是：如何比较高效的统计前面说的符合条件（cursum-target）的结点呢？
// 总不能又往回跑吧
// 这就需要使用哈希表去记录，方便查找

func pathSum(root *TreeNode, sum int) int {
	m := make(map[int]int)       // <prefixsum, count>
	m[0] = 1                     // 前缀和为0的路径初始就有一条，就是哪个结点都不挨
	return help(root, m, sum, 0) // 0是当前的前缀和
}

func help(root *TreeNode, m map[int]int, target, cursum int) int {
	// 递归终止
	if root == nil {
		return 0
	}

	// 本层要做的事
	res := 0
	cursum += root.Val // 当前前缀和
	// 更新res
	res += m[cursum-target] // 不存在就会是0
	// 更新路径上前缀和个数
	m[cursum]++

	// 进入下一层
	res += help(root.Left, m, target, cursum)
	res += help(root.Right, m, target, cursum)

	// 回到本层，恢复状态，去除当前节点的前缀和数量
	m[cursum]--

	return res
}
