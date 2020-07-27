/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/7/27 20:53
* @Description: The file is for
***********************************************************************/

package tree_traverse

import "fmt"

// 前序遍历：中左右 或者中右左

func PreOrderRecurse(root *TreeNode) {
	if root == nil {
		return
	}

	// 处理当前
	fmt.Println(root.Val)

	PreOrderRecurse(root.Left)
	PreOrderRecurse(root.Right)
}
