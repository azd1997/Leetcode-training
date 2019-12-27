package lt84

import "math"

// 柱状图中最大的矩形
// n个非负整数，表示柱状图中柱高，每个柱子宽度为1
// 求能勾勒出的矩形的最大面积

// 这题和装水容器、接雨水比较像，都是左右双指针内移的套路

// 根据题意
// 矩形有两类情况
// 1. 取heights最小值，以heights长度为宽，所得矩形面积为area1
// 2.
// 要取最大矩形面积 s=b(宽)*h(高) ，则肯定先取其中一边 b 为max，再将 b 慢慢缩小，h则尝试向增大的方向变化，才可能找到最大矩形面积
// 这里h是宽度范围里的最小值，当b缩小时，需要更新h值

// 当区间内包含 minHeight 时， h 必定为 min， 再怎么移动左右指针都不可能增大矩形面积， 所以每次都要寻找最小值
// 假设一开始heights有最小值nums[min]，其下标min，此时矩形值为 s0 = len(heights) * nums[min]
// 计算左半区间(不含min)的新最小值min1，并计算新的矩形值s1 = min * min1，
// 如果 s1 <= s0， 则左边，重复前边的操作：按min1分成新的两个区间，再去计算
// 如果 s1 > s0， 更新 maxS = s1, 并在 min1 左右分成两个区间看能不能找到更大的面积
// 从这个分析来看，适合递归求解


// 1. 自己给出的递归分治解法
// 不考虑内层寻找最小值的情况的话。外层其实类似二分查找，时间复杂度应该是O(logn)，但恶化的情况下会变成O(n)
// 而每一次findMaxArea都要O(n)时间。
// 时间O(nlogn),空间O(n)
//96/96 cases passed (664 ms)
//Your runtime beats 31.25 % of golang submissions
//Your memory usage beats 7.14 % of golang submissions (7.2 MB)
func largestRectangleArea1(heights []int) int {

	// 1. 特殊情况
	l := len(heights)
	if l==0 {return 0}
	if l==1 {return heights[0]}

	// 2. 一般情况
	var maxArea int
	findMaxArea(heights, &maxArea)
	return maxArea
}

func findMaxArea(ints []int, maxArea *int) {
	l := len(ints)
	if l==0 {return}	// 当ints没有数之时返回
	minI, minV := minInInts(ints)
	area := l * minV
	if area > *maxArea {*maxArea = area}	// 更新最大面积
	// 将区间一分为二再计算
	if minI != 0 || minI != l-1 {
		findMaxArea(ints[:minI], maxArea)
		findMaxArea(ints[minI+1:], maxArea)
	} else if minI == 0 {
		findMaxArea(ints[1:], maxArea)
	} else if minI == len(ints)-1 {
		findMaxArea(ints[:l-1], maxArea)
	}

}

// 这里是采用一次扫描的办法取最小值，但这样显然浪费了许多计算，先这样写着
func minInInts(ints []int) (index, value int) {
	value = math.MaxInt32
	for i:=0; i<len(ints); i++ {
		if ints[i] < value {
			index, value = i, ints[i]
		}
	}
	return index, value
}


// 分析前面解法，可以知道，想要找到最大矩形面积， 解法1前面的外层递归是不可避免的，总要比对这么多次数
// 时间的优化点在 minInInts 函数上想办法优化
// 可以想到： 利用栈操作来优化最小值的寻找。

// 此外，当数组本身有序的话，上面的思路分治一点优化作用也没有，因为每次都要在一个O(n)级别的数组里找最小值
// 因此最坏情况下时间复杂度变成了O(n^2)

// 从这里开始，参考官方题解。
// 1. 暴力解： 利用两个指针从左向右滑移(内外两层遍历)计算所有情况下的矩形面积， O(n^3)的时间复杂度（查找最小高度需O(n)），空间O(1)
// 2. 优化暴力 利用前一对柱子之间的最低高度来求出当前柱子间的最低高度 minHeight = min( minheight, heights[j] )  时间O(n^2)，空间O(1)
// 3. 分治解法 思路与我前面的解法一般。
// 4. 优化的分治解法 官方题解这里用线段树代替遍历寻找区间最小值，单词查询复杂度变成O(logn)，空间因为使用了线段树，占用O(n)
// 5. 使用栈


// 2. 基于栈的解法
//96/96 cases passed (12 ms)
//Your runtime beats 78.15 % of golang submissions
//Your memory usage beats 57.14 % of golang submissions (4.8 MB)
func largestRectangleArea2(heights []int) int {

	l := len(heights)

	if l == 0 {return 0}
	if l == 1 {return heights[0]}

	// 用切片模拟栈，为了节省切片扩容时间，预设切片空间为数组长度
	// 栈用来存heights数组中元素下标
	stack := make([]int, 1, l)
	stack[0] = -1		// 预填一个-1

	maxArea := 0
	var area int
	for i:=0; i<l; i++ {
		for stack[len(stack)-1] != -1 && heights[ stack[len(stack)-1] ] >= heights[i] {
			area = heights[stack[len(stack)-1]] * (i - stack[len(stack)-2] - 1)
			if area > maxArea {maxArea = area}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {		// 栈未到底
		area = heights[stack[len(stack)-1]] * (l - stack[len(stack)-2] - 1)
		if area > maxArea {maxArea = area}
		stack = stack[:len(stack)-1]
	}
	return maxArea
}