package lt297

import (
	"strconv"
	"strings"
)

// 二叉树的序列化与反序列化
// 将二叉树序列化成[1,2,3,null,..]形式的字符串
// 以及反序列化

// 序列化思路：BFS层次优先遍历
// 反序列化思路：同样按层构建树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{} // 相当于工具类

// 序列化： BFS层序遍历
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := []*TreeNode{root}
	encoded := ""
	for len(queue) != 0 { // 队列非空
		root = queue[0]
		queue = queue[1:] // 队头出队
		if root != nil {
			queue = append(queue, root.Left, root.Right) // 将左右子节点从队尾入队
			encoded += strconv.Itoa(root.Val) + ","
		} else {
			encoded += "null,"
		}
	}

	// 去除右边的"null,"，没必要存储
	encoded = trimRightNulls(encoded)
	// 去除末尾","，加上左右括号
	encoded = "[" + encoded[:len(encoded)-1] + "]"

	return encoded
}

func (c *Codec) deserialize(encoded string) *TreeNode {
	// 边界判断
	if len(encoded) == 0 {
		return nil
	}
	// 去除左右括号
	encoded = encoded[1 : len(encoded)-1]
	// 边界判断
	if len(encoded) == 0 {
		return nil
	}
	// 按","分割字符串
	items := strings.Split(encoded, ",")
	// 构造根节点
	rootVal, _ := strconv.Atoi(items[0]) // 第一个值不可能是null
	root := &TreeNode{Val: rootVal}

	// 构建辅助队列
	queue := []*TreeNode{root}
	i := 0
	for {
		if i++; i >= len(items) { // i先+1，如果超界，则退出
			break
		}

		node := queue[0]
		queue = queue[1:] // 弹出队头

		if items[i] != "null" {
			leftVal, _ := strconv.Atoi(items[i])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}

		if i++; i >= len(items) { // i先+1，如果超界，则退出
			break
		}
		if items[i] != "null" {
			rightVal, _ := strconv.Atoi(items[i])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
	}

	return root
}

func trimRightNulls(data string) string {
	target := "null,"
	size := len(target) //5
	for len(data) > size && data[len(data)-size:] == target {
		data = data[:len(data)-size]
	}
	return data
}
