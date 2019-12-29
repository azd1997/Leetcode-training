package lt74

// 搜索二维矩阵

// m*n二维矩阵
// 每行升序排列
// 每行第一个数比上一行最后一个数大

// 搜索目标数是否存在于矩阵中

// 思考
// 看到这有序区间查找元素第一反应就是二分查找了吧
// 1. 纯暴力遍历 O(m*n)/O(1)
// 2. 利用行间有序性，先线性遍历行首数字，找出目标值可能在哪一行，再线性遍历该行 O(m+n)/O(1)
// 3. 利用行间有序和行内有序， 进行两次二分搜索 O(logm+logn)/O(1) = O(logmn)/O(1)
// 4. 直接把整个二维矩阵看作一个长为mn的有序数组，进行一次二分查找 O(logmn)/O(1)

// 实现思考2,4

// 1. 解法1 （思考2）
//136/136 cases passed (8 ms)
//Your runtime beats 80.87 % of golang submissions
//Your memory usage beats 62.22 % of golang submissions (3.8 MB)
func searchMatrix(matrix [][]int, target int) bool {

	// 特殊情况  (假设每行等长)
	rows := len(matrix)
	if rows==0 {return false}
	cols := len(matrix[0])
	if cols==0 {return false}
	if matrix[0][0]>target || matrix[rows-1][cols-1]<target {return false}

	// 一般情况 经过前面筛选， target一定在矩阵范围内才可能找到
	atRow := -1
	for i:=0; i<rows; i++ {
		if matrix[i][0] < target {
			atRow++
			continue
		}
		if matrix[i][0] == target {return true}
		if matrix[i][0] > target {
			break
		}
	}
	for i:=1; i<cols; i++ {
		if matrix[atRow][i]==target {return true}
	}
	return false
}

// 2. 解法2 （思考4）
//136/136 cases passed (8 ms)
//Your runtime beats 80.87 % of golang submissions
//Your memory usage beats 62.22 % of golang submissions (3.8 MB)
func searchMatrix2(matrix [][]int, target int) bool {

	// 特殊情况  (假设每行等长)
	rows := len(matrix)
	if rows==0 {return false}
	cols := len(matrix[0])
	if cols==0 {return false}
	if matrix[0][0]>target || matrix[rows-1][cols-1]<target {return false}

	// 一般情况 经过前面筛选， target一定在矩阵范围内才可能找到。 在此区间进行一次二分查找
	// 索引 i [0, rows*cols-1] 映射到矩阵中的坐标是 (i/cols, i%cols)
	l, r, mid := 0, rows*cols-1, 0
	for l<=r {
		mid = l + (r-l)/2
		if matrix[mid/cols][mid%cols] > target {
			r = mid-1
			continue
		}
		if matrix[mid/cols][mid%cols] < target {
			l = mid+1
			continue
		}
		return true		// 等于target
	}

	return false
}

// 3. 解法2优化一下，当矩阵规模不大时直接采用顺序比较而非二分查找
//136/136 cases passed (8 ms)
//Your runtime beats 80.87 % of golang submissions
//Your memory usage beats 62.22 % of golang submissions (3.8 MB)
func searchMatrix3(matrix [][]int, target int) bool {

	// 特殊情况  (假设每行等长)
	rows := len(matrix)
	if rows==0 {return false}
	cols := len(matrix[0])
	if cols==0 {return false}
	if matrix[0][0]>target || matrix[rows-1][cols-1]<target {return false}

	// 一般情况 经过前面筛选， target一定在矩阵范围内才可能找到。 在此区间进行一次二分查找
	// 索引 i [0, rows*cols-1] 映射到矩阵中的坐标是 (i/cols, i%cols)
	l, r, mid := 0, rows*cols-1, 0
	for {
		// 顺序比较
		if r-l < 10 {	// 10是个比较小的数
			for i:=l; i<=r; i++ {
				if matrix[i/cols][i%cols] == target {return true}
			}
			return false
		}

		// 二分查找
		mid = l + (r-l)/2
		if matrix[mid/cols][mid%cols] > target {
			r = mid-1
			continue
		}
		if matrix[mid/cols][mid%cols] < target {
			l = mid+1
			continue
		}
		return true		// 等于target
	}

}