package search

// 广度优先搜索 与队列密切相关

// 模板1.
//func BFS(root, target *TreeNode) (step int) {
//	init a queue	// 一个队列用于存储所有等待处理的节点
//	var step int	// 记录从root到当前处理的节点的距离（步数）（用树的BFS来讲就是从上往下的层数）
//
//	add root to queue	// 将根节点加入到队列
//
//	// 开始BFS
//	for queue is not empty {
//		step = step + 1 	// 步数加一（到了下一轮）
//		// 遍历还在队列中的节点
//		size := queue.size
// 		for i:=0; i<size; i++ {
// 			current := the first node in queue
//			return step if current is target
//			for next in neighbors of current {
//				add next to queue
//			}
//			remove the first node in queue
//		}
//	}
//	return -1
//}

// 模板2.
// 对于存在环形的图结构，必须确保永远不会访问节点第二遍，不然可能无限循环。
// 针对这个需求，可以加一个哈希集来记录已经访问过的节点，避免重复
//
//func BFS(root, target *TreeNode) (step int) {
//	init a queue	// 一个队列用于存储所有等待处理的节点
//  init a set "used"	// 存储所有使用过的节点
//	var step int	// 记录从root到当前处理的节点的距离（步数）（用树的BFS来讲就是从上往下的层数）
//
//	add root to queue	// 将根节点加入到队列
//	add root to used
//
//	// 开始BFS
//	for queue is not empty {
//		step = step + 1 	// 步数加一（到了下一轮）
//		// 遍历还在队列中的节点
//		size := queue.size
// 		for i:=0; i<size; i++ {
// 			current := the first node in queue
//			return step if current is target
//			for next in neighbors of current {
//				if next is not in used {
//					add next to queue
//					add next to used
//				}
//			}
//			remove the first node in queue
//		}
//	}
//	return -1
//}

// 两种情况下不需要使用哈希集合：
// 1. 确定没有循环，例如，遍历树
// 2. 确定希望多次将结点添加到队列