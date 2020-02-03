package lt295

// 数据流的中位数

// 最简单直接的想法就是维护一个数据列表，记录中位数
type MedianFinder struct {
	data []int	// 数据，升序排列
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{data:make([]int,0)}
}


func (this *MedianFinder) AddNum(num int)  {
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


func (this *MedianFinder) FindMedian() float64 {
	n := len(this.data)
	if n % 2 == 0 {
		// 这里要注意避免被圆整
		return float64(this.data[n/2-1]) / 2 + float64(this.data[n/2]) / 2
	} else {
		return float64(this.data[n/2])
	}
}


// NOTICE！为了避免大量重新开辟数组空间带来的损耗，
// 可以改用有序链表加两个中位节点指针实现。


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */