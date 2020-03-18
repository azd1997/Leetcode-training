package lt836

import "math"

// 矩形重叠

// 1. 从数学的角度可以直接将两个矩形的边表示成线段，检查交点情况，来判断；
// 2. 从重叠的特性来看，必然有一个矩形的角是落在对方矩形内部的
// 	  因此可以检查矩形的四个角的坐标是否落在另一个矩形的内部

// 基于上面思路的第二点进行求解

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// rec1的四个角
	corners := fourCornerPoints(rec1)
	for _, c := range corners {
		if isPointInRec(c, rec2) {
			return true
		}
	}
	// rec2的四个角
	corners = fourCornerPoints(rec2)
	for _, c := range corners {
		if isPointInRec(c, rec1) {
			return true
		}
	}

	// 都没落在另一个矩形内
	return false
}

// 判断一个点是否在一个矩形内部
func isPointInRec(p [2]int, rec []int) bool {
	return p[0] > rec[0] && p[0] < rec[2] && p[1] > rec[1] && p[1] < rec[3]
}

func fourCornerPoints(rec []int) [][2]int {
	ret := make([][2]int, 4)
	ret[0] = [2]int{rec[0], rec[1]}
	ret[1] = [2]int{rec[2], rec[3]}
	ret[2] = [2]int{rec[0], rec[3]}
	ret[3] = [2]int{rec[1], rec[2]}
	return ret
}

// 很可惜上面这个解法是错误的，它漏掉了两个矩形有一个角重合的情况，例如
// [2 18 4 20] 和 [3 8 4 20]

// 官方题解给出了两种方法：
// 1. 排除法：排除不可能相交的情况：
// 一个矩形在另一个矩形右边/左边/上边/下边
// 右边： rec1[0] >= rec2[2]
// 左边： rec1[2] <= rec2[0]
// 上边： rec1[1] >= rec2[3]
// 下边： rec1[3] <= rec2[1]
// 2. 区域重叠：二维平面的两个矩形相交，则其投影到两个坐标轴后依然有重叠，而且是每个坐标轴都有重叠
// 投影到x轴： rec1 rec1[0]-rec1[2] ; rec2 rec2[0]-rec2[2]，
// 要有重叠则是 min(rec1[2], rec2[2]) > max(rec1[0], rec2[0])
// 投影到y轴同理

// 排除法
func isRectangleOverlap2(rec1 []int, rec2 []int) bool {
	if rec1[0] >= rec2[2] || rec1[2] <= rec2[0] ||
		rec1[1] >= rec2[3] || rec1[3] <= rec2[1] {
		return false
	}
	return true
}

// 区域重叠
func isRectangleOverlap3(rec1 []int, rec2 []int) bool {
	return (math.Min(float64(rec1[2]), float64(rec2[2])) > math.Max(float64(rec1[0]), float64(rec2[0]))) &&
		(math.Min(float64(rec1[3]), float64(rec2[3])) > math.Max(float64(rec1[1]), float64(rec2[1])))
}
