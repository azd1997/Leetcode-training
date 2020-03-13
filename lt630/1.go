package lt630

import (
	"container/heap"
	"math"
	"sort"
)

// 课程表III

//https://leetcode-cn.com/problems/course-schedule-iii/solution/ke-cheng-biao-iii-by-leetcode/

// 贪心是动态规划问题中的特例，本题可以应用贪心思想

// 根据贪心思想，对于结束时间d不同的的课程，先学习早结束的那个　总是更优的选择。
// 而当d相同时，先学习持续时间短的那个则是更优的选择。

// 因此，可以直接基于这两点进行排序

// 贪心思想　+ 排序 O(n2)/O(n)
func scheduleCourse(courses [][]int) int {

	n := len(courses)

	// 1. 先按d后按t排序
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1] // 早结束者排前面
	})

	// 2. 其实很像是动态规划，ts就是dp状态表
	ts := make([]int, n+1)   // ts[i]表示学完i门课最短需要的时间，初始化为MaxInt32
	max := 0                 // 最多可以学完的课程数
	for i := 0; i < n; i++ { // 第0,1,...,n-1门课
		t, d := courses[i][0], courses[i][1]
		ts[i+1] = math.MaxInt32
		for j := i; j >= 0; j-- { // 倒着来，避免影响前面的而对后面造成影响
			tmp := ts[j] + t               // 注意是j门课。加上当前这门课后的总用时
			if tmp <= d && tmp < ts[j+1] { // 使用课程i替换j+1门课程中其中一门的话总耗时更小
				ts[j+1] = tmp
				if max < j+1 {
					max = j + 1
				}
			}
		}
		// 此时ts[j+1]已更新为学完j+1门课所需的最短时间
	}
	return max
}

// 上面其实是用了线性遍历的方法来更新，其实是比较慢的
// 可以使用堆（或者说优先队列）来加速这一过程

// 贪心思想 + 优先队列
func scheduleCourse2(courses [][]int) int {

	n := len(courses)

	// 1. 先按d后按t排序
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1] // 早结束者排前面
	})

	// 2. 优先队列，这里使用一个堆。只存课程时间t
	th := new(tHeap)

	time := 0 // 总时间
	for i := 0; i < n; i++ {
		t, d := courses[i][0], courses[i][1]
		if time+t <= d { // 时间足够则time叠加上去，并将t入堆
			heap.Push(th, t)
			time += t
		} else if th.Len() != 0 && th.Peek() > t {
			// 堆非空，且堆顶的课程时间 > 当前这门课的时间，则交换
			time += t - heap.Pop(th).(int)
			heap.Push(th, t)
		}
	}

	return th.Len() // 最后堆中的课程数就代表了能够上完的最大课程
}

// tHeap，实现heap.Interface  根据课程时间作为比较的堆，时间最大在堆顶
type tHeap []int

func (h *tHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j] // 大顶堆
}

func (h *tHeap) Len() int {
	return len(*h)
}

func (h *tHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *tHeap) Push(t interface{}) {
	*h = append(*h, t.(int))
}

// 这里提供的Pop函数是将数组尾部的元素删除
func (h *tHeap) Pop() (t interface{}) {
	*h, t = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *tHeap) Peek() (t int) {
	return (*h)[0]
}

///////////////////////////////////////////////

// 由于本题不要求课程顺序，也就与课程ID无关，
// 因此前面在排序与入堆的时候都没有考虑记录课程ID信息
// 如果需要记录课程ID信息，需要使用索引化的排序方式以及优先队列或者是索引最大堆
