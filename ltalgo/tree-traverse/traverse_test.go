package traverse

import (
	"fmt"
	"testing"
)

var treedata = []int{1, 2, 3, 4, 5, 6, 7, 8}

var tree *TreeNode

func init() {
	tree = GenBinaryTree(treedata)
	fmt.Println("原始数组：", treedata)
}

func TestPreOrderTraverseRecurse(t *testing.T) {
	fmt.Println("前序遍历递归版本")
	arr := PreOrderRecurse(tree)
	fmt.Println("前序递归：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}

}

// [5 3 2 1 4 7 6 8]

func TestPreOrderTraverseIterate1(t *testing.T) {
	fmt.Println("前序遍历迭代版本1")
	arr := PreOrderIterate1(tree)
	fmt.Println("前序迭代1：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [5 3 2 1 4 7 6 8]

func TestPreOrderTraverseIterate2(t *testing.T) {
	fmt.Println("前序遍历迭代版本2")
	arr := PreOrderIterate2(tree)
	fmt.Println("前序迭代2：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [5 3 2 1 4 7 6 8]

func TestInOrderTraverseRecurse(t *testing.T) {
	fmt.Println("中序遍历递归版本")
	arr := InOrderRecurse(tree)
	fmt.Println("中序递归：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [1 2 3 4 5 6 7 8]

func TestInOrderTraverseIterate2(t *testing.T) {
	fmt.Println("中序遍历迭代版本")
	arr := InOrderIterate2(tree)
	fmt.Println("中序迭代：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [1 2 3 4 5 6 7 8]

func TestPostOrderTraverseRecurse(t *testing.T) {
	fmt.Println("后序遍历递归版本")
	arr := PostOrderRecurse(tree)
	fmt.Println("后序递归：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [1 2 4 3 6 8 7 5]

func TestPostOrderTraverseIterate(t *testing.T) {
	fmt.Println("后序遍历迭代版本")
	arr := PostOrderIterate(tree)
	fmt.Println("后序迭代：", arr)
	fmt.Println("原始数据：", treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [1 2 3 4 5 6 7 8]
