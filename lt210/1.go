package lt210

// 课程表II

// 1. 广度优先搜索
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 边界
	if numCourses <= 0 {
		return nil
	}

	// 1. 邻接表，记录当前节点(课程)的后继节点
	// 这里由于课程数量固定，用二维矩阵也是一样的
	graph := make([]map[int]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make(map[int]bool)
	}
	// 2. 入度表. 记录每门课程的先修数量
	indegrees := make([]int, numCourses)
	// 3. 填充邻接表和入度表
	for _, v := range prerequisites {
		graph[v[1]][v[0]] = true
		indegrees[v[0]]++
	}
	// 4. 辅助队列，将入度为0的课程先入队
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 3. BFS
	res := make([]int, 0)
	var nextCourses map[int]bool
	var cur int
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]      // 出队课程
		res = append(res, cur) // 加入返回数组

		// 得到所有的后继课程
		nextCourses = graph[cur]

		for next := range nextCourses {
			indegrees[next]--
			if indegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if numCourses != len(res) {
		return nil
	}
	return res
}

// 参照207的BFS，这里的graph显然可以省去，节省空间，但是很有可能增加了时间开销

// DFS解法
func findOrder2(numCourses int, prerequisites [][]int) []int {
	// 1.检查numCourses
	if numCourses <= 0 {
		return nil
	}

	// 2. 检查prerequisites
	if len(prerequisites) == 0 {
		// 课程没有依赖关系（图中没有有向边），说明随便怎样的顺序都可以
		order := make([]int, numCourses)
		for i := 0; i < numCourses; i++ {
			order[i] = i
		}
		return order
	}

	// 3. 标记数组 visited/marked/flags之类的名字
	// 0表示没访问，1表示当前轮已访问，2表示前面的其他轮已经访问
	visited := make([]int, numCourses)

	// 4. 初始化有向图 graph 的 begin，数组的每一项代表起点，是前驱结点
	graph := make([]map[int]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make(map[int]bool)
	}

	// 5. 初始化有向图 end， graph[i].map中存的是后继节点们
	for _, p := range prerequisites {
		graph[p[1]][p[0]] = true // 将后继节点加入
	}

	// 6. 使用栈stack记录递归的顺序
	// 这个Stackz在DFS过程中会将没有后继节点的课程放到栈底位置，
	// 没有前驱节点的课程压到栈顶
	stack := make([]int, 0) // 栈其实就是DFS路径
	for i := 0; i < numCourses; i++ {
		if dfs(i, graph, visited, &stack) { // 如果发现环，那么就返回空数组
			return nil
		}
	}

	// 7. DFS过程中一直没有出现成环情况
	// 那么说明当前的栈stack存的就是课程表路径,将顺序颠倒即可
	if len(stack) != numCourses {
		return nil
	} // 其实肯定是==的，只不过，假如stack没这么长，那样的话下面交换就会panic
	for i := 0; i < numCourses/2; i++ {
		stack[i], stack[numCourses-i-1] = stack[numCourses-i-1], stack[i]
	}

	return stack
}

// 从课程x开始DFS，判断是否成环。false不成环，true成环，
// 此外还返回path []int。仅当path是完整课表顺序才返回
func dfs(i int, graph []map[int]bool, visited []int, stack *[]int) bool {
	// 1. 检查visited数组
	if visited[i] == 1 { // 当前轮访问过i，说明成环
		return true
	}
	if visited[i] == 2 { // 其他轮访问过了，从i开始不会成环
		return false
	}

	// 2. 标记当前轮正在访问i
	visited[i] = 1

	// 3. 对后继节点进行DFS
	for j := range graph[i] {
		// 后继节点存在成环情况，直接返回true
		if dfs(j, graph, visited, stack) {
			return true
		}
	}

	// 4. i的所有后继节点都不成环，所以从i开始也不成环
	// 将visited[i]置2
	visited[i] = 2

	// 5. 将i入栈
	*stack = append(*stack, i)

	return false // 不成环
}
