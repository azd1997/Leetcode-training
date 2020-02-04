package lt347

import (
	"container/heap"
	"sort"
)

// 前k个高频元素


// 1. 最简单直接的做法是统计次数、按次数排序
// 错误答案
// 错在使用m2构建逆向映射。但是很显然元素出现次数有可能相等
func topKFrequent11(nums []int, k int) []int {

	// 统计各元素次数
	n := len(nums)
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[nums[i]]++
	}
	// 构建次数->元素的映射
	m2 := make(map[int]int)
	for k, v := range m {
		m2[v] = k
	}

	// 将元素次数信息倒入counts，并排序
	n = len(m)
	counts := make([]int, 0, n)
	for _, v := range m {
		counts = append(counts, v)
	}
	sort.Ints(counts)

	// 根据m2把对应的元素拿到
	res := make([]int, k)
	for i:=0; i<k; i++ {
		res[i] = m2[counts[n-1-i]]
	}
	return res
}


// 1. 最简单直接的做法是统计次数、按次数排序
// O(nlgn)/O(3n)
// 这依然是错的，因为go quickSort源码中并不是完全依赖Less而后Swap的
// 所以这样写less函数，只有当数组长度小于12时才有用，因为那时调用的是插入排序
func topKFrequent12(nums []int, k int) []int {

	// 统计各元素次数
	n := len(nums)
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[nums[i]]++
	}

	// 将去重后元素信息、次数信息分别倒入nums2, counts
	n = len(m)
	nums2, counts := make([]int, 0, n), make([]int, 0, n)
	for key, v := range m {
		nums2 = append(nums2, key)
		counts = append(counts, v)
	}
	sort.Slice(nums2, func(i, j int) bool {
		// 这个比较函数必须同步调换counts（这里要添加）和nums2(由sort.Slice负责)
		if counts[i] > counts[j] {
			counts[i], counts[j] = counts[j], counts[i]
			return true		// 使得高频元素排列在前
		}
		return false
	})

	return nums2[:k]
}



// 1. 最简单直接的做法是统计次数、按次数排序
// O(nlgn)/O(3n+k)
func topKFrequent13(nums []int, k int) []int {

	// 统计各元素次数
	n := len(nums)
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[nums[i]]++
	}

	// 将去重后元素信息、次数信息倒入helper
	n = len(m)
	helper := make([][2]int, 0, n)
	for key, v := range m {
		helper = append(helper, [2]int{key, v})
	}
	sort.Slice(helper, func(i, j int) bool {
		return helper[i][1] > helper[j][1]	// 降序排序
	})

	// 导出结果
	res := make([]int, k)
	for i:=0; i<k; i++ {
		res[i] = helper[i][0]
	}

	return res
}

// 1. 最简单直接的做法是统计次数、按次数排序
// 相比于解法14，有一些改良
// 仔细想一下，这个解法是错的，虽然能通过测试，但是它最后
// 返回的结果数组内没有任何大小顺序之分
// O(nlgn)/O(3n+k)
func topKFrequent14(nums []int, k int) []int {

	// 统计各元素次数
	n := len(nums)
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[nums[i]]++
	}

	// 将去重后元素信息、次数信息倒入helper
	n = len(m)
	counts := make([]int, 0, n)
	for _, v := range m {
		counts = append(counts, v)
	}
	sort.Ints(counts)

	//
	minK := counts[n-k]

	// 遍历哈希表，导出结果
	res := make([]int, 0, k)
	for k, v := range m {
		if v>=minK {res = append(res, k)}
	}

	return res
}

// 当然题目要求时间复杂度优于O(nlogn)，所以上面的解法是不合格的

// 求前K这种题一般都少不了堆
// 而这道是按频次排序，也就是需要优先队列

// 2. 优先队列  O(nlgn)/O(3n+k)
func topKFrequent2(nums []int, k int) []int {

	// 统计各元素次数
	n := len(nums)
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[nums[i]]++
	}

	// 将去重后元素信息、次数信息倒入优先队列中
	n = len(m)
	pq := make(PriorityQueue, n)
	i := 0
	for key, v := range m {
		pq[i] = &Item{key, v, i}
		i++
	}
	// 优先队列堆化
	heap.Init(&pq)

	// 导出结果
	res := make([]int, k)
	for i:=0; i<k; i++ {
		res[i] = heap.Pop(&pq).(*Item).value
	}

	return res
}


//===========优先队列=================

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}


//===================================


// 3. 其实借助堆排序可以实现O(nlgk)，这样就满足题意了
// 使用大小为K+1的优先队列(最小堆，优先度低的在堆顶)		// 其实有很多种变化啦，看自己喜欢
func topKFrequent3(nums []int, k int) []int {

	// 统计各元素次数
	n := len(nums)
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[nums[i]]++
	}

	// 将去重后元素信息、次数信息倒入优先队列中
	n = len(m)
	pq := make(PriorityQueue2, 0, k+1)
	heap.Init(&pq)
	for key, v := range m {
		heap.Push(&pq, &Item2{key, v})
		if pq.Len()>k {		// 弹出堆顶
			heap.Pop(&pq)
		}
	}

	// 导出结果
	res := make([]int, k)
	for i:=k-1; i>=0; i-- {
		res[i] = heap.Pop(&pq).(*Item2).value
	}

	return res
}


//===========优先队列=================

// An Item is something we manage in a priority queue.
type Item2 struct {
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue2 []*Item2

func (pq PriorityQueue2) Len() int { return len(pq) }

func (pq PriorityQueue2) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority		// 注意这里
}

func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue2) Push(x interface{}) {
	item := x.(*Item2)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}


//===================================


