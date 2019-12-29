package lt85

import (
	"fmt"
	"math/big"
	"strconv"
)

// 最大矩形
// 在仅包含 '0' / '1' 的二维二进制矩阵中找出只包含 1 的最大矩形

// 输入:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//输出: 6
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximal-rectangle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 第一反应就是： 从底往上，计算以每一行为x轴得到的直方图中的最大矩形，最后汇总得到真正的最大矩形

// 以上面例子为例，这个矩阵表示为 a[][]
// 最后一行 a[3] 只有 a[3][0]和a[3][3] = 1， 而且此时这两个不相邻
// 所以此时把这两个看作是两个直方图求最大矩形的子问题 （当然，由于单列太简单，直接计算就好）
// 那么，什么叫能表示为直方图呢？ 从当前值，比如说 a[3][0]开始往上，直到搜索到0，就找到能坍缩的直方图柱高
// 在这里 a[3][0]所代表的子问题中最大面积 maxArea = 4; a[3][3]中为3，所以不更新

// 最后一行处理完了，再来处理倒数第二行，遍历发现这行所有值均为1，也就是全相邻，直接将这行以及这行往上的所有内容转化为直方图问题即可，
// 而且不用继续行迭代。再网上迭代也不可能找到更大的矩形
// 在这行中按照 lt84直方图中最大矩形面积 计算就好

// 在这个思路中，非常重要的一环是判断 '1'的连续性，无论是行中连续性还是列中连续性
// 列中连续性直接向上遍历
// 行中连续性则向右遍历至'0'停，继续遍历至'1'又开始

// 代码实现如下

// 1. 自己解法，思路如上
// 总体时空复杂度均为 O(mn)
//66/66 cases passed (4 ms)
//Your runtime beats 83.74 % of golang submissions
//Your memory usage beats 60 % of golang submissions (5.7 MB)
func maximalRectangle(matrix [][]byte) int {

	// 行列数
	rows := len(matrix)
	if rows==0 {return 0}
	cols := len(matrix[0])	// 这里假设二维切片是等宽的
	if cols==0 {return 0}

	// 先遍历一次matrix，将之转化为直方图形式。 时间O(rows*cols), 空间(rows*cols)
	matrix1 := make([][]int, rows)
	matrix1[0] = make([]int, cols)
	for c:=0; c<cols; c++ {
		if matrix[0][c] == '1' {matrix1[0][c] = 1}
	}
	for r:=1; r<rows; r++ {
		matrix1[r] = make([]int, cols)
		for c:=0; c<cols; c++ {
			if matrix[r][c] == '1' {
				matrix1[r][c] = matrix1[r-1][c] + 1
			} else {
				matrix1[r][c] = 0
			}
		}
	}

	fmt.Printf("matrix1: \n%v\n", matrix1)

	// 按行从下往上遍历，若遇全连续(allOne=true)则不继续向上遍历
	//var allOne bool
	// 这里纠正上面说的遇全连续则不继续向上遍历，这样是有可能丢解的，必须全试
	// 比如 [[1] [1] [0] [1]]
	var maxArea, area int

	// 时间复杂度 外层遍历需要 O(rows*cols)
	// 在假设但是在计算面积时需要调用时间复杂度为O(n)空间复杂度为O(n)的子函数
	// 这个调用当且仅当连续2个格子'1'才会发生，所以极端情况下，cols三分(110110110...)
	// 这种情况下调用频率是 rows*cols* 1/3 * O(n)
	// 而这里的 n 是连续'1'的个数， 可以认为 每一行所有的子函数调用总体时空复杂度是O(cols),而且相当于在遍历行中元素时被均摊了
	// 因此，这个部分时间复杂度 O(rows*cols)
	for r:=rows-1; r>=0; r-- {

		//fmt.Printf("maxArea = %d, r = %d\n", maxArea, r)

		// 若前一行遍历时allOne=true则没必要继续遍历了
		//if allOne {break}

		// 遍历该行
		start, end := -1, -1	// 连续'1'的起始与结束
		for c:=0; c<cols; c++ {

			fmt.Printf("start=%d, end=%d, r=%d, c=%d, area=%d, maxArea=%d\n", start, end, r, c, area, maxArea)

			// 遇'0'继续，直至找到下一个'1'来作为start
			if matrix[r][c] == '0' {continue}
			//
			if matrix[r][c] == '1' {
				if start == -1 {start = c}
				if c+1 == cols || (c < cols-1 && matrix[r][c+1] == '0') {end = c}
				// start和end都统计出来后，要计算该子问题下最大面积并更新全局maxArea
				if start != -1 && end != -1 {
					// 单列时直接计算更新
					if start == end {
						if matrix1[r][start] > maxArea {maxArea = matrix1[r][start]}
					} else {
						// 否则按直方图计算
						area = largestRectangleAreaInLt84(matrix1[r][start:end+1])
						if area > maxArea {maxArea = area}
					}

					// 若全1则标志
					//if start==0 && end==cols-1 {allOne = true}

					// 计算后将start和end重置
					start, end = -1, -1
				}
			}
		}
	}

	return maxArea
}

