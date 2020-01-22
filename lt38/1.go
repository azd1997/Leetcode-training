package lt38

import "strconv"

// 外观数列

//是对前一项的描述。前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
//注意：整数序列中的每一项将表示为一个字符串。
//
// 
//
//示例 1:
//
//输入: 1
//输出: "1"
//示例 2:
//
//输入: 4
//输出: "1211"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-and-say
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 1. 不借助前置信息的话，可以直接遍历字符串，数有几个连续的字符，然后拼成新的字符；这个过程再进行递归

// 1. 暴力
// 时间复杂度不太好算，反正不是很高效的样子，但是击败了100%
// 这个题解区 jimmy00745 等的解法一致
func countAndSay(n int) string {
	if n==1 {return "1"}
	s := []byte("1")
	var tmp []byte
	var count int	// 同样数字连续出现计数
	for i:=2; i<=n; i++ {
		count, tmp = 0, make([]byte, 0)  //记得重置count和tmp
		for j:=0; j<len(s)-1; j++ {
			count++
			if s[j]!=s[j+1] {
				tmp = append(tmp, []byte(strconv.Itoa(count))...)
				tmp = append(tmp, s[j])
				count = 0
			}
		}
		// 还剩最后一个字符
		tmp = append(tmp, []byte(strconv.Itoa(count+1))...)
		tmp = append(tmp, s[len(s)-1])
		// 该把tmp交给s
		s = tmp
	}

	return string(s)
}

// 题解区各种花式解法...动态规划、递归、正则...

// 其实这题用动态规划并没有什么优势的感觉...