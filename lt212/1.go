package lt212

// 单词搜索II

// 和 单词搜索lt79 的区别在于：原来的目标单词变成单词列表；而返回是否存在也变成了返回所有匹配的单词
//
// 由此带来的最大区别是
// 单词搜索中只需要通过比较board[i][j]==word[idx]来决定是否继续递归搜索
// 而本题则是需要在一个单词列表匹配前缀
// 高效的匹配前缀的方法：前缀树


//====================前缀树实现========================

// 26叉树
type TrieNode struct {
	next []*TrieNode    // 子节点
	isEnd bool          // 标记单词结尾
}

func NewTrieNode() *TrieNode {
	next := make([]*TrieNode, 26)
	return &TrieNode{next:next}
}

func (n *TrieNode) Get(ch byte) *TrieNode {
	return n.next[ch-'a']
}

func (n *TrieNode) Put(ch byte) {
	n.next[ch-'a'] = NewTrieNode()
}

func (n *TrieNode) Has(ch byte) bool {
	return n.next[ch-'a']!=nil
}

func (n *TrieNode) IsEnd() bool {return n.isEnd}

func (n *TrieNode) SetEnd() {n.isEnd = true}

func (n *TrieNode) ResetEnd() {n.isEnd = false}

type Trie struct {
	root *TrieNode
}


/** Inserts a word into the trie. */
// 遍历单词每个字符，逐渐插入，并对最后一个字符SetEnd
func (this *Trie) Insert(word string)  {
	node := this.root
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		if !node.Has(cur) {
			node.Put(cur)
		}
		node = node.Get(cur)
	}
	node.SetEnd()
}


//=====================前缀树实现结束===============


func findWords(board [][]byte, words []string) []string {
	var res []string
	m := len(board)
	if m==0 {return res}
	n := len(board[0])
	if n==0 {return res}

	// 构建字典树 O(N2)
	trie := &Trie{root:NewTrieNode()}
	for _, word := range words {
		trie.Insert(word)
	}

	// 遍历所有点
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			DFS(board, i, j, trie.root, &res, []byte{})
		}
	}

	return res
}


// 方向：上右下左
var di = [4]int{-1,0,1,0}	// 纵坐标
var dj = [4]int{0,1,0,-1}	// 横坐标

func DFS(board [][]byte, i,j int, node *TrieNode, res *[]string, cur []byte) {

	//fmt.Println(i,j,cur)

	// 检查首字母是否存在于trie节点中
	ch := board[i][j]	// 把ch作为首字母在trie中检查
	if !node.Has(ch) {return}

	// 更新节点到ch所在节点
	node = node.Get(ch)

	// 检查ch节点是否是end(匹配到完整单词，加入到结果数组)
	cur = append(cur, ch)
	if node.IsEnd() {
		*res = append(*res, string(cur))
		node.ResetEnd()		// 防止重复的单词加入
	}

	// 继续搜索，搜寻可能有的单词（例如，已找到word，继续寻找wordpress）

	// 标记当前字母为'#'
	board[i][j] = '#'

	// 上右下左 搜索
	var newI, newJ int
	for k:=0; k<4; k++ {
		newI, newJ = i+di[k], j+dj[k]
		if (newI>=0 && newI<len(board)) &&
			(newJ>=0 && newJ<len(board[0])) &&
			board[newI][newJ]!='#' {
			DFS(board, newI, newJ, node, res, cur)
		}
	}

	// 恢复board[i][j]
	board[i][j] = ch
}


// 其实本题是找出所有单词，使用BFS也是完全OK的