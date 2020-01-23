package lt28

// 实现 strStr() 函数。
//
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1:
//
//输入: haystack = "hello", needle = "ll"
//输出: 2
//示例 2:
//
//输入: haystack = "aaaaa", needle = "bba"
//输出: -1
//说明:
//
//当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
//对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-strstr
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考，这其实就是字符串匹配！！
// 是时候练习一波 BF、RK、BM、KMP算法了！！

// 1. BF Brute Force 暴力匹配 （双指针，匹配两串头部）
// 常用，因为：
// 1. 平时匹配字符串长度不会太长
// 2. 实现简单，不易出错，容易修复。KISS（Keep it Simple and Stupid）原则
// 最坏时间复杂度O(mn) aaaaa 匹配aaab
// 平均时间复杂度O(mn) 空间O(1)
// 74/74 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 69.03 % of golang submissions (2.3 MB)
func strStr1(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {return 0}		// 空模式串可随意匹配
	if m < n {return -1}

	var finded bool

	for i:=0; i<=m-n; i++ {
		if haystack[i] == needle[0] {
			finded = true
			for j:=1; j<n; j++ {
				if haystack[i+j] != needle[j] {
					finded = false
					break
				}
			}
			if finded {return i}
		}
	}

	return -1
}

// 2. BF的另一种实现方式
//74/74 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 69.03 % of golang submissions (2.3 MB)
func strStr2(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {return 0}		// 空模式串可随意匹配
	if m < n {return -1}

	var i, j int
	for i<m && j<n {
		if haystack[i] == needle[j] {
			i, j = i+1, j+1
		} else {	// 一遇到不重复的，就把i指针回退到模式串头部对齐位置，j回退到0
			i, j = i-j+1, 0
		}
	}

	if j == n {return i-j}	// 说明模式串全匹配到了

	return -1
}

// 3. RK Rabin-Karp算法
// BF算法每次右移都需要重新尝试匹配子串，用的是循环比较字符的方式；
// RK则利用哈希算法加速这个比较过程：对模式串、以及主串的 m-n+1 个子串取哈希，然后进行哈希比较
// 因为哈希是数字，数值比较比字符比较要快的多，所以比较效率就上去了
// 但是由于计算哈希的过程仍需遍历子串所有字符，所以整体效率并没有提高
// 要想提高效率，就需要提高计算子串哈希的效率，这需要设计巧妙地哈希算法。

// 3.1 RK-1 假设所有字符都是英文小写字母(0~25)，将字符串表示为26进制数再转为10进制数
// 所得的数（哈希）一定不会产生冲突。我们先这么实现（尽管模式串长度一旦较长，数据表示就一定会溢出）
//74/74 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 69.03 % of golang submissions (2.3 MB)
func strStr3(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {return 0}		// 空模式串可随意匹配
	if m < n {return -1}

	needleHash := hashChar26(needle)

	for i:=0; i<=m-n; i++ {
		if hashChar26(haystack[i:i+n]) == needleHash {return i}
	}

	return -1
}

// 不使用大数的情况下，uint64差不多是最大范围
func hashChar26(str string) uint64 {
	var hash uint64
	for i:=0; i<len(str); i++ {
		hash = hash*26 + uint64(str[i]-'a')		// 'a' = 97; 'z' = 122
	}
	return hash
}

// 3.2 RK-2
// 在RK-1实现中有两个点可以优化：
// 一个是主串中的前后两个连续子串，计算哈希时重复计算了许多，例如 abc 和 bcd 中间是重复的（前一个*26就好）
// 另一个是 26^1,...,26^(n-1)可以提前计算好
//74/74 cases passed (476 ms)
//Your runtime beats 5.36 % of golang submissions
//Your memory usage beats 15.49 % of golang submissions (2.6 MB)
// 感觉这个运行时间过分了啊，怎么回事？
func strStr3_1(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {return 0}		// 空模式串可随意匹配
	if m < n {return -1}

	n26 := generate26n(n)
	needleHash := hashChar26(needle)

	// 比较第一个子串
	var prevHash, tmp uint64
	prevHash = hashChar26(haystack[0:n])
	var i int
	for {
		//fmt.Println("i=", i)
		if prevHash == needleHash {return i}
		if i + n == m {break}		// 遍历到最后一个子串
		i++
		tmp = prevHash - uint64(haystack[i-1]-'a')*n26[n-1]
		prevHash = tmp*26 + uint64(haystack[i+n-1]-'a')		// 计算下一个子串的哈希
	}

	return -1
}

