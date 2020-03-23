package lt1383

import (
	"container/heap"
	"sort"
)

// 单周赛180 t4
// 最大的团队表现值

// 这道问题的核心就是排序 ，TopK排序
// 由于排序依赖 s(speed)和e(efficiency)
// 先控制其中一项已排好序，再比较另一项引入带来的变化: **每个人作为最低效率时，在其左侧找到至多K个最大速度**
// 总体的算法思路是： 预排序 + 优先队列(堆)
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	// 按效率预排序
	// 为了固定工程师ID，需要将speed,efficiency合并再去排序
	nums := make([][2]int, n)
	for i := 0; i < n; i++ {
		nums[i] = [2]int{speed[i], efficiency[i]}
	}
	sort.Slice(nums, func(i, j int) bool { // 效率高者排前面，这样对于nums[i]来说，他肯定是[0:i]的最低效率者，然后在其中找最多K个最大速度者
		return nums[i][1] > nums[j][1]
	})

	// 维护一个优先队列
	pq := newPQ()

	// 维护所选的工程师速度总和sum，以及待返回的最大表现值res
	sum, res, tmp := 0, 0, 0

	// 遍历效率
	for i := 0; i < n; i++ {
		// 1. 速度入队
		heap.Push(pq, nums[i][0])
		// 2. 计算sum。 如果优先队列size>k出掉队列顶部
		sum += nums[i][0]
		if pq.Len() > k {
			sum -= heap.Pop(pq).(int)
		}
		// 3. 更新最大表现值
		tmp = sum * nums[i][1]
		if tmp > res {
			res = tmp
		}
	}

	return res % (1e9 + 7) // 注意只能在最后模，不能前面的迭代过程中模，因为需要比较
}

//////////////////////////////////////////
// 优先队列 （这里是大小为k的小顶堆）

// 存速度就好，反正题目不要求给出是选哪几个工程师
type priorityQueue []int

func newPQ() *priorityQueue {
	arr := make([]int, 0)
	pq := priorityQueue(arr)
	return &pq
}

func (q *priorityQueue) Len() int { return len(*q) }
func (q *priorityQueue) Less(i, j int) bool {
	return (*q)[i] < (*q)[j]
}
func (q *priorityQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}
func (q *priorityQueue) Push(v interface{}) {
	*q = append(*q, v.(int))
}
func (q *priorityQueue) Pop() (v interface{}) {
	*q, v = (*q)[:len(*q)-1], (*q)[len(*q)-1]
	return v
}
func (q *priorityQueue) Peek() int {
	return (*q)[0]
}
