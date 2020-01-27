package lt240

// 搜索二维矩阵II

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
// 1.直接解得话，就是先从左上角比较target，若target更大，则去该点的行后部和列后部去比较
// 2. 优化暴力解的话，就是利用二分查找，快速定位到最接近就目标的区间

//
//// 1. 直接解
//func searchMatrix(matrix [][]int, target int) bool {
//	m := len(matrix)
//	if m==0 {return false}
//	n := len(matrix[0])
//	if n==0 {return false}
//	if m==1 && n==1 {return matrix[0][0]==target}
//	if target<matrix[0][0] || target>matrix[m-1][n-1] {return false}
//
//	row, col := 0, 0	// row为根据第0列定位出的行，col同理
//	for i:=0; i<m; i++ {
//		if target==matrix[i][0] {return true}
//		if target<matrix[i][0] {row = i-1; break}
//	}
//	for j:=0; j<n; j++ {
//		if target==matrix[0][j] {return true}
//		if target<matrix[0][j] {col = j-1; break}
//	}
//	for j:=0; matrix[row][]
//
//
//
//}


// 上面的思路是错的，会漏掉

// 由于每行每列都是有序，只是相对于 搜索二维矩阵，少了每行一定比上一行大的设定
// 这使得这使得本题的搜索必须遍历所有行或者列
// 以遍历所有列为例，对所有列进行遍历搜索或者二分查找(优化)

// 暴力遍历 O(mn)/O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m==0 {return false}
	n := len(matrix[0])
	if n==0 {return false}
	if target<matrix[0][0] || target>matrix[m-1][n-1] {return false}

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if matrix[i][j] == target {return true}
		}
	}
	return false
}

// 利用二分查找优化 O(mlogn)/O(logn)
func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m==0 {return false}
	n := len(matrix[0])
	if n==0 {return false}
	if target<matrix[0][0] || target>matrix[m-1][n-1] {return false}

	for i:=0; i<m; i++ {
		if binarySearch(matrix[i], target, 0, n-1) {return true}
	}
	return false
}

// 由于在矩阵那对数组长度进行了检验，这里不必继续检查数组
func binarySearch(nums []int, target int, start, end int) bool {
	if start==end {return nums[start]==target}
	mid := (start + end)/2	// 这是数组索引，一般没必要担心整溢出
	if nums[mid] == target {return true}
	if nums[mid] > target {
		return binarySearch(nums, target, start, mid)
	}
	if nums[mid] < target {
		return binarySearch(nums, target, mid+1, end)
	}
	return false
}

// 官方题解采取了按对角线遍历然后每次都对对角线上该点的行、列进行二分查找
// 本质上和我这种二分没太大区别，优化了一点点

// 3. 利用这种矩阵的特性进行搜索空间的缩减
// 特性：某一个子矩阵左上角最小值，右下角最大值
func searchMatrix3(matrix [][]int, target int) bool {
	// 见官方题解
	return false
}

// 4. 利用该矩阵的特性的特点，减少搜索空间
// 以左下角为例，其列的上部 + 该点 + 其行的右部  组成升序序列
// 因此，如果该点<target则该点右移一格；否则上移一格
// O(n+m)/O(1)
func searchMatrix4(matrix [][]int, target int) bool {
	m := len(matrix)
	if m==0 {return false}
	n := len(matrix[0])
	if n==0 {return false}
	if target<matrix[0][0] || target>matrix[m-1][n-1] {return false}

	// 起始点为矩阵左下角
	row, col := m-1, 0
	for row>=0 && col<=n-1 {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {	// 找到
			return true
		}
	}
	return false
}