package lt315

import "sort"

// 计算右侧小于当前元素的个数

// 1. 暴力法 O(n2)/O(1)	(空间复杂度不算返回数组)
// 反正这题暴力思路很简单，先实现一下
func countSmaller1(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	res := make([]int, n)
	for i:=n-1; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			if nums[j] < nums[i] {res[i]++}
		}
	}
	return res
}


// 2. 动态规划的思想优化暴力解法  或者称 记忆化暴力解
// 空间复杂度仍为O(1)，但时间复杂度最好O(n)，最差O(n2)
func countSmaller2(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	res := make([]int, n)
	res[n-1] = 0
	for i:=n-2; i>=0; i-- {
		// 从自身出发，向右寻找第一个小于等于自己的元素d
		firstNotLarger := i+1		// 右边第一个<=自己的元素的下标
		for firstNotLarger < n-1 {	// 一直遍历到n-2这个元素，如果这个元素还比nums[i]大，那么显然nums[i]就和n-1比了
			if nums[firstNotLarger] <= nums[i] {break}
			firstNotLarger++
		}

		if nums[i]==nums[firstNotLarger] {res[i] = res[firstNotLarger]; continue}
		if nums[i]>nums[firstNotLarger] {res[i] = res[firstNotLarger] + 1; continue}
	}
	return res
}

// 上面的思路是错的
// i<j .  dp[j]右侧比dp[j]大的数不一定比dp[i]大


// 下面参考题解区 Adam Wong和liweiwei1419等题解


// 3. 暴力模拟法 + 二分查找  O(nlogn)/O(n)
// 后序遍历时，维护一个已排序数组sortedNums, 每遍历一个元素，
// 就把该元素添加到sortedNums中(保持升序，便于计算结果，降序当然也是行得通的)
// 这样，每次计算nums[i]的右侧更小数数目时直接在sortedNums中二分查找
func countSmaller3(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	sortedNum, res := make([]int, 0), make([]int, n)
	pos := 0	// 应插入位置
	for i:=n-1; i>=0; i-- {
		pos = binarySearchAndInsert(&sortedNum, nums[i])
		res[i] = pos	// pos是升序的sortNums应插入位置，小于nums[i]的元素就有pos个
	}
	return res
}

// 在升序排列的数组中查找target的应放入位置
// 要注意的是，如果发现a=target，target要放在a之前
// 例如 arr=[1,3,5,6], target=4，arr变为[1,3,4,5,6]，返回2
func binarySearchAndInsert(arr *[]int, target int) int {
	n := len(*arr)
	if n==0 {
		*arr = append(*arr, target)
		return 0
	}

	l, r, mid := 0, n-1, 0
	for l<=r {
		mid = (l+r)/2

		// target应插入第一个位置的情况
		if mid==0 && (*arr)[mid] >= target {
			*arr = append([]int{target}, *arr...)
			return 0
		}
		// target追加到最后的情况
		if mid==n-1 && (*arr)[mid] < target {
			*arr = append(*arr, target)
			return n
		}
		// mid在中间(0, n-1]，并恰好是target应插入位置
		if mid>0 && (*arr)[mid-1] < target && (*arr)[mid] >= target {
			*arr = append((*arr)[:mid+1], (*arr)[mid:]...)
			(*arr)[mid] = target
			return mid
		}
		// target应插入到mid左侧(保留mid)
		if (*arr)[mid] >= target {
			r = mid; continue
		}
		// target应插入到mid右侧(不保留mid)
		if (*arr)[mid] < target {
			l = mid+1; continue
		}
	}
	return l
}


// 4. 暴力模拟法 + 二分查找 (二分查找过程代码优化版)
func countSmaller4(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	sortedNum, res := make([]int, 0), make([]int, n)
	pos := 0	// 应插入位置
	for i:=n-1; i>=0; i-- {
		pos = binarySearchAndInsert(&sortedNum, nums[i])
		res[i] = pos	// pos是升序的sortNums应插入位置，小于nums[i]的元素就有pos个
	}
	return res
}

