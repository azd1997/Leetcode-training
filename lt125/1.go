package lt125

// 验证回文串

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 	1. 很容易想到，先遍历一遍，去掉其他符号，将大写字母转化为小写字母。再检查数组是否两端对称
// 2. 也可以用双指针迭代，遇杂字符指针前移，不停比较直至两个指针在中间相遇


// 'A'=65,'a'=97

// 1. 去杂后判断数组是否回文
// O(n)/O(n)
//476/476 cases passed (4 ms)
//Your runtime beats 71.11 % of golang submissions
//Your memory usage beats 74.47 % of golang submissions (2.9 MB)
func isPalindrome(s string) bool {
	// 特殊情况
	if s=="" {return true}

	// 去杂
	l := len(s)
	arr := make([]byte, 0, l)
	for i:=0; i<l; i++ {
		if s[i]>='A' && s[i]<='Z' {
			arr = append(arr, s[i]+32)
		} else if (s[i]>='a' && s[i]<='z') || (s[i] >= '0' && s[i] <= '9') {
			arr = append(arr, s[i])
		}
	}

	// 检查是否回文
	for i:=0; i<len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {return false}
	}
	return true
}

// 2. 双指针(两端指针)
//476/476 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 94.04 % of golang submissions (2.7 MB)
func isPalindrome2(s string) bool {
	// 特殊情况
	if s=="" {return true}

	// 左右指针内移
	length := len(s)
	l,r := 0, length-1
	for l<r {
		// 去除非字母非数字
		if s[l]<'0' || (s[l]>'9' && s[l]<'A') || (s[l]>'Z' && s[l]<'a') || s[l]>'z' {l++;continue}
		if s[r]<'0' || (s[r]>'9' && s[r]<'A') || (s[r]>'Z' && s[r]<'a') || s[r]>'z' {r--;continue}
		// 数字比较
		if s[l]>='0' && s[l]<='9' {if s[l]!=s[r] {return false}}
		// 字母比较
		if s[l]>='A' && s[l]<='Z' {if s[l]!=s[r] && s[l]+32!=s[r] {return false}}
		if s[l]>='a' && s[l]<='z' {if s[l]!=s[r] && s[l]-32!=s[r] {return false}}
		l++;r--
	}
	return true
}

