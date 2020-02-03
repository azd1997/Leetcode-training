package lt1342

import "sort"

// 数组大小减半

// 根据题意显然会想到利用贪心的思路
// 先统计各个元素的出现次数，按次数递减排序，看前边最少几个次数加一起过半

func minSetSize(arr []int) int {
	// 哈希集和记录元素出现次数
	m := make(map[int]int)
	for i:=0; i<len(arr); i++ {m[arr[i]]++}
	// 倒入数组进行排序
	arr2 := make([]int, 0, len(m))
	for _, v := range m {arr2 = append(arr2, v)}
	sort.Ints(arr2)
	count, sum := 0, 0
	for i:=len(arr2)-1; i>=0; i-- {
		sum += arr2[i]
		count++
		if sum >= len(arr)/2 {break}	// 题目说了arr长度为偶数，所以这里直接除2，如果是奇数要注意先加1再除2
	}
	return count
}
