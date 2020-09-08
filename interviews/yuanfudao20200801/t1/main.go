package main

import (
	"fmt"
	"sort"
)

func main() {
	N := 0
	fmt.Scan(&N)

	courses := make([]Course, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&courses[i].Start, &courses[i].End)
	}

	//fmt.Println(courses)

	ans := sol(courses)

	fmt.Println(ans)
}

type Course struct {
	Start int
	End int
}

// 输入
//4
//1 4
//1 2
//2 3
//3 4
// 输出
// 2


// 输入课程列表，得出最少需要同时修几门课，才能把所有课修完
// 其实没有任何优化点，要修完所有课，就是看最多有几门课同时存在
// 但是暴力解的话会是 O( n * Latest)
// 另一种思路是先排序，然后尽可能在一轮安排尽可能多的课程，同时在所有课程中减去其，再进行下一轮，直至某一轮再无课程需要修
// 轮数就是k
func sol(courses []Course) int {
	// 按开始时间排序，并且开始时间相同的结束时间早的排前面
	sort.Slice(courses, func(i, j int) bool {
		if courses[i].Start == courses[i].Start {
			return courses[i].End < courses[i].End
		}
		return courses[i].Start < courses[i].Start
	})

	k := 0
	used := make(map[Course]bool)
	n := len(courses)	// 课程数
	lastCourseEnd := 0
	for n > 0 {		// 只要课程表中还有课程，就继续循环
		//fmt.Println(k, used, n, lastCourseEnd, courses)
		for i:=0; i<len(courses); i++ {
			if !used[courses[i]] && courses[i].Start >= lastCourseEnd {	// 当前课程未使用 并且这门课开始晚于上一门选中的课的结束
				used[courses[i]] = true
				lastCourseEnd = courses[i].End
				n--		// 使用掉该门课
			}
		}
		lastCourseEnd = 0	// 一轮之后重置
		k++
	}
	return k
}