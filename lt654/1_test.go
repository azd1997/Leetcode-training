package lt654

import (
	"fmt"
	"testing"
)

// 测试线段树查询
func TestSegTreeQuery(t *testing.T) {
	arr := []int{3, 2, 1, 6, 0, 5}
	tree := buildSegmentTree(arr)
	fmt.Println(tree)
	fmt.Println(querySegmentTree(arr, tree, 0, 0, len(arr)-1, 1, 5))
}

// 测试线段树解法
func TestSol2(t *testing.T) {
	arr := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree2(arr)
	fmt.Println(root.Val)
}
