package lt210

// 课程表II


// 1. 广度优先搜索
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 边界
	if numCourses <= 0 {return nil}

	// 1. 邻接表，记录当前节点(课程)的后继节点
	// 这里由于课程数量固定，用二维矩阵也是一样的
	graph := make([]map[int]bool, numCourses)
	for i:=0; i<numCourses; i++ {
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
	for i:=0; i<numCourses; i++ {
		if indegrees[i]==0 {
			queue = append(queue, i)
		}
	}
	// 3. BFS
	res := make([]int, 0)
	var nextCourses map[int]bool
	var cur int
	for len(queue) != 0 {
		cur = queue[0]; queue = queue[1:]	// 出队课程
		res = append(res, cur)		// 加入返回数组

		// 得到所有的后继课程
		nextCourses = graph[cur]

		for next := range nextCourses {
			indegrees[next]--
			if indegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if numCourses != len(res) {return nil}
	return res
}

// 参照207的BFS，这里的graph显然可以省去，节省空间，但是很有可能增加了时间开销