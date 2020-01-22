package lt208


// 实现前缀树

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
//示例:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//说明:
//
//你可以假设所有的输入都是由小写字母 a-z 构成的。
//保证所有输入均为非空字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 官方题解 https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/shi-xian-trie-qian-zhui-shu-by-leetcode/
// 前缀树的应用
// 搜索框自动补全
// 拼写检查
// IP路由(最长前缀匹配)
// 九宫格打字预测
// 单词游戏


//还有其他的数据结构，如平衡树和哈希表，使我们能够在字符串数据集中搜索单词。为什么我们还需要 Trie 树呢？尽管哈希表可以在 O(1)O(1) 时间内寻找键值，却无法高效的完成以下操作：
//
//找到具有同一前缀的全部键值。
//按词典序枚举字符串的数据集。
//Trie 树优于哈希表的另一个理由是，随着哈希表大小增加，会出现大量的冲突，时间复杂度可能增加到 O(n)O(n)，其中 nn 是插入的键的数量。与哈希表相比，Trie 树在存储多个具有相同前缀的键时可以使用较少的空间。此时 Trie 树只需要 O(m)O(m) 的时间复杂度，其中 mm 为键长。而在平衡树中查找键值需要 O(m \log n)O(mlogn) 时间复杂度。




// 实现也见 lt14






// 字典树节点。 仅包含 'a' ~ 'z'
// 由于字母本身用作links下标，所以TrieNode只存连接关系
type TrieNode struct {
	links []*TrieNode	// 子节点的链接数组
	r int	// r=26，每个节点最多有R个子节点，用于初始化链接links数组
	isEnd bool	// 标记单词是否到达末尾
	// 这里isEnd仅用作字符串结束的标志，
	// 如果想要记录每个字符串出现的次数，比如说文件中的单次统计，
	// 可以将isEnd替换成occurance，记录字符串出现次数(每次插入删除时记得修改就行)

	// size int	// 非空子节点的个数
}

func NewTrieNode() *TrieNode {
	tnode := &TrieNode{}
	tnode.links = make([]*TrieNode, tnode.r)
	return tnode
}


func (tnode *TrieNode) ContainsKey(ch byte) bool {
	return tnode.links[ch-'a'] != nil
}

func (tnode *TrieNode) Get(ch byte) *TrieNode {
	return tnode.links[ch-'a']
}

func (tnode *TrieNode) Put(ch byte, node *TrieNode) {
	tnode.links[ch-'a'] = node
}

func (tnode *TrieNode) IsEnd() bool {
	return tnode.isEnd
}

func (tnode *TrieNode) SetEnd() {
	tnode.isEnd = true
}


// 前缀树

type Trie struct {
	root *TrieNode
	size int	// 不同单词的个数
	// totalSize int	// 所有单词所有出现，总次数
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := NewTrieNode()
	return Trie{root:root}
}

func (this *Trie) Root() *TrieNode {
	return this.root
}

func (this *Trie) Size() int {
	return this.size
}

// 向Trie中插入键
// 插入键时从根开始搜索它对应于第一个键字符的链接，链接存在则沿着链接移动到下一个键字符
// 链接不存在，则创建一个新节点，将其与父节点的链接相连，该链接与当前的键字符相匹配
// 重复前两行操作，直到键的最后一个字符，将该字符标记为结束节点。
/** Inserts a word into the trie. */
// O(m)/O(m)， m为键长，也就是单词长度，最坏情况下，这个单词一个前缀都没有，就需要O(m)的空间了
//
// 1、从头到尾遍历字符串的每一个字符
//2、从根节点开始插入，若该字符存在，那就不用插入新节点，要是不存在，则插入新节点
//3、然后顺着插入的节点一直按照上述方法插入剩余的节点
//(4)、为了统计每一个字符串出现的次数，应该在最后一个节点插入后occurances++，表示这个字符串出现的次数增加一次
//
func (this *Trie) Insert(word string)  {
	node := this.root
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		if !node.ContainsKey(cur) {
			node.Put(cur, NewTrieNode())
		}
		node = node.Get(cur)
	}
	node.SetEnd()
	this.size++
}


/** Returns if the word is in the trie. */
// O(m)/O(1)
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.IsEnd()
	// 节点必须非空且是word的结尾
}

// 搜索一个前缀或者整个键，返回搜索结束时的节点
func (this *Trie) SearchPrefix(word string) *TrieNode {
	node := this.root
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		if node.ContainsKey(cur) {
			node = node.Get(cur)
		} else {return nil}
	}
	return node
}


/** Returns if there is any word in the trie that starts with the given prefix. */
// 和Search的区别是这里只需要遍历到prefix末尾然后返回true，不关心isEnd
// O(m)/O(1)
func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

//参考：https://blog.csdn.net/johnny901114/article/details/80711441
//https://www.jianshu.com/p/6f81da81bd02

// Trie树的删除操作：3种情况
// 1，如果单词是另一个单词的前缀，只需要把该word的最后一个节点的isWord的改成false
// 2，如果单词的所有字母的都没有多个分支，删除整个单词
// 3，如果单词的除了最后一个字母，其他的字母有多个分支，把最后一个有分支的字符之后的字符节点全删掉
//
//
// 思路：
// 1.从根节点的子节点开始，判断该当前节点是否为空，若为空且没有到达要删除字符串的
// 最后一个字符，则不存在该字符串。若已经到达叶子节点但是没有遍历完整个字符串，
// 说明字符串也不存在。例如要删除的是"aazz1111"，但只匹配到"aazz"
// 2. 只有当要删除的字符串找到时并且最后一个字符是叶子节点才需要删除，
// 而且任何删除动作都只发生在叶子节点。例如要删除"byebye"，但trie里还有"byebyeha"
// 说明byebye不需要删除，只需将isEnd置否(occurance清零)
// 3. 遍历到最后一个字符时，也就是说字典树中有这个单词。必须
// 将最末尾那个字符节点的isEnd置否(occurance清零)，这样就表示这个单词不存在了，
// 至于是否删除这个单词的后缀部分节点，取决于情况2

// TODO：先搁置，写不来了。

func (this *Trie) Delete(word string) bool {
	if this.root.links == nil || word=="" {
		return false
	}

}

// node表示当前节点，d表示当前是word第几个字符
func (this *Trie) delete(node *TrieNode, word string, d int) {
	if d==len(word) {
		node.isEnd = false	// 先将isEnd置否
	} else {
		this.delete(node.links[word[d] - 'a'], word, d+1)
	}

	if node.IsEnd() {
		return
	}

	item := node
	for item != nil {}
}