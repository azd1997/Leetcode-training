package lt108

// 将有序数组转换为二叉搜索树

//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 由于树做的比较少，这里搬运题解区 江不知 的描述：

//二叉搜索树
//二叉搜索树（Binary Search Tree）是指一棵空树或具有如下性质的二叉树：
//
//若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值
//若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值
//任意节点的左、右子树也分别为二叉搜索树
//没有键值相等的节点
//基于以上性质，我们可以得出一个二叉搜索树的特性：二叉搜索树的中序遍历结果为递增序列。
//
//那么现在题目给了我们一个递增序列，要求我们构造一棵二叉搜索树，就是要我们实现这一特性的逆过程。
//
//还记得什么是中序遍历吗？中序遍历的顺序为：左节点 \to→ 根节点 \to→ 右节点。这个遍历过程可以使用递归非常直观地进行表示。
//
//如何构造树
//构造一棵树的过程可以拆分成无数个这样的子问题：构造树的每个节点以及节点之间的关系。对于每个节点来说，都需要：
//
//选取节点
//构造该节点的左子树
//构造该节点的右子树
//因题目要求构造一棵「高度平衡」的树，所以我们在选取节点时选择数组的中点作为根节点，以此来保证平衡性。
//
//以题目给出的 [-10,-3,0,5,9] 为例。
//
//我们选取数组中点，即数字 0 作为根节点。此时，以 0 为分界点将数组分为左右两个部分，左侧为 [-10, -3]，右侧为 [5, 9]。因该数组为升序排列的有序数组，所以左侧数组值均小于 0，可作为节点 0 的左子树；右侧数组值均大于 0，可作为节点 0 的右子树。
//
//作者：jalan
//链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/solution/tu-jie-er-cha-sou-suo-shu-gou-zao-di-gui-python-go/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 也就是不断区区间中点作为根节点，构建左右子树，直至再无左右子树

// 1. 递归做法
// 注意：由于题解不唯一，所以单个测例测试时可能与题目给出答案不吻合，但提交时却只看是否高度平衡
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {return nil}
	mid := len(nums)/2
	left, right := nums[:mid], nums[mid+1:]
	node := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
	return node
}


// 以下参考题解区 Windliang

// 2. 基于栈的解法
// 一部分递归可以转为动态规划问题实现空间换时间，从自顶向下再向顶改成了自底向上
// 而另一部分则不行，只能用栈去模拟递归过程，没啥好处，但是能让人更清楚递归的过程
// 迭代栈 模拟 递归
// 递归就是压栈出栈过程，因此需要一个栈保存递归的参数
// 在这里，就是子树的数组前后下标，以及内部定义的root(也就是解法1里的那个node)
// 由于go中切片是引用传递，并不占额外空间，所以，在go中可以使用切片来代替前后下标（只要这个前后下标是为了划子区间而不是别的目的，就可以这样使用）
// 为了方便，定义一个新结构体，用来存储这两个参数
type MyNode struct {
	node TreeNode
	slice []int
}

// 模拟递归 自顶向下再向顶
func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 初始化栈，用切片模拟
	stack := make([]*MyNode, 0)

	// 1. 将根节点信息压栈
	mid := len(nums) / 2              // 由于这里求中数是数组索引求中，其实不用担心溢出，数组没那么大
	node := &TreeNode{Val: nums[mid]} // 此时是真正根节点
	stack = append(stack, &MyNode{*node, nums})
	nums1 := nums //避免修改nums

	// 只要nums1还能再细分或者栈未空，就得继续迭代
	for len(nums1) > 1 || len(stack) > 0 {

		// 1. 迭代栈，模拟递归， 不停生成左子树
		// 这一步会把所有的左子树找齐
		for len(nums1) > 1 { // 要生成子树，个数需要大于1
			nums1 = nums1[:mid]  // 左子树
			mid = len(nums1) / 2 // 左子树中点
			node.Left = &TreeNode{Val: nums1[mid]}
			node = node.Left // 全局根节点转移为其左子树根节点
			stack = append(stack, &MyNode{*node, nums1})
		}

		// 2. 无法生成左子树后，开始出栈，生成右子树
		// 这时会将原来的区间从小到大不断弹出来，不断生成这些区间的右子树根节点
		mynode := stack[len(stack)-1];
		stack = stack[:len(stack)-1] // 出栈
		minSlice := mynode.slice     // 就是指前面生成左子树过程中最小的用来得左子树根节点的区间
		mid = len(minSlice) / 2      // 其实这里不用重新赋值，就是上面遗留下来的mid
		nums1 = minSlice[mid+1:]
		// 3. 弹出的区间越来越大，其中需要重新将右子树根节点及新的更小的右子树对应的区间压入栈
		for len(nums1) > 1 { // 如果有两个数及以上的话，就要往下构建右子树
			mid = len(nums1) / 2 // 右子树mid
			node.Right = &TreeNode{Val: nums1[mid]}
			node = node.Right
			stack = append(stack, &MyNode{*node, nums1})
		}
	}

	// 最后弹出的node就是一开始压入的根节点
	return node
}


