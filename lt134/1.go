package lt134

// 加油站

// 1. 暴力模拟 O(n2)
// gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]
// 线性遍历每一个站点作为起始站点，然后判断从其出发能不能够绕一圈
func canCompleteCircuit(gas []int, cost []int) int {
	var g, c []int
	for i:=0; i<len(gas); i++ {
		g = append(gas[i:], gas[:i]...)
		c = append(cost[i:], cost[:i]...)
		if can(g, c) {return i}
	}
	return -1
}

func can(g []int, c []int) bool {
	curgas := 0
	for i:=0; i< len(g); i++ {    // 比如有5个站，要走一个回路，就有5段路程，都走完了就是走到了
		curgas += g[i]
		if curgas < c[i] {return false}
		curgas -= c[i]  // 去往下一站
	}
	return true
}


// 2. 优化为O(n)
// TODO