package unionfind

// 并查集V5

// 前面V2/V3/V4都是在不断优化union操作，但其实find操作也是可以优化的
// 在V4版本的树形结构下，find操作需要从当前结点不断向上寻找到祖宗节点
// 如果途中没有其他分支或分支较少，也就是说树趋于细长，
// 那么相当于白向上跑这么久，没有分支，就意味着不需要检查这么多嘛

// 例如
// 4 -> 3 -> 2 -> 1 -> 0 |
// find(4)需要向上找四步
// 但其实我们可以想像要是每次能多走几步并且将细长的树变得矮胖就好了

// 对于4，先向上“跳两步”接在2下面，与3成为兄弟；
// 接着再将4的这颗子树向上跳2步，接在了0下面。
// 这样树变成了：

// 			0
// 		   / \
//        1   4
//           / \
//          3   2

// 树的高度就变矮了

// 这个策略叫做路径压缩
// 核心在于每次向上跳两步，跳三/四/更多步当然也是可以的。
// 因为最后祖宗节点一直都指向自己，所以向上跳不会跳出子树范围。

// UnionFindV5 第五版的并查集。增加路径压缩的优化，只需要在find时增加一句代码
type UnionFindV5 struct {
	parent []int // parent数组，同一连通区域的结点拥有同一个根节点/公共祖宗
	rank   []int // 当前子树的高度/当前节点的秩
	count  int
}

// NewUnionFindV5 新建
func NewUnionFindV5(n int) *UnionFindV5 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 一开始每个元素的parent就是自己，也就是都是孤立的
	}
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[i] = 1 // 每个节点所在子树初始都含有自身
	}
	return &UnionFindV5{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

// Find 找出元素p（下标p）对应的id O(1)
func (uf *UnionFindV5) Find(p int) int {
	if p < 0 || p >= uf.count {
		return -1 // 报错
	}
	// 不断向上寻找parent，返回当前节点p的祖宗
	for p != uf.parent[p] {
		// find的同时将细长的树拉得矮胖
		uf.parent[p] = uf.parent[uf.parent[p]] // 向上跳两步，并将树变矮
		p = uf.parent[p]
	}
	return p
}

// IsConnected 判断两个元素是否连接在一起 O(1)
func (uf *UnionFindV5) IsConnected(p, q int) bool {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return false
	}

	return pf == qf
}

// Union 把所有与p相连的元素（包括p）都改成与q相连 O(n)
func (uf *UnionFindV5) Union(p, q int) {
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

// 既然能将树的高度变矮，为什么不直接在find的过程中将所有结点都接在子树的根节点下面呢？
// 例如

// 			0
// 		  /| |\
//       1 2 3 4

// 这样就需要递归操作了，见V6版本
