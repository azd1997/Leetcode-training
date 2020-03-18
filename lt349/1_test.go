package lt349

import (
	"fmt"
	"testing"
)

func TestSol3(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	ret := intersection3(nums1, nums2)
	fmt.Println(ret)
}
