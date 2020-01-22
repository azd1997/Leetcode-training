package lt14

import (
	"errors"
	"github.com/azd1997/Leetcode-training/lt208"
)

// 最长公共前缀

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
// 1. 直接遍历数组，将第一字符串s作为最长公共前缀，与第二个进行匹配，假设匹配到第i位，
// 之后就对不上，那么接着拿s[:i+1]去和第三个字符串匹配，最后剩下来的就是最长公共前缀。
// 这样的做法 O(n*k)/O(1) n是数组长度，k是第一个字符串长度
// 2. 第二种思路是哈希表(map或者array)，将第一个字符串的所有前缀载入表中，然后让第二个字符串从长到短去匹配，并删除不匹配的前缀
// 	最后剩下的哈希表中最长的键就是最长公共前缀。 O(n*k)/O(k) 本质上和前一种一样，就不实现了

// 1.
//118/118 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 59.49 % of golang submissions (2.4 MB)
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n==0 {return ""}
	if n==1 {return strs[0]}

	prefix, k := strs[0], len(strs[0])
	for i:=1; i<n; i++ {
		if len(strs[i])<k {k = len(strs[i])}
		for j:=0; j<k; j++ {
			if prefix[j] != strs[i][j] {k = j}	// j是第一个不匹配的位置，k是前缀长度
		}
	}
	return prefix[:k]
}

// 水平扫描
// 这种做法就像是所有字符串都去从第一个位置匹配，向后统一平推
func longestCommonPrefix2(strs []string) string {

	var index, indexOut, j int
	l := len(strs)
	if l == 0 {
		goto RESULT
	}

	// 1.求出indexOut，最大公共长度
	indexOut = len(strs[0])
	for j=1;j<l;j++ {
		if len(strs[j]) < indexOut {
			indexOut = len(strs[j])
		}
	}
	//fmt.Println("最长公共长度：", indexOut)

	// 2.求出最长前缀
	for {

		// 有一种可能：第一个字符串长度比最短字符串长度要大
		// indexMax = l-1
		if index == indexOut {
			goto RESULT
		}

		// 遍历找出相同字符
		for j=0;j<l;j++ {
			if strs[j][index] != strs[0][index] {
				goto RESULT
			}
		}
		// 若全都和第一个字符串该位相等，则index前移，说明0～index-1范围都是相同前缀
		index++
		//fmt.Println("index此刻值为：", index)
	}
	//fmt.Println("index最终值为：", index)

	// 不要直接放在for循环中，只要前面没有return等退出语句，它始终会执行一次
RESULT:
	// 若没有公共前缀，则index=0，应返回空字符串而不是用index去取子串
	if index == 0 {
		return ""
	}
	// 返回最长前缀，单个字符串的0～index范围的子串
	return strs[0][:index]

}


// 3.前面的水平扫描写得太啰嗦，简化一下
// 第一种解法的问题在于，如果数组末尾有个很短的字符串，
// 但是找公共前缀的过程中还是会比较前面全部的字符串
// 水平扫描的思想可以更快检查出这种情况
func longestCommonPrefix3(strs []string) string {
	n := len(strs)
	if n==0 {return ""}
	if n==1 {return strs[0]}

	for i:=0; i<len(strs[0]); i++ {
		// 字符会从前向后一个一个拿出来和其他所有字符串比较
		for j:=1; j<n; j++ {
			if i==len(strs[j]) || strs[0][i] != strs[j][i] {return strs[0][:i]}
		}
	}
	return strs[0]
}


// 官方题解还给出了：
// 分治解法(不断将区间二分，求子区间的最长公共子串，不断向上汇总)、
// 二分查找解法(先找最短字符串(长度k)，在[0,k]区间使用二分查找不断试探)
// 这两种解法效率上都不如前面的扫描做法

// 官方题解还进行了拓展
// 如果要求一个可变的字符串q与某固定的字符串数组S的最长公共子串。
// 意味着这样的求最长公共子串函数经常被调用，那么优秀的做法是将S构建成字典树(前缀树)trie甚至是压缩前缀树

func LargestCommonPrefix(S []string, q string) string {
	if len(S)==0 {return ""}
	if len(S)==1 {return S[0]}

	// 将S构建字典树
	trie := &Trie{lt208.Constructor()}
	for _, s := range S {
		trie.Insert(s)
	}
	// 查询q与S最长公共前缀
	return trie.LongestCommonPrefix(q)
}





// 字典树的实现见lt208
// 这里给Trie增加方法






