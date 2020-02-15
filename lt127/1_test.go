package lt127

import (
	"fmt"
	"testing"
)

// 测试对于以指针访问的数组，在扩容后，这个指针是否依然指向原数组内存
func TestSliceResizeAndPointer(t *testing.T) {
	array := make([]int, 0, 3)	// cap = 3
	p := &array
	des := []int{0,0,0}
	for i:=0; i<5; i++ {
		if i == cap(array) {	// 故意模拟空间不够时系统将数据搬至另一个连续空间的情况
			*p = append(des[3:], *p...)
		}
		*p = append(*p, i)	// 到i=3时就开始扩容了
		fmt.Printf("array = %v, array cap = %d, array address = %p, p = %p, parray = %v\n", array, cap(array), &array, p, *p)
	}

	//array = [0], array cap = 3, array address = 0xc00009e080, p = 0xc00009e080, parray = [0]
	//array = [0 1], array cap = 3, array address = 0xc00009e080, p = 0xc00009e080, parray = [0 1]
	//array = [0 1 2], array cap = 3, array address = 0xc00009e080, p = 0xc00009e080, parray = [0 1 2]
	//array = [0 1 2 3], array cap = 4, array address = 0xc00009e080, p = 0xc00009e080, parray = [0 1 2 3]
	//array = [0 1 2 3 4], array cap = 8, array address = 0xc00009e080, p = 0xc00009e080, parray = [0 1 2 3 4]
}

func TestLadderLength(t *testing.T) {
	var tests = []struct{
		begin string
		end string
		words []string
		ans int
	}{
		{"hit", "cog", []string{"hot","dot","dog","lot","log"}, 0},
		{"hit", "cog", []string{"hot","dot","dog","lot","log", "cog"}, 5},
		{"hit", "cog", []string{"hot","dot","lot","log", "cog"}, 5},
	}

	for _, test := range tests {
		ret := ladderLength2(test.begin, test.end, test.words)
		if ret != test.ans {
			t.Errorf("begin = %s, end = %s, words = %s, ans = %d, but return %d\n", test.begin, test.end, test.words, test.ans, ret)
		}
	}
}
