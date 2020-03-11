package lt669

// 修剪二叉搜索树

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 示例 1:

// 输入:
//     1
//    / \
//   0   2

//   L = 1
//   R = 2

// 输出:
//     1
//       \
//        2
// 示例 2:

// 输入:
//     3
//    / \
//   0   4
//    \
//     2
//    /
//   1

//   L = 1
//   R = 3

// 输出:
//       3
//      /
//    2
//   /
//  1

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/trim-a-binary-search-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 如果是普通的二叉搜索树，那么需要遍历整棵树，对所有子树都去找有没有要修剪的
// 但是这里是二叉搜索树，因此可以直接通过中序遍历定位到L和R处，
// 假设往左边找到结点A， A.Val >= L，并且A的左子树的最大值<L，那么直接将为A删除其左子树
// 同理，找到B，使B满足最右边 B.Val <= R， 直接删除其右子树

// 但是按照这个思路，可以看到是有漏洞的（见上面的例子）

// 也就是说，尽管是二叉搜索树也需要遍历整棵树，只是中间部分不需要检查

// 前边瞎想...没有用

// 递归真的是玄学啊，不讲道理，根本不关系中间是怎么修剪的过程，就像那道汉诺塔...

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	// 比较当前节点与边界的关系
	if root.Val > R {
		return trimBST(root.Left, L, R)
	} // >R时修剪后的树一定是来自左边
	if root.Val < L {
		return trimBST(root.Right, L, R)
	} // <L时修剪后的树一定是来自右边

	// 更新左右孩子
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)

	// 返回更新后的root
	return root
}
