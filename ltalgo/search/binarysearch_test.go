package search

import (
	"testing"
	"fmt"
)

func TestBS(t *testing.T) {
	nums := []int{1,3,5,7,9}
	target := 3
	existed := binarySearch2(nums, target)
	fmt.Println(existed)
}
