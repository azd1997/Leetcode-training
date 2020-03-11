package traverse

import (
	"fmt"
	"testing"
)

var treedata = []int{1, 2, 3, 4, 5, 6, 7, 8}

var tree *TreeNode

func init() {
	tree = GenBinaryTree(treedata)
}

func TestPreOrderTraverseRecurse(t *testing.T) {
	arr := PreOrderRecurse(tree)
	fmt.Println(arr)
	fmt.Println(treedata)
	if arr == nil {
		t.Error("1")
	}

}

// [5 3 2 1 4 7 6 8]

func TestPreOrderTraverseIterate1(t *testing.T) {
	arr := PreOrderIterate1(tree)
	fmt.Println(arr)
	fmt.Println(treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [5 3 2 1 4 7 6 8]

func TestPreOrderTraverseIterate2(t *testing.T) {
	arr := PreOrderIterate2(tree)
	fmt.Println(arr)
	fmt.Println(treedata)
	if arr == nil {
		t.Error("1")
	}
}

// [5 3 2 1 4 7 6 8]