func binarySearchAndInsert2(arr *[]int, target int) int {
	n := len(*arr)
	if n==0 {
		*arr = append(*arr, target)
		return 0
	}

	l, r, mid := 0, n, 0	// 注意r初始为n
	for l<r {
		mid = (l+r)/2

		if (*arr)[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 结束之后l=r就是应插入位置
	*arr = append((*arr)[:mid+1], (*arr)[mid:]...)	// 能够处理插入到头部和尾部的情况
	(*arr)[r] = target
	return r
}


// 5. 二叉搜索树 O(nlogn)/O(n)
type BSTNode struct {
	val, count int	// val为当前值，count为其右小于该值的数目
	left, right *BSTNode
}

func NewBSTNode(val int) *BSTNode {
	return &BSTNode{val:val}	// count零值为0
}

func (node *BSTNode) Insert(insertNode *BSTNode, count *int) {

	if node.val >= insertNode.val {
		// 插入的结点值更小，则被比较结点的计数count++，然后将插入到左子树(若不为空)
		node.count++
		if node.left != nil {
			node.left.Insert(insertNode, count)
		} else {
			node.left = insertNode
		}
	} else {
		// 插入的结点更大，需要在右子树(若不为空)继续找
		*count += node.count + 1
		if node.right != nil {
			node.right.Insert(insertNode, count)
		} else {
			node.right = insertNode
		}
	}
}

// 利用平衡二叉树，每个节点在插入树中记录比自己小的结点总数
func countSmaller5(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	res := make([]int, n)
	res[n-1] = 0
	node := NewBSTNode(nums[n-1])	// 先生成最后一个节点(或者说第一个)
	for i:=n-2; i>=0; i-- {
		node.Insert(NewBSTNode(nums[i]), &res[i])
		// 插入结束后，res[i]即被赋值为结果
	}
	return res
}


// 6. 利用归并排序 O(nlogn)
// TODO: 没太看懂，但是在归并过程中不断开辟新数组，总数组开辟量O(nlogn)感觉不友好
func countSmaller6(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	// 答案数组
	res := make([]int, n)
	// 默认全为0

	// helper数组用于将i和nums[i]关联
	helper := make([][2]int, n)
	for i:=0; i<n; i++ {helper[i] = [2]int{i, nums[i]}}

	// 在归并排序的过程中为res填值
	mergeSort(&helper, &res)

	return res
}

func mergeSort(helper *[][2]int, res *[]int) {
	n := len(*helper)
	if n<2 {return}

	// 不断将区间二分
	mid := n/2
	helperLeft, helperRight := append([][2]int{}, (*helper)[:mid]...), append([][2]int{}, (*helper)[mid:]...)
	mergeSort(&helperLeft, res)
	mergeSort(&helperRight, res)

	// 触底后开始不断合并
	*helper = (*helper)[:0]		// 清空helper
	merge(&helperLeft, &helperRight, helper, res)
}

func merge(helperLeft, helperRight, helper *[][2]int, res *[]int) {
	n1, n2 := len(*helperLeft), len(*helperRight)
	i, j := 0, 0	// 分别为left和right的游标

	for i<n1 && j<n2 {
		if (*helperLeft)[i][1] <= (*helperRight)[j][1] {
			*helper = append(*helper, (*helperLeft)[i])
			(*res)[(*helperLeft)[i][0]] += j
			i++
		} else {
			*helper = append(*helper, (*helperRight)[j])
			j++
		}
	}

	for ; i<n1; i++ {
		*helper = append(*helper, (*helperLeft)[i])
		(*res)[(*helperLeft)[i][0]] += j
	}
	for ; j<n2; j++ {
		*helper = append(*helper, (*helperRight)[j])
	}
}


// 7. 使用树状数组 O(nlogn)/O(n)
// (1)离散化：将数组nums元素转化为其排名1~n的排名数组，使得数值连续，
// 可以在使用树状数组时节省空间. 例如 [5,2,6,1] => [3,2,4,1] 不影响本题求解
// (2)从后向前填表：
// 		1> 读到1，其排名为1，首先在树状数组"1"的位置+1，
// 		查询排名在1之前的数的时候显然是没有的，所有res[n-1]=0
//		2> 读到6，排名为4，首先在树状数组"4"的位置+1，然后在树状数组
//		中查询排名在"4"之前的元素个数有多少个，结果是1，所以nums[n-2]=1
//	https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/solution/shu-zhuang-shu-zu-by-liweiwei1419/
// TODO: 下次再研究
func countSmaller7(nums []int) []int {
	n := len(nums)
	if n==0 {return nil}

	// 答案数组
	res := make([]int, n)
	// 默认全为0

	nums2 := append([]int{}, nums...)
	sort.Ints(nums2)

	// 排名表
	ranks := make(map[int]int)
	for i:=0; i<n; i++ {ranks[nums2[i]] = i+1}

	// 构建树状数组，倒序填res表
	tree := NewFenwickTree(n)
	for i:=n-1; i>=0; i-- {
		// 1. 查询当前数的排名
		curRank := ranks[nums[i]]
		// 2. 在树状数组中curRank位置 + 1
		// 其实是保证curRank位置记录的是比自己小的排名的数量加上自身
		tree.update(curRank, 1)
		// 3. 查询一下 <= curRank-1 的元素有多少
		res[i] = tree.query(curRank-1)
	}

	return res
}


// 虽然树状数组的细节没太搞明白，但是应用于本题的总体思路是：
// for [n-1, 0] {
//		nums[i] => rank(nums[i])
// 		前缀和 = query(i) = A[1]+...+A[i]
// }
// 由于排名是独一无二的(当然那种相同数排名一致的这里也能处理，而且更省空间)
// 我这里实现是用sort.Ints排序得到排名，排名是独一无二的
// 当我遍历到nums[i]时，update就会置A[rank(nums[i])]为1，
// 然后更新rank(nums[i]及其后的前缀和
// 因此，当updata进去就立马query，就能得到A[1]到A[i]这个区间和
// 这个和也表示A[1:i]中1的个数，也就是<nums[i]的右侧元素个数
// 当然为了处理nums[i]存在重复元素的情况，查询的是curRank-1所对应的前缀区间和


type FenwickTree struct {
	tree []int
	length int
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		tree:   make([]int, n+1),
		length: n,
	}
}

// 单点更新：将index位置 + 1。
// index为当前元素(nums[i])对应的rank
// 单点更新只关心一个数变化了多少
func (t *FenwickTree) update(index, delta int) {
	// 从下到上，最多到size
	for index <= t.length {
		t.tree[index]  += delta
		index += t.lowbit(index)
	}
}

// 区间查询： 查询 <=index 的元素个数
// index为前缀区间的最大索引，返回前缀区间[0,index]的所有元素之和
// 查询的语义是“前缀和”: 一个数组从头开始的区间里所有元素的和
// 对于A=[-,1,2,3,4,5,6,7,8]
// 	  K= -,0,1,0,2,0,1,0,3
//    C=[-,1,2,1,4,1,2,1,8]		// 预处理结束后树状数组就是这样固定的
// 要注意的是，真正生成过程在update和query操作的过程中进行的
// 例如要查询前6个元素的前缀和，则sum = C[6]+C[4]，
// 也就是 sum([0,6]) = A[1]+...+A[6] = C[4] + C[6]
func (t *FenwickTree) query(index int) int {
	// 从右到左查询
	sum := 0
	for index > 0 {
		sum += t.tree[index]
		index -= t.lowbit(index)
	}
	return sum
}

// lowbit 计算  2^k
// x的二进制表示从右向左数有多少个0(记为k)，遇1则止
// 这个k值的意义在于：
// 树状数组其实是排名数组A=[0,1,2,3,..](第一个位置不要)的预处理数组
// C=[0,sum1,sum2,sum3,...](将前缀和标记为sum)。C被构建成多叉树的形式，所以叫树状数组
// k的意义在于： 记C下标为x，则2^k就是数组C中的元素来自数组A的个数。
// 对于A=[-,1,2,3,4,5,6,7,8]
// 	  K= -,0,1,0,2,0,1,0,3
//    C=[-,1,2,1,4,1,2,1,8]		// 预处理结束后树状数组就是这样固定的
func (t *FenwickTree) lowbit(x int) int {
	return x & (-x)
}























