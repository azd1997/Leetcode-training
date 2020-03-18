package lt4

import "testing"

func TestSol(t *testing.T) {
	testCases := []struct {
		desc  string
		nums1 []int
		nums2 []int
		ans   float64
	}{
		{
			desc:  "1",
			nums1: []int{1, 3},
			nums2: []int{2},
			ans:   2.0,
		},
		{
			desc:  "2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			ans:   2.5,
		},
		{
			desc:  "3",
			nums1: []int{},
			nums2: []int{1, 2},
			ans:   1.5,
		},
		{
			desc:  "4",
			nums1: []int{1, 2},
			nums2: []int{-1, 3},
			ans:   1.5,
		},
		{
			desc:  "5",
			nums1: []int{2, 2, 2, 2},
			nums2: []int{2, 2, 2},
			ans:   2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ret := findMedianSortedArrays(tC.nums1, tC.nums2)
			if ret != tC.ans {
				t.Errorf("测例[%s]: ans=%f, but ret=%f\n", tC.desc, tC.ans, ret)
			}
		})
	}
}