// 直方图求解
// 这里不需要做什么检查工作
func largestRectangleAreaInLt84(heights []int) int {
	l := len(heights)

	// 构建leftFirstLess
	var p int
	leftFirstLess := make([]int, l)
	leftFirstLess[0] = -1
	for i:=1; i<l; i++ {
		p = i-1
		for p>=0 && heights[p] >= heights[i] {
			p = leftFirstLess[p]	// p 快速左移
		}
		leftFirstLess[i] = p
	}
	// 构建rightFirstLess
	rightFirstLess := make([]int, l)
	rightFirstLess[l-1] = l
	for i:=l-2; i>=0; i-- {
		p = i + 1
		for p<=l-1 && heights[p] >= heights[i] {
			p = rightFirstLess[p]
		}
		rightFirstLess[i] = p
	}

	// 求包含每个柱子的矩形区域的最大面积，从中选出最大者
	var maxArea, area int
	for i:=0; i<l; i++ {
		area = (rightFirstLess[i] - leftFirstLess[i] - 1) * heights[i]
		if area > maxArea {maxArea = area}
	}
	return maxArea
}


// 好，自己上来就解的解法可行，并且结果还不错，接下来看看题解

// 2.
// 在某解答中看到思路和解法1一致的，只不过，其每行直接调用 largestRectangleAreaInLt84 而不是像我这般
// 对行中元素做了分段，从代码上更容易理解，这里再改写一次
func maximalRectangle2(matrix [][]byte) int {

	// 行列数
	rows := len(matrix)
	if rows==0 {return 0}
	cols := len(matrix[0])	// 这里假设二维切片是等宽的
	if cols==0 {return 0}

	// 先遍历一次matrix，将之转化为直方图形式。 时间O(rows*cols), 空间(rows*cols)
	matrix1 := make([][]int, rows)
	matrix1[0] = make([]int, cols)
	for c:=0; c<cols; c++ {
		if matrix[0][c] == '1' {matrix1[0][c] = 1}
	}
	for r:=1; r<rows; r++ {
		matrix1[r] = make([]int, cols)
		for c:=0; c<cols; c++ {
			if matrix[r][c] == '1' {
				matrix1[r][c] = matrix1[r-1][c] + 1
			} else {
				matrix1[r][c] = 0
			}
		}
	}

	// 按行从下往上遍历(从上往下也没问题)
	var maxArea, area int
	for r:=rows-1; r>=0; r-- {
		area = largestRectangleAreaInLt84(matrix1[r])
		if area > maxArea {maxArea = area}
	}

	return maxArea
}


