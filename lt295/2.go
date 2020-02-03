package lt295

import (
	"container/heap"
	LTheap "github.com/azd1997/Leetcode-training/ltcontainer/heap"
)

// 显然这样的要求更适合使用堆来实现。使用数组实现的插入时间复杂度在O(n)，
// 使用堆则为logn

// 使用小顶堆，存数据流最大的n/2个数，小顶堆的大小维持在n/2，
// 当n为偶数则返回小顶堆中最小的两个数
type MedianFinder2 struct {
	heap LTheap.IntHeap
}


/** initialize your data structure here. */
func Constructor2() MedianFinder2 {
	intheap := LTheap.IntHeap{}
	heap.Init(&intheap)
}


func (this *MedianFinder2) AddNum(num int)  {
	n := len(this.data)

	if n==0 {
		// 当前data还是空数组
		this.data = append(this.data, num)
	} else if num <= this.data[0] {
		// 新数num比data[i]都小
		this.data = append([]int{num}, this.data...)
	} else if num >= this.data[n-1] {
		// 新数num比data[i]都大
		this.data = append(this.data, num)
	} else {
		// 新数在data数组中间，需要遍历找到插入位置
		idx := 0
		for i:=1; i<n; i++ {
			if num>=this.data[i-1] && num<=this.data[i] {idx = i}
		}
		tmp := make([]int, n+1)
		for i:=0; i<idx; i++ {
			tmp[i] = this.data[i]
		}
		for i:=idx+1; i<=n; i++ {
			tmp[i] = this.data[i-1]
		}
		tmp[idx] = num
		this.data = tmp
	}
}


func (this *MedianFinder2) FindMedian() float64 {
	n := len(this.data)
	if n % 2 == 0 {
		// 这里要注意避免被圆整
		return float64(this.data[n/2-1]) / 2 + float64(this.data[n/2]) / 2
	} else {
		return float64(this.data[n/2])
	}
}