// 要求26^(n-1)不会超出表示范围。
func generate26n(n int) []uint64 {
	res := make([]uint64, n)
	for i:=0; i<n; i++ {
		res[i] = pow(26, i)
	}
	return res
}

// 指数运算，不想使用math.Pow
func pow(x uint64, y int) uint64 {
	var res uint64 = 1
	for i:=0; i<y; i++ {
		res = res * x
	}
	return res
}

// 由于这个哈希算法的设计使得只需要遍历一遍主串就可以计算出所有哈希。时间复杂度O(m)	// 这里m指主串规模
// (上面的实现不是一次性计算所有子串哈希，而是分开计算，其实也是一种小优化)
// 而模式串与子串哈希比较也只需要比较m-n+1次
// 所以整体时间复杂度变成O(n)
// 上面之所以时间比较低，可能是涉及了过多的类型转换，暂时不纠结

// 上面这个哈希算法 优点是不可能碰撞，缺点是很容易溢出，超出数值表示范围
// 那使用这样的哈希算法呢？
// a~z 对应数字 0~25. 将字符代表的数字相加得到哈希。
// 显然碰撞几率大了许多，但是很难溢出（一般情况下都不可能溢出）
// 使用这样的哈希算法怎么办呢？
// 哈希相等的时候再进行逐字符的比较呗！
// 或者将 每个字母 映射为 质数， 再取其和， 那么碰撞的概率又降低了， 而且也很难溢出
// 只要所设计的哈希算法碰撞概率不过高，使复杂度退化到最差情况，都是可以的（相较于BF有优化）


// 4. BM Boyer-Moore算法 最快的字符串匹配算法
// 各种文本编辑器的查找、正则表达式匹配实现可能都是它，但比较复杂

// 先介绍下BM的几个规则

// 1. BM算法匹配模式串与子串时是从右模式串右端开始匹配，也就是倒序匹配
// 2. 坏字符规则：
// 		2-1. 最右不匹配的主串中字符 称为 坏字符
// 		2-2. 若坏字符在模式串中不存在，则直接将模式串向后滑动n位，因为中间都不可能匹配这个坏字符
//		2-3. 若坏字符在模式串中存在，取其 最右出现位置， 记x, 坏字符对应的模式串位置为 s ， 则直接将模式串向后滑动 s-x 位
// 		2-4. 利用坏字符规则，BM最好时间复杂度可以达到O(m/n)
// 		2-5. 坏字符规则的问题是 s - x 有可能为负数或0， 这就会使得模式串倒滑或者无限循环卡在那
// 		2-6. 由坏字符规则得到滑动距离 l1
// 3. 好后缀规则：
//		3-1. 坏字符之后的可能有0个、1个或多个已经匹配的字符，如非0个，称为 好后缀
//		3-2. 将好后缀到模式串坏字符位置前面去找，如果找到，则将模式串滑到最右匹配好后缀位置与好后缀对齐的位置； 没找到则直接滑到好后缀之后
// 		3-3. 3-2的问题是有可能好后缀与模式串前部不匹配，但是好后缀的后缀子串可能和模式串的前缀子串匹配，这个也需要考虑进去
// 		3-4. 因此当好后缀不能匹配到时要继续检查好后缀的前缀子串与模式串的前缀子串，找出最长公共子串，再计算滑动距离
//		3-5. 好后缀规则不存在错过匹配的可能，但它比坏字符规则要慢一些
// 		3-6. 由好后缀规则得到滑动距离 l2
// 4. 最终的滑动距离就是 max(l1, l2)
// 5. 加速模式串中最右出现坏字符的查找：模式串预填充哈希表bc, 记录模式串中各字符最右出现位置
// 6. 加速好后缀规则中好后缀后缀子串与模式串前缀字串的匹配：
// 		6-1. 预处理模式串前缀子串（或者说好后缀后缀子串（就是模式串后缀子串）），
// 		6-2. 用一哈希表suffix，键为后缀子串长度，值为后缀子串在模式串前面出现的位置； 用数组也可以替代哈希表
//		6-3. 后缀子串在前面有多个匹配怎么办？不能只存最后后匹配位置，那样仍可能过度滑动； 还要考虑最长匹配前缀
// 		6-4. 相应的，用以哈希表或数组存bool值，标记模式串后缀子串是否是模式串的前缀子串
// 		6-5. 只需要选择prefix[x]=true且suffix[x]最大的x去计算滑动距离即可
// 		6-6. 那么怎么计算suffix/prefix呢？暴力匹配肯定是低效的，
//		6-7. 遍历模式串，用模式串[0:i]与模式串取公共后缀，若能取到，则记录suffix中的值（并且会不断更新，找到最右的位置）
//		6-7. 取到公共后缀的同时，看下公共后缀的首字母在模式串的下标是不是0，是说明是前缀子串，prefix中值设为true
//		6-8. 如何计算滑动距离？如果suffix[i]!=-1，则说明好后缀完整存在与模式串中，那么滑移 j-suffix[i]+1位
// 		6-9. 如果suffix[i]==-1再判断 k = m-r; prefix[k]=true, 滑移r位

