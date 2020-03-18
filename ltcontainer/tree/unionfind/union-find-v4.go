package unionfind

// 并查集V4

// UnionFindV4 第四版的并查集。 union操作时间复杂度O(logn)，find O(logn)
// 维护parent树的平衡性。 基于rank(秩)的优化。 rank[p]表示p子树的高度
type UnionFindV4 struct {
	parent []int // parent数组，同一连通区域的结点拥有同一个根节点/公共祖宗
	rank   []int // 当前子树的高度/当前节点的秩
	count  int
}

// NewUnionFindV4 新建
func NewUnionFindV4(n int) *UnionFindV4 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 一开始每个元素的parent就是自己，也就是都是孤立的
	}
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[i] = 1 // 每个节点所在子树初始都含有自身
	}
	return &UnionFindV4{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

// Find 找出元素p（下标p）对应的id O(1)
func (uf *UnionFindV4) Find(p int) int {
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
func (uf *UnionFindV4) IsConnected(p, q int) bool {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return false
	}

	return pf == qf
}

// Union 把所有与p相连的元素（包括p）都改成与q相连 O(n)
func (uf *UnionFindV4) Union(p, q int) {
	pf, qf := uf.Find(p), uf.Find(q) // 各自的祖宗
	if pf == -1 || qf == -1 {
		return // 异常
	}

	if pf == qf {
		return //已经连接在一起
	}
	// 不等的话，则将p/q的祖宗合并
	// 将小的（也就矮）接在大的那一片下面
	if uf.rank[pf] > uf.rank[qf] {
		uf.parent[qf] = pf
	} else if uf.rank[pf] < uf.rank[qf] {
		uf.parent[pf] = qf
	} else { // ==
		uf.parent[pf] = qf // pf让一步，挂在qf下面
		uf.rank[qf]++
	}
}
