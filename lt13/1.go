package lt13

import "fmt"

//罗马数字转整数


//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//
//示例 1:
//
//输入: "III"
//输出: 3
//示例 2:
//
//输入: "IV"
//输出: 4
//示例 3:
//
//输入: "IX"
//输出: 9
//示例 4:
//
//输入: "LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.
//示例 5:
//
//输入: "MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
//在真实的面试中遇到过这道题？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/roman-to-integer
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



// 思考：
// 	1. 直观来讲，就是遍历一遍，找到符合条件的罗马字符/字符串，然后转为数字。为了方便记录当前是否是特殊情况，可以用一个布尔变量表示


// 1. 遍历字符串，if-else
//3999/3999 cases passed (12 ms)
//Your runtime beats 54.55 % of golang submissions
//Your memory usage beats 89.78 % of golang submissions (3.1 MB)
func romanToInt(s string) int {
	if s=="" {return 0}
	sum := 0
	for i:=0; i<len(s); i++ {
		switch s[i] {
		case 'I':
			if i<len(s)-1 && (s[i+1]=='V' || s[i+1]=='X') {
				sum -= 1
			} else {sum += 1}
		case 'V':
			sum += 5
		case 'X':
			if i<len(s)-1 && (s[i+1]=='L' || s[i+1]=='C') {
				sum -= 10
			} else {sum += 10}
		case 'L':
			sum += 50
		case 'C':
			if i<len(s)-1 && (s[i+1]=='D' || s[i+1]=='M') {
				sum -= 100
			} else {sum += 100}
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		default:
			// 遇到非罗马数字字符，出错
			return -1
		}
	}

	return sum
}


func romanToInt2(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// 由于这些字符也是单个byte表示的，所以也就不转rune了
	indexEnd := len(s)-1
	index := 0
	var res int
	var now, next int
	goto LOOP

LOOP:
	// 至少后边还有两个字符
	if index < indexEnd {
		now = roman[string(s[index])]
		next = roman[string(s[index+1])]
		// 处理六类特殊情况
		if now < next {
			switch now {
			case 1:
				if next != 5 || next != 10 {
					fmt.Println("输入的字符串不符合罗马数字规则")
					break
				}
			case 10:
				if next != 50 || next != 100 {
					fmt.Println("输入的字符串不符合罗马数字规则")
					break
				}
			case 100:
				if next != 500 || next != 1000 {
					fmt.Println("输入的字符串不符合罗马数字规则")
					break
				}
			default:
				fmt.Println("输入的字符串不符合罗马数字规则")
				break
			}

			res = res - now + next
			// 这里要注意next索引是否已为最后一个
			if index == indexEnd - 1 {
				return res
			}
			// 否则继续循环
			index = index + 2
			goto LOOP
		}

		// 常规情况下解析单个字符
		res += now
		index++
		goto LOOP
	}
	// index = indexEnd情况
	res += roman[string(s[index])]
	return res
}



// 3. 哈希表
// 这样的写法优雅一些，执行效率也高一些
//3999/3999 cases passed (8 ms)
//Your runtime beats 81.41 % of golang submissions
//Your memory usage beats 18.64 % of golang submissions (5.1 MB)
func romanToInt3(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"IV":4,
		"IX":9,
		"XL":40,
		"XC":90,
		"CD":400,
		"CM":900,
	}

	sum := 0
	i := 0	// 把i作用域放大，方便最后判断还有没有剩余
	for i=0; i<len(s)-1; {
		// 先比较特殊情况是否发生
		if v, ok := roman[s[i:i+2]]; ok {
			sum += v; i += 2
		} else if v1, ok1 := roman[s[i:i+1]]; ok1 {	// 普通情况
			sum += v1; i++
		} else {return -1}	// 包含非法字符
	}
	// 最后可能剩字符(i移位到最后一位)也可能不剩(i位于最后一位的后一位)
	if i == len(s)-1 {
		if v, ok := roman[s[i:i+1]]; ok {
			sum += v
		} else {return -1}
	}

	return sum
}