// 下面来实现它

// 由于经过前面的测试，测例字符串只包含英文小写字母，
// 所以所有需要用到哈希表的地方都可以用长度为26的数组来模拟

// 首先是generateBC BC(Bad Char) 对模式串做处理
func generateBC(str string) (bc []int) {
	bc = make([]int, 26)
	for i:=0; i<26; i++ {
		bc[i] = -1
	}
	for i:=0; i<len(str); i++ {
		bc[str[i]-'a'] = i
	}
	return bc
}

// 接着实现基于坏字符规则的BM算法
func BM1(str1, str2 string) int {
	// 预处理模式串得到 bc
	bc := generateBC(str2)

	var i, j int	// i为主串指针，从主串首部向后移动； j为模式串指针
	m, n := len(str1), len(str2)

	for i <= m-n {
		i = 0
		for j:=n-1; j>=0; j-- {
			if str1[i+j] != str2[j] {break}		// 最右坏字符位置 j
		}
		if j < 0 {return i} 	// 说明匹配到了，返回子串首部位置

		// 坏字符规则计算出的后滑距离
		i = i + (j - bc[str1[i+j]-'a'])
	}

	return -1
}

// 接下来实现 generateGS GS(Good Suffix)
func generateGS(str string) (suffix []int, prefix []bool) {
	n := len(str)
	suffix, prefix = make([]int, n), make([]bool, n)
	for i:=0; i<n; i++ {suffix[i] = -1}

	var j, k int	// j为倒序匹配的指针； k为公共子串长度
	for i:=0; i<n-1; i++ {		// 用以得到子串 str[0:i]
		j, k = i, 0
		for j>=0 && str[j] == str[n-1-k] {	// str[0:i]与str[0:n]求公共后缀
			j, k = j-1, k+1
			suffix[k] = j+1		// 最后在此停止时，j 已经指向公共后缀位置前一个，所以要加1
		}
		if j == -1 {prefix[k] = true}	// 说明公共后缀子串也是模式串前缀子串
	}

	return suffix, prefix
}

// 计算好后缀规则下的后滑距离.	j为坏字符对齐的模式串位置， n为模式串长度
func moveByGS(j, n int, suffix []int, prefix []bool) int {
	k := n-1 - j		// 好后缀长度
	if suffix[k] != -1 {return j-suffix[k]+1}	// 好后缀存在于模式子串非后缀位置
	for r:=j+2; r<=n-1; r++ {
		if prefix[n-r]==true {return r}
	}
	return n
}

// 最后，加上好后缀，完成BM算法 str1主串， str2模式串
func BM(str1, str2 string) int {
	// 预处理模式串得到 bc
	bc := generateBC(str2)
	// 预处理模式串得到 suffix, prefix
	suffix, prefix := generateGS(str2)

	var i, j int	// i为主串指针，从主串首部向后移动； j为模式串指针
	m, n := len(str1), len(str2)

	for i <= m-n {
		for j=n-1; j>=0; j-- {
			if str1[i+j] != str2[j] {break}		// 最右坏字符位置 j
			//fmt.Printf("j = %d\n, str1[%d+%d] = %s\n", j, i, j, string(str1[i+j]))
		}

		//fmt.Printf("j = %d\n, str1[%d+%d] = %s\n", j, i, j, string(str1[i+j]))
		//fmt.Printf("j = %d\n", j)

		if j < 0 {return i} 	// 说明匹配到了，返回子串首部位置

		// 坏字符规则计算出的后滑距离
		d1 := j - bc[str1[i+j]-'a']

		// 基于好后缀规则计算后滑距离
		var d2 int
		if j < n-1 {	// 说明存在好后缀
			d2 = moveByGS(j, n, suffix, prefix)
		}

		// 合并求滑移距离
		d := d1	// 滑移距离
		if d2 > d1 {d = d2}

		// 向后滑移
		i = i + d
	}

	return -1
}

