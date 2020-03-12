package lt1071

import "sort"

// 字符串的最长公因子

// 提示：

// 1 <= str1.length <= 1000
// 1 <= str2.length <= 1000
// str1[i] 和 str2[i] 为大写英文字母

// 思考：
// 1. 计算两个字符串长度
// 2. 求短字符串的公因子（从最长公因子到最短公因子）
// 3. 依次用短字符串的公因子去试长字符串
// 试公因子时先使用长度检验（长度需能整除），然后再去字符串匹配

// 不对，是先根据两个字符串长度m，n取长度的所有公因子，从大到小去试
// 那么怎么求两个正整数的所有公因子呢？ 9 = 3 × 3 × 1 ； 6 = 3 × 2 × 1		// 公因子从大到小为 [3,1]
// 求两个数的所有公因数的方法：
// 1. 辗转相除求最大公因数max（最大公因数一定是其他公因数的组合）
// 2. for i in [1, sqrt(max)] 去试其 j， j要是整数，且 i * j = max
// 3. 经过步骤2得到所有的公因数并按降序排列

func gcdOfStrings(str1 string, str2 string) string {
	// 确定长短
	m, n := len(str1), len(str2)

	// 确定长度的公因数
	cfs := calcAllCommonFactors(m, n)

	// 求“公共因子”
	for i := 0; i < len(cfs); i++ {
		if checkSubStr(cfs[i], m, str1) && checkSubStr(cfs[i], n, str2) && str1[:cfs[i]] == str2[:cfs[i]] {
			return str1[:cfs[i]]
		}
	}
	// 没有“公共因子”
	return ""
}

// 检查str[:a]是否能将str整除. n是str的长度
func checkSubStr(a, n int, str string) bool {
	for i := 0; i < a; i++ {
		for j := i + a; j < n; j += a {
			if str[j] != str[j-a] {
				return false
			}
		}
	}
	return true
}

func calcAllCommonFactors(a, b int) []int {
	// max common factor
	maxCF := gcd(a, b)
	res := make([]int, 0)
	for i := 1; i*i <= maxCF; i++ {
		if maxCF%i == 0 {
			res = append(res, i, maxCF/i)
		}
	}

	// 排序
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})

	return res
}

func gcd(a, b int) int {
	tmp := a
	for tmp != 0 { // 例如 a=9, b=6 => a=6, b=3 => a=3, b=0(tmp=0) => a就是最大公因数
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}

///////////////////////////////////

// 官方题解还给出了一种基于数学的解法
// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/solution/zi-fu-chuan-de-zui-da-gong-yin-zi-by-leetcode-solu/
// 如果str1，str2有“公共因子”，必然有 str1+str2 = str2+str1
// 并且“公共因子”长度为两字符串长度的最大公约数

func gcdOfStrings2(str1 string, str2 string) string {
	//
	if str1+str2 != str2+str1 {
		return ""
	}

	m, n := len(str1), len(str2)
	// 确定长度的公因数
	_gcd := gcd(m, n)
	return str1[:_gcd]
}

// 其实我自己上面的解法也做了些无用功：如果两个字符串存在“公共因子”，其长度必然为两字符串长度的最大公约数。
