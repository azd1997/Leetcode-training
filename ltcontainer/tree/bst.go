package tree

import "reflect"

// Value 必须是可比较的，这里不做检查和约束
// 值。 空接口，类似泛型的方案
type Value interface{}

// compare 辅助函数，用来比较两个Value
// 0 相等
// 1 左大于右
// -1 右大于左
// -2 类型不一致
// -3 类型不可比较
// -4 未知Kind
func compare(v1, v2 Value) int {
	v1Value := reflect.ValueOf(v1)
	v2Value := reflect.ValueOf(v2)

	// 必须检查两个Value类型一致
	if v1Value.Type() != v2Value.Type() {
		return -2
	}

	//  可比较性
	if !v1Value.Type().Comparable() {
		return -3
	}

	// 比较
	switch v1Value.Kind() {
	case reflect.Int:
		if v1Value.Int() > v2Value.Int() {
			return 1
		} else if v1Value.Int() < v2Value.Int() {
			return -1
		} else {return 0}
	default:
		// 不知名的Kind报错
		return -4
	}
}

// Node 二叉搜索树节点
type Node struct {
	Val Value
	Left, Right *Node
}

// 二叉树节点该有的一些方法

// 1. 非空
// 2. 大小
// 3. 插入/更新
// 4. 按值删除
// 5. 包含

// BST 二叉搜索树
type BST struct {
	root *Node
	size int
}

// NewBST 新建空树
func NewBST() *BST {
	return &BST{}
}

// NewBSTFromSlice 根据切片新建二叉搜索树
func NewBSTFromSlice(data []Value) *BST {return nil}

// IsEmpty 判空
func (bst *BST) IsEmpty() bool {
	return bst.size == 0	// bst.root == nil 也可以
}

// Size 元素数量
func (bst *BST) Size() int {
	return bst.size
}

// Add 插入
// 不断将v与树原有节点进行比较，判断是否应该去其左右子树继续寻找（联系二分查找的思想）
func (bst *BST) Add(v Value) {
	// 树空
	if bst.IsEmpty() {bst.root.Val = v}

	// 递归寻找插入位置
	newRoot := add(bst.root, v)
	if newRoot == nil {return}	// 插入失败
	bst.root = newRoot
}

// 返回插入后树的新根（因为根可能会改变）
func add(root *Node, v Value) *Node {
	// 空节点，也就是找到可插入位置
	if root == nil {return &Node{Val:v}}
	// 发现重复元素
	if root.Val == v {return root}

	// 根据比较结果向左右子树试探
	cmp := compare(root.Val, v)
	switch cmp {
	case 1:
		root.Left = add(root.Left, v)
	case -1:
		root.Right = add(root.Right, v)
	default:
		// 其他结果都是不应当的
		return nil	// 应当报错
	}

	return root		// 返回当前根节点
}

// Del 按值删除
// 有三种情况
// 待删除节点为叶节点； 为不完全节点； 为具备两个孩子的完全节点
func (bst *BST) Del(v Value) {
	// 记得接收返回值，root可能更改
	bst.root = findAndDel(bst.root, v)
}

// 删除root子树中等于该v的节点。 返回新root
func findAndDel(root *Node, v Value) *Node {
	// 外层先是遍历框架。 找到待删除节点之后再根据待删除节点的情况去处理

	if root.Val == v {
		// 找到待删除节点.进行删除操作
		return del(root)
	}
	cmp := compare(root.Val, v)
	switch cmp {
	case 1:
		return findAndDel(root.Left, v)
	case -1:
		return findAndDel(root.Right, v)
	default:
		return root		// 返回root，删除失败
	}
}

