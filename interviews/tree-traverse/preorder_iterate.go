/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/7/27 21:09
* @Description: The file is for
***********************************************************************/

package tree_traverse

import "fmt"

func PreOrderIterate(root *TreeNode) {
	if root == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	cur := root		// 游标节点

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			// 处理当前节点
			fmt.Println()
		}
	}


}
