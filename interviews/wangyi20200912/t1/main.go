/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/12/20 3:00 PM
* @Description: The file is for
***********************************************************************/

package main

import "fmt"

type Node struct {
	left, right *Node
	id int
}

func main() {
	m, n := 0, 0
	fmt.Scan(&m, &n)

	nodes := make([]*Node, m+1)	// 下标对应id，0不使用

	for i:=0; i<n; i++ {
		id := 0
		rela := ""
		next := 0
		fmt.Scan(&id, &rela, &next)

		if nodes[id] == nil {
			nodes[id] = &Node{id:id}
		}
		if nodes[next] == nil {
			nodes[next] = &Node{id:next}
		}
		if rela == "left" {
			nodes[id].left = nodes[next]
		} else if rela == "right" {
			nodes[id].right = nodes[next]
		}
	}

	ans := sol(nodes[1])

	fmt.Println(ans)
}

// 求树中两个孩子都是叶子节点的结点数量
func sol(root *Node) int {
	if root == nil {return 0}
	cnt := 0
	help(root, &cnt)
	return cnt
}


func help(root *Node, cnt *int) int {
	if root == nil {
		return 1
	}

	//fmt.Println(root, *cnt)

	l := help(root.left, cnt)
	r := help(root.right, cnt)
	if l == 1 && r == 1 {	// 说明root是叶子节点
		return -1
	}
	if l == -1 && r == -1 {
		*cnt++
		return 0
	}
	return 0
}