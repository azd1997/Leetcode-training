package unionfind

// 并查集

// UnionFindV1 第一版的并查集。 union操作时间复杂度O(n)有待改进
type UnionFindV1 struct {
	ids   []int // id数组，id相同的元素(注意这里所说的元素指的是数组下标)是“合并的”
	count int
}

// NewUnionFindV1 新建
func NewUnionFindV1(n int) *UnionFindV1 {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i // 一开始每个元素对应的id都不相同，也就是都是孤立的
	}
	return &UnionFindV1{
		ids:   ids,
		count: n,
	}
}

// Find 找出元素p（下标p）对应的id O(1)
func (uf *UnionFindV1) Find(p int) int {
	if p < 0 || p >= uf.count {
		return -1 // 报错
	}
	return uf.ids[p]
}

// IsConnected 判断两个元素是否连接在一起 O(1)
func (uf *UnionFindV1) IsConnected(p, q int) bool {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return false
	}

	return pf == qf
}

// Union 把所有与p相连的元素（包括p）都改成与q相连 O(n)
func (uf *UnionFindV1) Union(p, q int) {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return // 异常
	}

	if pf == qf {
		return //已经连接在一起
	}
	// 不等的话，则修改p或者q的区域，这里改p所在的区域
	for i := 0; i < uf.count; i++ {
		if uf.ids[i] == pf {
			uf.ids[i] = qf
		}
	}

	return
}