// 3. 位运算解法 —— 暴力解优化
// 未能通过。
// [["0","1","1","1","1","1","0","1","1","1","1","1","0","1","0","1","1","0","0","1","1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","0","0","0","1","1","0","1","1","1","0","1","1","1","1","0","0","1","1","1","0","0","0","1","1","1","0","1","1","1","1","1","1","1","1","1","1","1","1","1","0","0","1","1","1","1","0","1","1","1","1","1","1","1","1","0","0","1","0","1","0","1","0"]]
// 由于用int64表示，位数有限制
func maximalRectangle3(matrix [][]byte) int {
	// 行列数
	rows := len(matrix)
	if rows==0 {return 0}
	cols := len(matrix[0])	// 这里假设二维切片是等宽的
	if cols==0 {return 0}

	// 将每行转化为二进制数
	nums := make([]int64, rows)
	for i:=0; i<rows; i++ {
		//nums[i], _ = binary.ReadUvarint(bytes.NewReader(matrix[i]))	这个不行
		nums[i], _ = strconv.ParseInt(string(matrix[i]), 2, 64)
	}
	//fmt.Printf("nums=%v\n", nums)

	var maxArea, area, j, num int

	// 遍历所有行
	for i:=0; i<rows; i++ {
		j, num = i, int(nums[i])
		// 将第i行 连续的 和接下来的所有行 作与运算
		for j < rows {
			// 与运算之后，num化为二进制中的 1， 表示从第i到第j行，可以组成矩形的那几列
			num = num & int(nums[j])

			if num == 0 {break}
			l, curnum := 0, num
			// 每次循环将curnum与 其左移一位 的数 作与运算
			// 最终的循环次数 l 表示最宽的矩形宽度
			for curnum != 0 {
				l++
				curnum = curnum & (curnum << 1)
			}
			area = l * (j-i+1)
			if area > maxArea {maxArea = area}

			j++
		}
	}

	return maxArea
}

// 4. 解法三的改进版
// 使用大数表示
func maximalRectangle4(matrix [][]byte) int {
	// 行列数
	rows := len(matrix)
	if rows==0 {return 0}
	cols := len(matrix[0])	// 这里假设二维切片是等宽的
	if cols==0 {return 0}

	// 将每行转化为二进制数
	nums := make([]big.Int, rows)
	for i:=0; i<rows; i++ {
		nums[i] = big.Int{}
		nums[i].SetString(string(matrix[i]), 2)		// 不能使用 SetBytes()	我们要位操作，所以得将字符表示成二进制位先
	}
	fmt.Printf("nums=%v\n", nums)

	var maxArea, area, j int
	var num big.Int

	// 遍历所有行
	for i:=0; i<rows; i++ {
		j, num = i, nums[i]
		//fmt.Printf("num addr = %p, nums[i] addr = %p\n", &num, &nums[i])
		// 将第i行 连续的 和接下来的所有行 作与运算
		for j < rows {
			// 与运算之后，num化为二进制中的 1， 表示从第i到第j行，可以组成矩形的那几列

			num.And(&num, &nums[j])

			fmt.Printf("j=%d, num=%v\n", j, &num)

			if num.Cmp(big.NewInt(0)) == 0 {break}

			bits1, bits2 := make([]big.Word, len(num.Bits())), make([]big.Word, len(num.Bits()))
			l, curnum, curnumLsh := 0, new(big.Int).SetBits(bits1), new(big.Int).SetBits(bits2)
			//fmt.Printf("num addr = %p, curnum addr = %p\n", &num, &curnum)
			// 每次循环将curnum与 其左移一位 的数 作与运算
			// 最终的循环次数 l 表示最宽的矩形宽度
			for curnum.Cmp(big.NewInt(0)) != 0 {
				l++
				curnumLsh = curnum
				fmt.Printf("num = %v, curnum = %v, curnumLsh = %v\n", &num, curnum, curnumLsh)
				curnumLsh.Lsh(curnumLsh, 1)
				fmt.Printf("num = %v, curnum = %v, curnumLsh = %v\n", &num, curnum, curnumLsh)

				curnum.And(curnum, curnumLsh)
				//fmt.Printf("num addr = %p, curnum addr = %p\n", &num, &curnum)
				fmt.Printf("num = %v, curnum = %v, curnumLsh = %v\n", &num, &curnum, &curnumLsh)
				//fmt.Printf("l=%d, curnum=%v\n", l, &curnum)
			}
			area = l * (j-i+1)
			if area > maxArea {maxArea = area}
			//fmt.Printf("num = %v, area = %d, maxArea = %d\n", num, area, maxArea)
			//fmt.Printf("nums = %v\n", nums)
			j++
		}
	}

	return maxArea
}