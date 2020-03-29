package lt820

// 单词的压缩编码

// # 是完全隔断，且不计入下标
// 这道题主要的知识点是后缀树。 将words构建成一颗后缀树
// 后缀相同的字符串可以进行压缩

// 也就是说
// 第一种是纯暴力思路，枚举每个单词的所有后缀，看当前是否是其他单词的后缀，是则可以压缩
// 第二种思路就是trie树，当然这里用来存后缀而非前缀
// 第三种思路，可以利用字符串哈希算法，暴力匹配后缀
// 第四种思路是按后缀进行排序相邻

// 利用哈希表 O(sum(wordL ^ 2)) / O(sum(wordL))
func minimumLengthEncoding(words []string) int {
	// 存储所有单词，便于查找
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}
	// 遍历单词列表，尝试删去每个单词的后缀
	// 也就是压缩掉那些是别人后缀的单词
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(set, word[i:])
		}
	}
	// 再一遍查看set，看还剩多少个单词，每个单词后面再加一个'#'
	ans := 0
	for word := range set {
		ans += len(word) + 1
	}
	return ans
}

// 2. Trie树   O(sum(wordL)) / O(S * sum(wordL))	// S为每个TrieNode存储的额外信息
func minimumLengthEncoding2(words []string) int {
	// 初始化一颗Trie树（建一个根节点）
	trie := &TrieNode{}
	// 存储每个单词的首字母的Trie树结点，以及该单词在单词列表中的下标
	nodes := make(map[*TrieNode]int)
	for i, word := range words {
		cur := trie
		for j := len(word) - 1; j >= 0; j-- {
			cur = cur.get(word[j])
		}
		nodes[cur] = i
	}
	// 查看nodes，看还剩多少个单词(在Trie树中为叶子节点)，每个单词后面再加一个'#'
	ans := 0
	for node := range nodes {
		if node.count == 0 {
			ans += len(words[nodes[node]]) + 1
		}
	}
	return ans
}

// Trie树结点基本定义
type TrieNode struct {
	children [26]*TrieNode
	count    int
}

func (n *TrieNode) get(char byte) *TrieNode {
	if n.children[char-'a'] == nil {
		n.children[char-'a'] = &TrieNode{}
		n.count++
	}
	return n.children[char-'a']
}
