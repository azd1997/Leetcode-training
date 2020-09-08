package lt00000

import (
	"math"
)

// 统计好三元组

// 思路：根据前两项条件进行回溯，得到全部排列，再根据第三项条件进行筛选


func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	path := make([]int, 0, 3)	// 选出的三元组(i,j,k)，注意只保存下标，而非值
	dfs(arr, a, b, c, path, &res)
	return res
}

func dfs(arr []int, a int, b int, c int, path []int, res *int) {
	// 终止
	if len(path) == cap(path) {
		// 检查是否满足第三项条件
		if int(math.Abs(float64(arr[path[0]] - arr[path[2]]))) <= c {
			//fmt.Printf("(arr[%d], arr[%d], arr[%d]) = (%d,%d,%d)\n",
			//	path[0], path[1], path[2], arr[path[0]], arr[path[1]], arr[path[2]])
			*res ++
		}
		return
	}

	// 选择
	if len(path) == 2 {	//准备选择第三个数
		for x:=path[1]+1; x < len(arr); x++ {
			// 检查是否满足第二项条件
			if int(math.Abs(float64(arr[path[1]] - arr[x]))) <= b {
				// 做选择
				path = append(path, x)
				// 继续递归
				dfs(arr, a, b, c, path, res)
				// 撤销选择
				path = path[: len(path)-1 : cap(path)]
			}
		}
		return
	}

	if len(path) == 1 {	//准备选择第二个数
		for x:=path[0]+1; x < len(arr)-1; x++ {		// <len(arr)-1
			// 检查是否满足第一项条件
			if int(math.Abs(float64(arr[path[0]] - arr[x]))) <= a {
				// 做选择
				path = append(path, x)
				// 继续递归
				dfs(arr, a, b, c, path, res)
				// 撤销选择
				path = path[: len(path)-1 : cap(path)]
			}
		}
		return
	}

	if len(path) == 0 {	//准备选择第一个数
		for x:=0; x < len(arr)-2; x++ {		// <len(arr)-2
			// 做选择
			path = append(path, x)
			// 继续递归
			dfs(arr, a, b, c, path, res)
			// 撤销选择
			path = path[: len(path)-1 : cap(path)]
		}
		return
	}
}