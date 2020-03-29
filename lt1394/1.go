package lt1394

// 幸运数字

// 哈希表暴力干
func findLucky(arr []int) int {
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v]++
	}

	luckys := make(map[int]int)
	for k, v := range mp {
		if k == v {
			luckys[k] = v
		}
	}

	if len(luckys) == 0 {
		return -1
	}

	maxK, maxV := 0, 0
	for k, v := range luckys {
		if v > maxV {
			maxK, maxV = k, v
		}
	}
	return maxK
}

// 一种更简洁一些的做法：利用数组作哈希表，同时还保证了有序性
func findLucky2(arr []int) int {
	limit := 500
	cnts := make([]int, limit+1)
	for _, v := range arr {
		cnts[v]++
	}

	// 倒序遍历cnts
	index := -1
	for i := limit; i >= 1; i-- {
		if cnts[i] == i {
			index = i
			break // 第一个找到的就是最大的幸运数
		}
	}

	return index
}