// 删除当前节点， 返回顶替的节点（后继节点）
func del(cur *Node) *Node {
	// 情况1 cur为叶子节点，直接用nil节点顶替（注意：事实上情况1会被情况2处理掉）
	if cur.Left == nil && cur.Right == nil {return nil}
	// 情况2 cur为不完全节点，用不为空的孩子顶替
	if cur.Left == nil {return cur.Right}
	if cur.Right == nil {return cur.Left}
	// 情况3 cur为完全节点
	//  这种情况下需要找到左子树最大节点或者右子树最小节点来顶替自己
	//  也就是找自己数值相邻的两个之一顶替
	// 	找到后继者之后，返回后继者并且将原后继者位置删除
	// 方便的做法是直接将后继者值赋给cur，然后删除后继者
	// 这里采用右子树最小节点顶替（沿右子树左侧链路到底）
	rightMin := getMinNode(cur.Right)
	cur.Val = rightMin.Val
	// 看似是递归调用，其实由于rightMin是叶子或者不完全节点，一遍就删掉了
	cur.Right = findAndDel(cur.Right, rightMin.Val)
	return cur
}

func getMinNode(root *Node) *Node {
	for root.Left != nil {root = root.Left}
	return root
}

// Contains 包含
func (bst *BST) Contains(v Value) bool {
	return contains(bst.root, v)
}

func contains(root *Node, v Value) bool {
	// 递归终止
	if root == nil {return false}
	if root.Val == v {return true}
	// 否则类似二分，进入相应子树搜寻
	cmp := compare(root.Val, v)
	switch cmp {
	case 1:
		return contains(root.Left, v)
	case -1:
		return contains(root.Right, v)
	default:	// 报错,这里简单的返回false
		return false
	}
}

////////////////////////////////////////////

// AddOne 所有节点加1（假设Value存int，这里只作演示）
//  选择某一种遍历方式，注意使用标记避免重复加1
func (bst *BST) AddOne() {
	addOne(bst.root)
}

// 前序遍历,不要额外标记
func addOne(root *Node) {
	// 边界
	if root == nil {return}
	// 当前+1
	root.Val = root.Val.(int) + 1
	// 左子树
	addOne(root.Left)
	// 右子树
	addOne(root.Right)
}

// Equals 判断两棵树相同
// 同样也是选择一种遍历方式
func (bst *BST) Equals(bst2 *BST) bool {
	return equals(bst.root, bst2.root)
}

func equals(root1, root2 *Node) bool {
	// 都空
	if root1 == nil && root2 == nil {return true}
	// 有一个空
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {return false}
	// 都不空时值不等
	if root1.Val != root2.Val {return false}
	// 当前值相等，还需要比较子树
	return equals(root1.Left, root2.Left) && equals(root1.Right, root2.Right)
}

// IsValid 有效性
// 注意这里不能简单的将node/left/right三者的值进行比较
// 因为二叉搜索树的定义是 左子树的值均 < node < 右子树的值
func (bst *BST) IsValid() bool {
	return isValid(bst.root, nil, nil)
}

// 根据传入的最大最小值节点限制，判断root子树是否有效
// min，max为空时说明没有限制
func isValid(root, min, max *Node) bool {
	// 边界
	if root == nil {return true}
	// min限制且root不>left则不对（这里对异常情况统一作无效处理，
	// 是因为若元素都不可比较，自然树是无效的）
	if min != nil && compare(root.Val, min.Val) != 1 {
		return false
	}
	// max有限制且root不<right则不对
	if max != nil && compare(root.Val, max.Val) != -1 {
		return false
	}
	// 如果 left < root < right，还要递归检查左右子树
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}


//////////////////////////////////////////////////////////

// BST的遍历框架
// 基本上所有操作都是按这个模板，先找后操作
func traverse(root *Node, target Value) {
	// 1. 找到目标
	if root.Val == target {
		// do something
	}
	// 2. 向左右子树之一递归
	if compare(root.Val, target) == 1 {
		traverse(root.Left, target)
	}
	if compare(root.Val, target) == -1 {
		traverse(root.Right, target)
	}
}

// Tips: 如果当前节点的操作会对其下边的子树产生影响，
// 需要狗造辅助函数增长参数列表，借助参数传递信息
