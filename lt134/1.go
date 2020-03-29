package lt134

// 加油站

// 1. 暴力模拟 O(n2)
// gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]
// 线性遍历每一个站点作为起始站点，然后判断从其出发能不能够绕一圈
func canCompleteCircuit(gas []int, cost []int) int {
	var g, c []int
	for i := 0; i < len(gas); i++ {
		g = append(gas[i:], gas[:i]...)
		c = append(cost[i:], cost[:i]...)
		if can(g, c) {
			return i
		}
	}
	return -1
}

func can(g []int, c []int) bool {
	curgas := 0
	for i := 0; i < len(g); i++ { // 比如有5个站，要走一个回路，就有5段路程，都走完了就是走到了
		curgas += g[i]
		if curgas < c[i] {
			return false
		}
		curgas -= c[i] // 去往下一站
	}
	return true
}

// 2. 优化为O(n)
// 设置两个变量。 sum判断当前指针的有效性；total判断整个数组是否有解，
// 有就返回sum得到的下标，否则返回-1
func canCompleteCircuit2(gas []int, cost []int) int {
	total := 0
	j := -1
	// 这里利用的是这样一个性质： 将total为区间 [0:n] 的数字（汽油-花费）之和
	// 假如说当前用sum记录了 [0:k] （从第0个加油站出发，发现到不了第k个加油站）
	// 那么只要total >=0，则必然有 total([k+1:n]) > 0，且满足 [0:k]的花费缺口
	// 那么只需要一次遍历，找到第一个不满足的位置（j），从下一个位置开始出发，必然能完成行车循环
	for i, sum := 0, 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]   // 当前这个出发点是否有解
		total += gas[i] - cost[i] // 总体是否有解
		if sum < 0 {              // 说明[j+1:i]这个区间内从j+1开始出发是行不通的
			j = i // 记录当前停止的位置
			sum = 0
		}
	}
	if total >= 0 { // total是把所有汽油和花费相减，看总量是否满足行车循环。>=才有解
		return j + 1
	}
	return -1
}
