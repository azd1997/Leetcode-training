package lt202

// 快乐数

//编写一个算法来判断一个数是不是“快乐数”。
//
//一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。
//
//示例: 
//
//输入: 19
//输出: true
//解释:
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/happy-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 1. 目前只能想到通过模10与除10得到所有位的数字，再去计算。最后要想是快乐数，必须是只有一个1的形式
// 		同时为了避免死循环，需要使用哈希集合等结构标记计算结果，如果计算结果重了就退出
//


// 还是看题解吧
//

// 1. 尾递归(参考Knife`)
// 1~4中只有1是快乐数，5之后的数要么回归到1要么回归到4或者3 （这属于找规律了...）
// 只有n>4时才进行递归，否则做判断
//401/401 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 97.65 % of golang submissions (2 MB)
func isHappy1(n int) bool {
	if n > 4 {
		// 递归
		sum := 0
		bit := 0
		for n > 0 {
			bit = n % 10
			sum += bit * bit
			n = n / 10
		}
		return isHappy1(sum)
	}

	return n==1
}


// 2. 使用哈希集和记录平方和sum，通过是否存在于sum来控制循环结束 (参考Shan)
// 说实话，这可能是现阶段我能想到的解法...
// 但这个解法并不好，内存易炸，前面的递归解法也一样，只是因为题目测例比较容易过
func isHappy2(n int) bool {
	set := make(map[int]bool)	// 集合，存sum
	sum := 0
	bit := 0

	for {
		// 计算sum
		for n > 0 {
			bit = n % 10
			sum += bit * bit
			n = n / 10
		}

		// 快乐
		if sum == 1 {return true}

		// 不快乐，检查sum是否出现过，决定是否继续下一次迭代
		if set[sum] {
			return false
		} else {
			// 标记sum,并进入下一次迭代
			set[sum] = true
			n = sum
			sum = 0
		}
	}
}

// 3. 快慢指针控制循环结束 (参考 金字塔下的小蜗牛)
// 当fast和slow相等之后，就是过了一次循环周期，此时要检查是不是因为1引起的循环，不是就是快乐数
func isHappy3(n int) bool {
	slow, fast := calcSum(n), calcSum(calcSum(n))
	for slow!=fast {
		slow = calcSum(slow)
		fast = calcSum(calcSum(fast))
	}
	return slow==1
}

func calcSum(n int) int {
	sum, bit := 0, 0
	for n > 0 {
		bit = n % 10
		sum += bit * bit
		n = n / 10
	}
	return sum
}