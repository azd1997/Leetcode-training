package lt127

// 单词接龙

// 这种题一般是用图论考虑，转化成图上节点，看起点能否走到最终，最短距离几何
// BFS/DFS均可

// 为了描述单词与单词之间可达，需将单词进行预处理：
// Hot -> [*ot H*t Hot](暂且称Hot的next列表)
// 当另一个单词的next列表与Hot的next列表存在重叠时，说明二者可达
// 有一点要注意：如果本身这一位就相同(例如Hot与HugDE第一位)，
// 那便没必要进行next数组的匹配

// 当然每一个单词对应一个next列表是可行的，但是浪费太多空间
// 可以只使用一个哈希表，所有next列表的元素作键，对应的若干单词作值
// 形成 H*t -> Hit/Hot 这样的映射

// 1. BFS O(n*L) / O(n*L)	n为单词长度,L为wordList长度
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	// 1. 由于题目描述可以假设beginWord与endWord不同，且wordList非空
	// 那么这里不作检查

	n := len(beginWord)		// 每个单词的长度

	// 2. wordList预处理，得到前面说的 next状态 -> 原单词 的映射
	set := make(map[string]*[]string)	// 这里最好接数组指针
	var nextword string
	for _, word := range wordList {
		for i:=0; i<n; i++ {
			nextword = word[:i] + "*" + word[i+1:]
			if rawwords, ok := set[nextword]; ok {
				*rawwords = append(*rawwords, word)
			} else {
				set[nextword] = &[]string{word}
			}
		}
	}

	// 3. 构造辅助队列，队列存的值为从起点到当前字符串走了几步
	type Elem struct {
		word string
		step int	// 从beginWord走到word花了几步
	}
	queue := []Elem{{beginWord, 1}}

	// 4. 构造访问哈希表，用以确保不重复处理相同的单词
	visited := make(map[string]bool)
	visited[beginWord] = true

	// 5. BFS
	for len(queue) != 0 {
		node := queue[0]; queue = queue[1:]		// 出队

		// 遍历node.word的所有next状态
		for i:=0; i<n; i++ {
			nextword := node.word[:i] + "*" + node.word[i+1:]
			// 遍历每一个next状态对应的rawword数组，此时都是当前node.word的邻接单词
			// 注意next不一定一定对应有rawwords数组，需要检查(这是因为beginWord不一定在wordList中)
			if _, ok := set[nextword]; !ok {continue}
			for _, adjacent := range *set[nextword] {
				// 下一步就可以变换到endWord
				if adjacent == endWord {
					return node.step + 1
				}
				// 否则需要将步数+1，继续压入队列中;
				// 并且还需将该单词标记为访问过，访问过的单词就没必要继续入队了
				if !visited[adjacent] {
					visited[adjacent] = true
					queue = append(queue, Elem{adjacent, node.step + 1})
				}
			}
		}
	}
	return 0	// 不可达(不可达的具体原因:end不在wordlist或者wordlist缺少中间单词，这里不关心)
}


// 2. 双向BFS O(n*L) / O(n*L)
// 参考官方题解描述，此题中按wordList构建的图可能很大，双向BFS可有效缩减搜索时间与空间复杂度
// 双向BFS的终止条件为从头开始的word和从尾开始的word，他们两个的adjacent相等
// 由于直接在原BFS代码处叠加再写一份显得代码过于冗长，将之提取为一个函数，用于求
// 某一个节点的adjacent以及是否达到双向BFS终止条件
// 此外，为了在相遇时能够知道总工走了多少步，visited表应该记录走的步数，而不是仅仅记录是否访问过
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	// 1. 检查wordList是否含有endWord，不含直接返回0
	// 双向BFS必须要先检查endWord存在与否，单向则可以不需要
	endExisted := false
	for _, word := range wordList {
		if word == endWord {endExisted = true}
	}
	if !endExisted {return 0}

	n := len(beginWord)		// 每个单词的长度

	// 2. wordList预处理，得到前面说的 next状态 -> 原单词 的映射
	set := make(map[string]*[]string)	// 这里最好接数组指针
	var nextword string
	for _, word := range wordList {
		for i:=0; i<n; i++ {
			nextword = word[:i] + "*" + word[i+1:]
			if rawwords, ok := set[nextword]; ok {
				*rawwords = append(*rawwords, word)
			} else {
				set[nextword] = &[]string{word}
			}
		}
	}

	// 3. 构造辅助队列
	queueBegin := []Elem{{beginWord, 1}}
	queueEnd := []Elem{{endWord, 1}}

	// 4. 构造访问哈希表，用以确保不重复处理相同的单词
	// 记录的值为相应的begin、end走到当前node的步数
	visitedBegin := make(map[string]int)
	visitedEnd := make(map[string]int)
	visitedBegin[beginWord], visitedEnd[endWord] = 1, 1

	// 5. BFS
	for len(queueBegin) != 0 && len(queueEnd) != 0 {
		// 通俗地讲，就是两边都看一看自己下一步能否与对方相遇，能则返回，不能则继续往前一步

		// begin开始的BFS
		ans := visitNode(set, &queueBegin, visitedBegin, visitedEnd)
		if ans > -1 {
			return ans
		}
		// end开始的BFS
		ans = visitNode(set, &queueEnd, visitedEnd, visitedBegin)
		if ans > -1 {
			return ans
		}
	}

	return 0	// 不可达(不可达的具体原因:end不在wordlist或者wordlist缺少中间单词，这里不关心)
}

type Elem struct {
	word string
	step int	// 从beginWord或者endWord走到word花了几步
}

func visitNode(set map[string]*[]string, queue *[]Elem, visited, otherVisited map[string]int) int {
	node := (*queue)[0]; *queue = (*queue)[1:]
	n := len(node.word)
	// 遍历node.word的所有next状态
	for i:=0; i<n; i++ {
		nextword := node.word[:i] + "*" + node.word[i+1:]
		// 遍历每一个next状态对应的rawword数组，此时都是当前node.word的邻接单词
		// 注意next不一定一定对应有rawwords数组，需要检查(这是因为beginWord不一定在wordList中)
		if _, ok := set[nextword]; !ok {continue}
		for _, adjacent := range *set[nextword] {
			// 下一步就和另外一端开始的BFS相遇
			if v, ok := otherVisited[adjacent]; ok {	//
				return node.step + v
			}
			// 否则需要将步数+1，继续压入队列中;
			// 并且还需将该单词标记为访问过，访问过的单词就没必要继续入队了
			if _, ok := visited[adjacent]; !ok {
				visited[adjacent] = node.step + 1
				*queue = append(*queue, Elem{adjacent, node.step + 1})
			}
		}
	}

	return -1	// 当前节点下一步没法与另一端的BFS相遇
}