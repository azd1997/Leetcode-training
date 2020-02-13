package lt166

import (
	"math"
	"strconv"
)

// 分数到小数

//

func fractionToDecimal(numerator int, denominator int) string {

	// 特例： x/0 异常（这里返回最小值）
	if denominator == 0 {return strconv.Itoa(math.MinInt32)}

	// 特例： 0/x = 0
	if numerator == 0 {
		return "0"
	}

	res := ""

	// 如果分子分母符合不相同，需要先加个负号
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		res += "-"
	}

	// 用int64表示被除数和除数
	dividend := int64(numerator)
	divisor := int64(denominator)
	if numerator < 0 {dividend = -dividend}
	if denominator < 0 {divisor = -divisor}

	// 记录商的整数部分
	res += strconv.FormatInt(dividend/divisor, 10)

	// 余数
	remainder := dividend % divisor

	// 余数为0，说明能整除，直接返回
	if remainder == 0 {
		return res
	}

	// 对于不能整除的情况，加一个"."
	res += "."

	// 哈希表用来记录余数是否出现重复(余数重复，就会使得除出来的小数循环)
	// 值为每一次余数对应的字符串中小数位置
	set := make(map[int64]int)
	for remainder != 0 {
		if set[remainder]!=0 {
			i := set[remainder]
			res = res[:i] + "(" + res[i:] + ")"
			// 出现循环了，那么需要在循环的那个位置前插入"("
			break
		}

		set[remainder] = len(res)
		remainder *= 10
		res += strconv.FormatInt(dividend/divisor, 10)
		remainder = remainder % divisor
	}
	return res
}
