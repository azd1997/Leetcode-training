package lt113

// 路径总和II

// 返回所有符合条件的路径

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 其实是一种【回溯】（DFS）

// 看下面的解法，看上去没有问题，而且很小心的在所有做了选择之后需要return的地方撤销了选择
// 但是仍然是错误的解法
// 这是因为path始终会修改，即便res追加的是*path，也没有用
// 要么在添加时拷贝，见pathSum1
// 要么传参时就拷贝。见pathSum2

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	help(root, sum, &path, &res)
	return res
}

func help(root *TreeNode, sum int, path *[]int, res *[][]int) {

	//fmt.Println(root, sum, *path)

	// 防止root为空，引起panic； 同时也是触底返回
	if root == nil {
		return
	}

	// 先更新path和sum（当前决策）
	*path = append(*path, root.Val)
	sum -= root.Val

	// 检查当前
	if root.Left == nil && root.Right == nil { // 叶子结点
		if sum == 0 { // 找到符合条件的路径
			*res = append(*res, *path)
			//fmt.Println(*res)
			// 撤销选择
			*path = (*path)[:len(*path)-1]
			return
		}
		// 撤销选择
		*path = (*path)[:len(*path)-1]
		return
	}

	// 继续处理下一层（下层的决策）
	help(root.Left, sum, path, res)
	help(root.Right, sum, path, res)

	// 撤销选择
	*path = (*path)[:len(*path)-1]
}

// 2.
func pathSum1(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	help1(root, sum, &path, &res)
	return res
}

func help1(root *TreeNode, sum int, path *[]int, res *[][]int) {

	//fmt.Println(root, sum, *path)

	// 防止root为空，引起panic； 同时也是触底返回
	if root == nil {
		return
	}

	// 先更新path和sum（当前决策）
	*path = append(*path, root.Val)
	sum -= root.Val

	// 检查当前
	if root.Left == nil && root.Right == nil { // 叶子结点
		if sum == 0 { // 找到符合条件的路径
			*res = append(*res, append([]int{}, (*path)...))
			//fmt.Println(*res)
			// 撤销选择
			*path = (*path)[:len(*path)-1]
			return
		}
		// 撤销选择
		*path = (*path)[:len(*path)-1]
		return
	}

	// 继续处理下一层（下层的决策）
	help1(root.Left, sum, path, res)
	help1(root.Right, sum, path, res)

	// 撤销选择
	*path = (*path)[:len(*path)-1]
}

// 3.
func pathSum2(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	help2(root, sum, &[]int{}, &res)
	return res
}

func help2(root *TreeNode, sum int, path *[]int, res *[][]int) {

	//fmt.Println(root, sum, *path)

	// 防止root为空，引起panic； 同时也是触底返回
	if root == nil {
		return
	}

	// 先更新path和sum（当前决策）
	*path = append(*path, root.Val) // 要警惕这步，一旦扩容，则发生变化了，但外部不可见这个变化
	sum -= root.Val

	// 检查当前
	if root.Left == nil && root.Right == nil { // 叶子结点
		if sum == 0 { // 找到符合条件的路径
			*res = append(*res, *path)
		}
		return
	}

	// 继续处理下一层（下层的决策）
	tmp1, tmp2 := append([]int{}, (*path)...), append([]int{}, (*path)...)
	help2(root.Left, sum, &tmp1, res)
	help2(root.Right, sum, &tmp2, res)

	// 撤销选择
	*path = (*path)[:len(*path)-1]
}
