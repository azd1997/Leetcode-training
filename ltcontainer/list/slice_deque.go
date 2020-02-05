package ltlist

import (
	"log"
)

type SliceDeque struct {
	data       []int
	start, end int // 实际数据的起始与结束 [start, end)
	size       int // 切片预设的开辟空间占用，
	// 不限制最终尺寸，但是预设合理的size可以减少数据搬移，提高效率
	step int
}

func (s *SliceDeque) Data() []int {
	return s.data
}

func (s *SliceDeque) Step() int {
	return s.step
}

func (s *SliceDeque) Size() int {
	return s.size
}

// 队头插入元素，首先看start是否为0，不为0前面还有空间，直接填
// 若为0，表示没空间了，将现存数据整体后移step位，再前插
func (s *SliceDeque) PushFront(v interface{}) {
	if s.start == 0 {
		s.moveWholeByStep() // 搬迁数据
		s.data[s.start-1] = v.(int)
		s.start--
	} else if s.start > 0 { // 前有余位
		s.data[s.start-1] = v.(int)
		s.start--
	} else if s.start == -1 { // 插入第一个数据
		s.data[s.start-1] = v.(int)
		s.start++
		s.end++
	}
}

func (s *SliceDeque) PushBack(v interface{}) {
	if s.start == -1 { // 插入第一个数据
		s.data[s.end] = v.(int)
		s.end++
		s.start++
	} else if s.end < s.size { // 还有位置
		s.data[s.end] = v.(int)
		s.end++
	} else { // s.end==s.size 没位置了，让data自己扩容
		s.data = append(s.data, v.(int))
		s.size = len(s.data)
		//fmt.Println(s.size)
		s.end++
	}
}

func (s *SliceDeque) PopFront() interface{} {
	if s.start == -1 {
		log.Fatalln("no elem to popfront")
	}
	res := 0
	if s.start == s.end-1 { // 只有一个元素
		res = s.data[s.start]
		s.start, s.end = -1, 0 // 恢复初始位置
	} else { // 至少有两个元素
		res = s.data[s.start]
		s.start++
	}
	return res
}

func (s *SliceDeque) PopBack() interface{} {
	if s.start == -1 {
		log.Fatalln("no elem to popback")
	}
	res := 0
	if s.start == s.end-1 { // 只有一个元素
		res = s.data[s.start]
		s.start, s.end = -1, 0 // 恢复初始位置
	} else { // 至少有两个元素
		res = s.data[s.end-1]
		s.end--
	}
	return res
}

func (s *SliceDeque) SeekFront() interface{} {
	if s.start == -1 {
		log.Fatalln("no elem to popfront")
	}
	return s.data[s.start]
}

func (s *SliceDeque) SeekBack() interface{} {
	if s.start == -1 {
		log.Fatalln("no elem to popfront")
	}
	return s.data[s.end-1]
}

func (s *SliceDeque) IsEmpty() bool {
	return s.start==-1
}

func (s *SliceDeque) moveWholeByStep() {
	if s.end+s.step <= s.size {
		// 只将数据部分拷走，后序遍历
		for i := s.end; i > s.start; i-- {
			s.data[i+s.step] = s.data[i]
		}
		s.start, s.end = s.start+s.step, s.end+s.step
	} else { // 空间不够，重新分配内存
		space := make([]int, s.step)
		newdata := append(space, s.data...)
		//fmt.Println("newdata", newdata)
		s.data = newdata
		s.start, s.end = s.start+s.step, s.end+s.step
		s.size += s.step
	}
}

func NewSliceDeque(size int) *SliceDeque {
	return &SliceDeque{
		data:  make([]int, size),
		start: -1,
		end:   0,
		size:  size,
		step:  5,
	}
}