// 醉了，这模拟得比递归解法本身难理解多了...
// 而且出错了

// 先使用前后索引而不是子区间切片，完整还原Windliang答案来看看
type MyNode2 struct {
	node TreeNode
	start, end int
}

// 模拟递归 自顶向下再向顶
func sortedArrayToBST22(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 初始化栈，用切片模拟
	stack := make([]*MyNode2, 0)

	// 1. 将根节点信息压栈
	start, end := 0, len(nums)
	mid := (start + end) / 2
	root := &TreeNode{Val:nums[mid]}
	curRoot := root
	stack = append(stack, &MyNode2{*curRoot, start, end})

	for end - start > 1 || len(stack) > 0 {

		// 1. 迭代栈，模拟递归， 不停生成左子树
		// 这一步会把所有的左子树找齐
		for end - start > 1 { // 要生成子树，个数需要大于1
			mid = (start + end) / 2		// 当前根节点
			end = mid	// 左子树区间的末尾
			mid = (start + end) / 2		// 当前根节点左子树的中点
			curRoot.Left = &TreeNode{Val:nums[mid]}
			curRoot = curRoot.Left
			stack = append(stack, &MyNode2{*curRoot, start, end})
		}

		// 2. 无法生成左子树后，开始出栈，生成右子树
		// 这时会将原来的区间从小到大不断弹出来，不断生成这些区间的右子树根节点
		mynode := stack[len(stack)-1]; stack = stack[:len(stack)-1] // 出栈
		start, end = mynode.start, mynode.end
		mid = (start + end) / 2
		start = mid + 1		// 右子树的start
		curRoot = &mynode.node		// 当前根节点
		if start < end {	// 如果范围内有数
			mid = (start + end) / 2
			curRoot.Right = &TreeNode{Val:nums[mid]}
			curRoot = curRoot.Right
			stack = append(stack, &MyNode2{*curRoot, start, end})
		}
	}

	return root
}

// 同样没能通过测试，暂时放弃，不想浪费时间

// 3. BFS
type MyNode3 struct {
	root *TreeNode
	start, end int
}

func sortedArrayToBST3(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 初始化队列，用切片模拟
	queue := make([]*MyNode3, 0)

	// 随便构造一个根节点值载入队列
	root := &TreeNode{Val:0}
	queue = append(queue, &MyNode3{root, 0, len(nums)-1})

	// BFS
	var (
		myroot *MyNode3
		start, end, mid int
		curRoot *TreeNode
	)
	for len(queue)>0 {
		myroot = queue[0]; queue = queue[1:]	// 弹出队首
		start, end = myroot.start, myroot.end
		mid = (start + end) / 2
		curRoot = myroot.root
		curRoot.Val = nums[mid]		// 根节点赋值
		if start<mid {	// [start, mid]还可分
			curRoot.Left = &TreeNode{Val:0}		// 也是先随便填一个0
			queue = append(queue, &MyNode3{curRoot.Left, start, mid})
		}
		if mid+1<end {	// [mid+1, end]还可分
			curRoot.Right = &TreeNode{Val:0}
			queue = append(queue, &MyNode3{curRoot.Right, mid+1, end})
		}
	}

	// 分析一下：一开始都只是填充了个0进去，后来才赋值

	return root
}

// 崩了，又没通过
