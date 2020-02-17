package lt207

// 课程表

// 1. 入度表(广度优先遍历)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 入度表
	indegrees := make([]int, numCourses)
	// 根据前置课程计算各个课程节点的入度
	for _, v := range prerequisites {
		indegrees[v[0]]++
	}
	// 辅助队列
	queue := make([]int, 0)
	// 将入度为0的课程率先压入队尾
	for i:=0; i<numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS
	for len(queue) != 0 {
		// 取出队头元素，为当前入度为0的课程，pre记录的是它的编号(0~n-1)
		pre := queue[0]; queue = queue[1:]
		// 课程数-1，相当于原先入度为0的课程都处理过了，不会成环
		numCourses--
		// 遍历先决条件数组，找到先决修课程为pre的，将其入度-1
		// 如果-1后为0，那么入队
		for _, v := range prerequisites {
			if v[1] != pre {continue}
			indegrees[v[0]]--
			if indegrees[v[0]] == 0 {
				queue = append(queue, v[0])
			}
		}
	}

	// 如果存在环，最后不可能为0
	return numCourses == 0
}

// 2. 深度优先遍历
// 通过DFS判断图中是否有环
func canFinish2(numCourses int, prerequisites [][]int) bool {
	// 构建邻接矩阵
	adjacency := make([][]int, numCourses)
	for i:=0; i<numCourses; i++ {adjacency[i] = make([]int, numCourses)}
	for _, v := range prerequisites {
		adjacency[v[1]][v[0]] = 1	// 表示课程v[1]到v[0]有一条边到达(v[1]是v[0]的先修课程)
	}
	flags := make([]int, numCourses)
	for i:=0; i<numCourses; i++ {	// i为课程编号
		if !dfs(adjacency, flags, i) {return false}
	}

	return true
}

// 从课程i这个节点出发深度优先搜索，看是否存在环
func dfs(adjacency [][]int, flags []int, i int) bool {
	if flags[i] == 1 {return false}		// 本轮访问过
	if flags[i] == -1 {return true}		// 其他轮访问过
	flags[i] = 1	// 标志本轮已经访问过这门课了
	for j:=0; j<len(adjacency); j++ {
		// 如果有课程j的先修课程为当前课程i并且
		if adjacency[i][j] == 1 && !dfs(adjacency, flags, j) {
			return false
		}
	}
	flags[i] = -1	// 本轮结束，将flag标记为-1
	return true
}