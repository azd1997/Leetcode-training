package unionfind

// 并查集V3

// 增加维持parent树平衡性的步骤

// UnionFindV3 第三版的并查集。 union操作时间复杂度O(logn)，find O(logn)
// 维护parent树的平衡性
type UnionFindV3 struct {
	parent []int // parent数组，同一连通区域的结点拥有同一个根节点/公共祖宗
	size   []int // 当前子树的结点数量
	count  int
}

// NewUnionFindV3 新建
func NewUnionFindV3(n int) *UnionFindV3 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 一开始每个元素的parent就是自己，也就是都是孤立的
	}
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1 // 每个节点所在子树初始都含有自身
	}
	return &UnionFindV3{
		parent: parent,
		size:   size,
		count:  n,
	}
}

// Find 找出元素p（下标p）对应的id O(1)
func (uf *UnionFindV3) Find(p int) int {
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
func (uf *UnionFindV3) IsConnected(p, q int) bool {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return false
	}

	return pf == qf
}

// Union 把所有与p相连的元素（包括p）都改成与q相连 O(n)
func (uf *UnionFindV3) Union(p, q int) {
	pf, qf := uf.Find(p), uf.Find(q) // 各自的祖宗
	if pf == -1 || qf == -1 {
		return // 异常
	}

	if pf == qf {
		return //已经连接在一起
	}
	// 不等的话，则将p/q的祖宗合并
	// 将小的（也就矮）接在大的那一片下面
	if uf.size[pf] > uf.size[qf] {
		uf.parent[qf] = pf
		uf.size[pf] += uf.size[qf] // 增加结点数
	} else {
		uf.parent[pf] = qf
		uf.size[qf] += uf.size[pf] // 增加结点数
	}
}

// 基于size(集合大小/子树结点数)的优化还是有些问题。

// 考虑p子树结点数很多，但树趋于扁平；而q子树节点数较少，但树趋于细长
// 按照V3的Union策略，就会使得q挂在p的下面，而这就比较不平衡

// 下一个版本考虑基于层数或者说树高度的优化
