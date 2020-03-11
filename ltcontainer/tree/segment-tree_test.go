package tree

import (
	"fmt"
	"testing"
)

func TestBuildSegmentTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	st := BuildTree(arr)
	fmt.Println(st.tree)
}

func TestUpdateSegmentTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	st := BuildTree(arr)
	fmt.Println(st.tree)

	st.Update(3, 11)
	fmt.Println(arr)
	fmt.Println(st.arr)
	fmt.Println(st.tree)
}

func TestQuerySegmentTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	st := BuildTree(arr)
	fmt.Println(st.tree)

	ret := st.Query(1, 5)
	fmt.Println(ret)
}
