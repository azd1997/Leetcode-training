package lt365

// 水壶问题

// 自己没想出来咋做
// 题解区主要有三种做法：
// 1. DFS穷举所有状态
// 2. BFS穷举所有状态
// 3. 数学解法。当且仅当 z = k * gcd(x,y) 时才为true

// 重点关注如何穷举：其实就是模拟
// 每一次的时候都有下面几种选择可以做：
// 把x倒空；把x装满
// 把y倒空；把y装满
// 把x倒到y直至倒空或装满； 把y倒到x中直至倒空或装满
//
// 每次选择结束之后检查是否符合条件 true
//
// （由于x*y比较大，使用递归可能导致函数调用栈溢出，建议使用栈迭代）
// 其实这里DFS就是一种搜索、回溯

// 1. DFS O(xy)/O(xy)
func canMeasureWater(x int, y int, z int) bool {
	if x == 0 {
		return y == z
	}
	if y == 0 {
		return x == z
	}

	// DFS
	stack := make([][2]int, 1) // 存[x,y]当前的装水量，初始值为0
	stack[0] = [2]int{0, 0}
	visited := make(map[[2]int]bool) // 记录当前 [x,y]是否访问过
	var cur [2]int
	for len(stack) != 0 {
		// 取出当前栈顶的装水情况
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		// 检查是否true
		if cur[0] == z || cur[1] == z || cur[0]+cur[1] == z {
			return true
		}
		// 是否已访问过，已访问过则跳过
		if visited[cur] {
			continue
		}
		// 添加到已访问
		visited[cur] = true
		// 做选择，将作出的六种选择后的新状态都入栈
		stack = append(stack,
			[2]int{x, cur[1]}, // 将x装满
			[2]int{0, cur[1]}, // 将x倒空
			[2]int{cur[0], y}, // 将y装满
			[2]int{cur[0], 0}, // 将y倒空
			[2]int{cur[0] - min(cur[0], y-cur[1]), cur[1] + min(cur[0], y-cur[1])}, // 将x倒入y，直至y满或x空
			[2]int{cur[1] + min(cur[1], x-cur[0]), cur[1] - min(cur[1], x-cur[0])}, // 将y倒入x，直至x满或y空
		)
	}
	return false // 所有状态都试过了，没有，只能返回false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 提交后栈内存分配不足，崩掉了

// 2. BFS
// 在本题和dfs其实一样，就是把栈换成队列
// 该崩还是崩

// 3. 数学
// https://leetcode-cn.com/problems/water-and-jug-problem/solution/shui-hu-wen-ti-by-leetcode-solution/
func canMeasureWater2(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return x+y == z || z == 0 // 注意z==0的情况
	}
	return z%gcd(x, y) == 0
}

// 辗转相除法 求最大公约数
func gcd(a, b int) int {
	tmp := a
	for tmp != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}
