package lt218

import "sort"

// 天际线问题

// 只能求助题解区了，不会就要挨打...

// 1. 扫描线解法 参考题解区 ivan_allen
// ivan_allen的解法是C++版本，使用multiset
// multiset基于平衡二叉树，对添加其中的元素自动排序。multiset相比于set允许重复元素
// 知道了这些后，使用z
func getSkyline1(buildings [][]int) [][]int {

}


// 2. 扫描线解法 参考题解区 codes
func getSkyline2(buildings [][]int) [][]int {

	// 1. 定义拐点。
	// 每个建筑的左上角和右上角称拐点， 每个建筑物两个拐点
	type boxing struct {
		// node : 建筑物序号
		// pos : 拐点x坐标
		// height: 拐点y坐标
		// side: 标记拐点为当前建筑物的左侧拐点还是右侧拐点
		node, pos, height, side int
	}

	// 2. 收集所有建筑物拐点
	var box []*boxing
	for i, v := range buildings {
		box = append(box,
			&boxing{i, v[0], v[2], 0},
			&boxing{i, v[1], v[2], 1})
	}

	// 3. 对所有拐点根据横坐标pos排序
	sort.Slice(box, func(i, j int) bool {
		return box[i].pos < box[j].pos
	})

	// 4.
	var hill, res [][]int
	maxhill := -1
	for i:=0; i<len(box); i++ {
		v := box[i]
		if v.side==0 {
			hill = append(hill, []int{v.height, v.node})
		} else {
			for idx, vv := range hill {
				if vv[1] == v.node {
					hill = append(hill[:idx], hill[idx+1:]...)
					break
				}
			}
		}

		// 当前位置(x)，最高位置(y)查找
		maxh := 0
		for _, vv := range hill {
			if vv[0] > maxh {maxh = vv[0]}
		}

		// 同一个位置x出现多个y的处理
		if !(i<len(box)-1 && box[i].pos==box[i+1].pos) {
			if maxhill < maxh || maxhill > maxh {
				res = append(res, []int{v.pos, maxh})
			}
			maxhill = maxh
		}
	}

	return res
}

