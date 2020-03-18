package unionfind

// 并查集 使用QuickUnion

// V1版本的并查集采用的是相互连接的元素拥有相同的id，但是其union操作比较慢
// V2版本的并查集不再记录id，而是记录当前节点的父亲结点，union操作时直接将
// 两个节点的祖宗节点进行挂载（连接）即可（树形结构）
// 即便采用的树形结构，这里仍可以使用数组作为基本存储结构，结点指针即为数组下标

// 既然采用了树形结构，那么就需要考虑树高度的平衡性，这样才能降低时间复杂度。那么请看V3版本

// UnionFindV2 第二版的并查集。 union操作时间复杂度O(logn)，find O(logn)
type UnionFindV2 struct {
	parent []int // parent数组，同一连通区域的结点拥有同一个根节点/公共祖宗
	count  int
}

// NewUnionFindV2 新建
func NewUnionFindV2(n int) *UnionFindV2 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 一开始每个元素的parent就是自己，也就是都是孤立的
	}
	return &UnionFindV2{
		parent: parent,
		count:  n,
	}
}

// Find 找出元素p（下标p）对应的id O(1)
func (uf *UnionFindV2) Find(p int) int {
	if p < 0 || p >= uf.count {
		return -1 // 报错
	}
	// 不断向上寻找parent，返回当前节点p的祖宗
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

// IsConnected 判断两个元素是否连接在一起 O(1)
func (uf *UnionFindV2) IsConnected(p, q int) bool {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return false
	}

	return pf == qf
}

// Union 把所有与p相连的元素（包括p）都改成与q相连 O(n)
func (uf *UnionFindV2) Union(p, q int) {
	pf, qf := uf.Find(p), uf.Find(q) // 各自的祖宗
	if pf == -1 || qf == -1 {
		return // 异常
	}

	if pf == qf {
		return //已经连接在一起
	}
	// 不等的话，则将p/q的祖宗合并
	uf.parent[pf] = qf // pf退让一步，让qf做了自己爹...
}
