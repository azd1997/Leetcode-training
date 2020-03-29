package lt137

// 只出现一次的数字II

// 思路：
// 使用哈希表或排序，非常简单。
// 要求O(n)/O(1)，只能是位操作
//
// 0^a=a; a^a=0
// 所以出现个数为奇数个，最后都会是原样:
// 0^a = a; 0^a^a^a = a
// 这里需要想办法区分这两种情况
// 那就是利用两个掩码而不是一个

func singleNumber(nums []int) int {
	seenOnce, seenTwice := 0, 0
	for _, num := range nums {
		// 第一次出现时加到seenOnce而不加到seenTwice
		// 第二次出现时，从seenOnce移除num，并加到seenTwice
		// 第三次出现时，从seenTwice移除num

		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}
	return seenOnce
}

// 这是从题解区看到的评论：
// @jinlinpang 这一题位运算的思路是K为奇数次的解法，实现是K=3的解法。K为偶数次的解法参考题目136。

// 思路是通用的，按照位运算的解题思路，k=5的解法如下：

// public class SingleNumberSolution137 {

//     public int singleNumber(int[] nums) {
//         int seenOnce = 0, seenTwice = 0, seenThird = 0,seenForth = 0;
//         for (int num : nums) {
//             seenOnce = ~seenTwice & ~seenThird & ~seenForth & (seenOnce ^ num);//若seenTwice,seenThird,seenForth不改变，改变seenOnce
//             seenTwice = ~seenOnce & ~seenThird & ~seenForth & (seenTwice ^ num);//若seenOnce,seenThird,seenForth不改变，改变seenTwice
//             seenThird = ~seenOnce & ~seenTwice & ~seenForth & (seenThird ^ num);
//             seenForth = ~ seenOnce & ~seenTwice & ~seenThird & (seenForth ^ num);
//         }
//         return seenOnce;
//     }
// }
// 如此类推，k=2n+1的通用算法也呼之欲出了。

// 一个比较好的题解
// https://leetcode-cn.com/problems/single-number-ii/solution/zi-dong-ji-wei-yun-suan-zui-xiang-xi-de-tui-dao-gu/
//
// 好理解的思路是：
// 对于本题，k=3（除了特定数以外其他都重复三次）0^a^a^a = a
// 就没办法向上一题 k=2那样 利用 a^a=0 了
// 但是可以这么思考：
// 考虑一个数用32bit表示。对于第j位上的1，
// 可知所有的数在该位置的1的总个数 模3 得到的结果就是多出来的那个只出现一次的数在该位的状态
//

func singleNumber2(nums []int) int {
	var i, res uint
	for ; i < 32; i++ {
		var cnt uint
		for _, num := range nums {
			cnt += (uint(num) >> i) & 1
		}
		res |= (cnt % 3) << i
	}
	return int(res)
}

// 这道题这么做，由于go1.13之前不支持int型的移位操作，因此这里都将之转为uint
// 但是题目测例是包含负数的，因此这样解没法通过。
