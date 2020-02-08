package lt179

import (
	"sort"
	"strconv"
	"strings"
)

// 最大数

// 要使重排后的数值最大，那显然num同一位上数值越大，越应该在字符串拼接时排在最前面
// 其实就是一个排序问题

func largestNumber(nums []int) string {

	// 1. 将所有nums转化为字符串
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}

	// 2. 将numStrs按每位数字字节大小排序(数字大小关系等于于字节大小关系)
	sort.Slice(numStrs, func(i, j int) bool {
		return ALargerThanB(numStrs[i], numStrs[j])
	})
	// sort.Strings(numStrs) // 没太搞懂StringSlice用的less (return p[i]<p[j])
	// 字符串之间是怎么比较大小的，如果是从前向后按byte比，那么这里可以用；否则不能

	// 3. 需要注意的是，如果排序后最大位置的都是"0"，那说明全是"0"
	// 直接返回"0"
	if numStrs[0] == "0" {return "0"}

	// 4. 将numStrs合并成字符串
	return strings.Join(numStrs, "")
}

// 备注：像我下面这个思路去比较的话，
// 当出现连续相等之后可能需要不停的去首尾交替的比较，太难处理
// 比较两个数字的字符串，如果a>b，返回true
//func ALargerThanB(a, b string) bool {
//	la, lb := len(a), len(b)
//
//	if la == lb {
//		for i := 0; i < la; i++ {
//			if a[i] > b[i] {return true}
//			if a[i] < b[i] {return false}
//		}
//		return true		// 等长等大，谁前谁后无所谓
//	}
//
//	if la > lb {
//		x, y := la/lb, la%lb
//		for k:=0; k<x; k++ {
//			for i:=0; i<lb; i++ {	// 注意a的下标别出界
//				if a[i+k*lb] > b[i] {return true}
//				if a[i+k*lb] < b[i] {return false}
//			}
//		}
//		for i:=0; i<y; i++ {
//			if a[i+k*lb] > b[i] {return true}
//			if a[i+k*lb] < b[i] {return false}
//		}
//
//		ia := 0		// a的每一小段的起始下标
//		for ia < la {
//			for i:=0; i<lb && i+ia<la; i++ {	// 注意a的下标别出界
//				if a[i+ia] > b[i] {return true}
//				if a[i+ia] < b[i] {return false}
//			}
//			ia += lb
//		}
//	}
//
//
//
//
//
//
//
//
//
//
//
//	idx := 0	// 长的那一个字符串的索引
//	for idx < l {
//		for i := 0; i < l; i++ {
//			if a[i] > b[i] {
//				return true
//			}
//			if a[i] < b[i] {
//				return false
//			}
//		}
//	}
//
//
//	// [121,12] 12 121 > 121 12
//	// 如果 a[lb]==a[0] 这种时候还需要继续比较，怎么比较呢？
//	// 循环比较
//	// 例如 12 1212121(注意最后的1)
//	// 把 长 按 短的长度进行分割 跟短进行比较
//
//	// 如果共同长度上都相等，那么谁大谁小取决于长的那个数字的下一个位与开头位的比较
//	if la==lb {		// 二者长度相同且完全一致，谁前谁后无所谓
//		return true
//	} else if la > lb {
//		return a[lb] > a[0]
//	} else {return b[0] > b[la]}
//}


// 其实应该换个思路，不是比较a>b而是比较 ab>ba，这样事情就简单许多了
func ALargerThanB(a, b string) bool {
	// return strings.Compare(a+b, b+a) >= 0
	return a+b >= b+a		// 内建的>=<比Compare更快
}