// 应用字典树求最长公共前缀。那么在公共前缀的部分，必然是没有分叉的
// 也就是这一段(公共前缀)每一个节点的子节点列表中都必须保证只有一个是非空的
// 为了更方便的查询每个节点的非空子节点数，设一size.
// 【其实如果是要方便记录非空节点数的话，用哈希集合代替数组会更方便】
// 通过数组实现的字典树其实就是26叉树
type TrieNode struct {
	next []*TrieNode	// 下一层
	R int		// r = 26，用来让每个节点初始化next
	size int	// 非空子节点数
	isEnd bool	// 标志单词结尾
}
func NewTrieNode() *TrieNode {
	tnode := &TrieNode{}
	tnode.next = make([]*TrieNode, tnode.R)
	return tnode
}
func (tnode *TrieNode) Size() int {return tnode.size}
func (tnode *TrieNode) IsEnd() bool {return tnode.isEnd}
func (tnode *TrieNode) SetEnd() {tnode.isEnd = true}
func (tnode *TrieNode) Get(ch byte) *TrieNode {
	if ch<'a' || ch>'z' {return nil}
	return tnode.next[ch-'a']
}
func (tnode *TrieNode) Put(ch byte)	{
	if ch<'a' || ch>'z' {return}
	if tnode.next[ch-'a']==nil {
		tnode.next[ch-'a'] = NewTrieNode()
		tnode.size++	// 非空子节点+1
	}
}
func (tnode *TrieNode) Has(ch byte) bool {
	if ch<'a' || ch>'z' {return false}
	return tnode.next[ch-'a'] != nil
}
func (tnode *TrieNode) DelSubNode(ch byte) error {
	if ch<'a' || ch>'z' {return errors.New("illegal input char")}
	if !tnode.Has(ch) {return errors.New("nonexistent char")}
	tnode.next[ch-'a'] = nil
	return nil
}
func (tnode *TrieNode) DisEnd(ch byte) error {
	if ch<'a' || ch>'z' {return errors.New("illegal input char")}
	if !tnode.Has(ch) {return errors.New("nonexistent char")}
	tnode.Get(ch).isEnd = false
	return nil
}

// Trie
type Trie struct {
	root *TrieNode
	// size等其他属性这里忽略，不设计
}

func NewTrie() *Trie {return &Trie{root:NewTrieNode()}}
func (trie *Trie) Root() *TrieNode {return trie.root}

func (trie *Trie) Insert(word string) {
	node := trie.root
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		if !node.Has(cur) {
			node.Put(cur)	// 没有就放进去
		}
		node = node.Get(cur)
	}
	node.SetEnd()	// 最后记得给单词结尾
}

// 把word当做前缀去树里找
// 不存在，则返回nil
// 存在且为自己，或者存在且为其他单词前缀，则返回末尾节点
func (trie *Trie) SearchPrefix(word string) *TrieNode {
	node := trie.root
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		if !node.Has(cur) {return nil}
		node = node.Get(cur)
		// 这时不需要关心node.IsEnd
	}
	return node
}

// 搜索单词是否存在于前缀树
func (trie *Trie) Search(word string) bool {
	node := trie.SearchPrefix(word)
	return node!=nil && node.IsEnd()
}

// 搜索word是不是字典树中的某个单词的前缀或本身
// 与search的区别在于，startwith只在乎是否存在
func (trie *Trie) StartWith(word string) bool {
	node := trie.SearchPrefix(word)
	return node != nil
}

// 删除某个单词
func (trie *Trie) Delete(word string) error {
	node := trie.root
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		if !node.Has(cur) {
			return errors.New("no such word")
		}
		node = node.Get(cur)
	}
	// 现在node是word[len(word)-1]的节点
	// 但是得先检查是不是isEnd，是才说明有word
	if !node.isEnd {return errors.New("no such word")}

	// 现在要判断它是叶子节点还是不是
	if node.size > 0 {	// 不是叶子节点
		node.isEnd = false		// 只需要置否
	} else {	// 叶子节点
		// 需要从叶子节点一直往上删直到有一个节点n
		// 它具有两个及以上非空子节点
		// 或者它的isEnd是true
		//node.isEnd = false	// 先把当前这个节点置为false
		//for node.size==1 && !node.isEnd {
		//
		//}
		//TODO: 这里删除操作，要么将TrieNode再加一个prev指针指向上层节点(双向链表)
		// 要么又从头遍历，暂时不知道怎么优雅的实现，先搁置
	}
}


// 查找word与Trie树当前状态的最长公共前缀
func (trie *Trie) LongestCommonPrefix(word string) string {
	node := trie.Root()
	prefix := make([]byte, 0, len(word))
	var cur byte
	for i:=0; i<len(word); i++ {
		cur = word[i]
		// 这三点条件保证了是公共前缀
		if node.Has(cur) && node.Size() == 1 && !node.IsEnd() {
			prefix = append(prefix, cur)
			node = node.Get(cur)
		} else {
			return string(prefix)
		}
	}
	return string(prefix)
}