// 实际的 BM算法 实现还有一些别的优化，现在这个简化版本在极端情况下（比如模式串为aaaa这种包含很多重复的字符）
// 预处理计算suffix、prefix时间复杂度就会退化成 O(n2)

// 有时候为了降低预处理内存消耗，会只使用好后缀规则实现BM。


// 最后将BM应用到题目
//74/74 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 15.49 % of golang submissions (2.6 MB)
func strStr4_BM(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {return 0}		// 空模式串可随意匹配
	if m < n {return -1}

	return BM(haystack, needle)
}


// 5. KMP Knuth-Morris-Pratt算法
// 设主串是a， 模式串b
// KMP和BM算法一样，也是寻找规律使得模式串能够一次性向后多滑动几步
// KMP 也有两个概念： 坏字符（和BM一样）、好前缀（从前向后遍历，已匹配的部分叫好前缀）
// KMP的思想是： 当遇到坏字符且有好前缀的时候，寻找 好前缀的后缀子串 与 模式串前缀子串 求最长公共子串，再据此向后滑移
// 				而好前缀一定是模式的前缀子串
// 这其实是当定位到 最左坏字符 时，设其在模式串位置j	（从它取最左坏字符就知道比BM慢了）
// 那么 b[0:j] 就是 好前缀
// 现在要找到 b[0:j]的后缀子串与前缀子串匹配的最长情况
// 分别称为 好前缀的最长可匹配 后缀与前缀
// 那么 b[0:j] 就是取 b的所有前缀子串 （除了本身）
// 因此可以预处理模式串
// 同样预处理的方法和前面BM好后缀的处理类似，也是以子串长度为索引，构建数组哈希表
// 这样一个数组称为 next， 很多书称这个next为失效函数
// next存的是 模式串前缀子串 （好前缀的候选）的 最长可匹配前缀子串 末尾位置 （如果找不到最长可匹配，则是-1）

// 假设这样的 next数组已经获得，那么 KMP算法实现如下：
func fakeNext(b string) []int {
	next := make([]int, len(b))
	return next
}

func BM_NextAlreadyExists(a, b string) int {
	next := fakeNext(b)

	var j int	// 坏字符位置
	for i:=0; i<len(a); i++ {
		for j>0 && a[i]!=b[j] {		// a[i]!=b[j]不匹配时，j需要回退到好前缀的最长可匹配前缀子串末尾后一位（可匹配的部分已匹配过，无需再匹配））
			j = next[j-1] + 1		// j-1就是好前缀长度， +1是因为前面已匹配的无需再比
		}
		if a[i] == b[j] {j++}		// 如果两指针对应的位置字符匹配，则将j指针后移。若j后移一位之后到达模式串长度，那说明完全匹配
		if j == len(b) {return i-len(b)+1}	// 完全匹配，就返回匹配子串首部位置
	}

	return -1
}

// 接下来就是计算 next数组 的过程
// 最简单的方法就是: 比如计算 next[4]， 就把 b[0:4] 的所有后缀子串与所有相应长度的前缀子串匹配，找出最长可匹配的情况
// 但太低效
// KMP使用了类似动态规划的做法：
// 1. 按下标从小到大， 计算next值。当计算到 next[i]时，next[0]~next[i-1]已经得到
// 如果 next[i-1] = k-1, 也就是说 b[0:k)子串是b[0:i)的最长可匹配子串
// 2. 如果子串 b[0:k) 的下一个字符 b[k] 和 b[0:i] 下一个字符 b[i] 匹配，
// 那么说明 b[0:k+1) 是 b[0:i+1) 的最长可匹配前缀子串，所以 next[i] = k
// 3. 如果b[k]与b[i]不匹配:
// 		假设b[0:i)最长可匹配后缀子串是 b[r:i)， 所以 b[r:i-1)必然是 b[0:i-1)可匹配后缀子串，但不一定是最长。
//		既然 b[0:i-1) 最长可匹配后缀子串 对应的模式串的 前缀子串 的下一个字符不等于 b[i]
//		那么就可以考察 b[0:i-1) 的 次长可匹配的



// TODO: KMP算法