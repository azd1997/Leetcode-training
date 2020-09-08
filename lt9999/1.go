package lt9999

import "fmt"

//[1,2,3,4,5,6,7]
//3
//[7,1,4,6,null,5,3,null,null,null,null,null,2]
//3

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	if distance < 2 {
		return 0
	}

	total := 0
	help(root, &total, distance)
	return total
}

// 计算node子树中的好叶节点对的数量，并且返回此时路径最大值
// 对于每科子树而言，其根节点需要记录：
// 1. 其下所有叶子节点到其的路径长度（当然，已经>=distance-1的可以剪枝）的列表list
// 2. 到递归到该根节点时，组合其左右子树的listA,listB，再加上2，看有没有<=distance的，有则给总体返回值加1，
// 对于递归函数而言，应该返回叶节点到根节点路径长度的list

// 而且这个list使用map替代slice可能更好

// 返回的list是升序排列的
func help(node *TreeNode, total *int, distance int) map[int]int {     // <path, number>
	if node == nil {
		return map[int]int{}    // 注意是空，长度为0
	}

	listA := help(node.Left, total,distance)
	listB := help(node.Right, total,distance)

	// 叶节点到根节点路径已经超了distance
	if listA[-1] > 0 || listB[-1] > 0 {
		return map[int]int{-1:1}    //数量不重要了，随便填一个
	}

	if len(listA) == 0 && len(listB) == 0 {     // 这里要注意区别叶子节点和不满足条件的子树根节点，那么设置一个-1的键用来标识
		return map[int]int{0:1}     // 是说明是叶子节点
	}

	if len(listA) == 0 {
		checkListAndAdd(listB, distance)
		return listB
	}

	if len(listB) == 0 {
		checkListAndAdd(listA, distance)
		return listA
	}

	// 然后就是都有可能满足条件的叶节点

	// 穷举组合
	fmt.Println(*total, node)
	for k, v := range listA {
		for k1, v1 := range listB {
			if k + k1 + 2 <= distance {
				fmt.Println(k,v,k1,v1,*total)
				*total = *total + v * v1
			}
		}
	}
	// 将list合并，并再次检查
	for k, v := range listA {
		listB[k] = listB[k] + v
	}
	checkListAndAdd(listB, distance)

	return listB
}

func checkListAndAdd(list map[int]int, distance int) {
	// 输入时如果是空，那么直接返回空
	if len(list) == 0 {return}

	// 输入时不为空，那么对所有键的值进行加1操作，如果加之后值过大，则删除键
	for k := range list {
		if k + 1 > distance - 2 {
			delete(list, k)
			continue
		}
		// 否则
		list[k+1] = list[k]     // 不用担心数据会被覆盖掉，因为for range是取了份拷贝的
	}
	// 如果准备输出时发现list空，则返回-1
	if len(list) == 0 {
		list[-1] = 1
	}
	return
}
