/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/11/20 7:12 PM
* @Description: The file is for
***********************************************************************/

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 使用题目中描述的方法，将 raw_str 压缩后返回
 * @param raw_str string字符串 未压缩的原始字符串
 * @return string字符串
 */
func compress( raw_str string ) string {
	// write code here

	n := len(raw_str)
	if n < 4 {return raw_str}

	newStrSlice := make([]byte, 0, n)

	curCh := raw_str[0]
	curChStartIdx := 0
	raw := []byte(raw_str)
	for i:=1; i<n; i++ {
		if raw_str[i] != curCh {
			// 处理前一段相同字符
			process(curCh, curChStartIdx, i, raw, &newStrSlice)

			// 处理完了
			curCh = raw_str[i]
			curChStartIdx = i
		}
	}
	// 最后还得处理剩下的curCh
	process(curCh, curChStartIdx, n, raw, &newStrSlice)

	return string(newStrSlice)
}

func process(curCh byte, curChStartIdx int, i int, raw_str []byte, newStrSlice *[]byte) {
	cnt := i-curChStartIdx	// 连续的相同字符数量
	if cnt < 4 {	// 不压缩
		*newStrSlice = append(*newStrSlice, raw_str[curChStartIdx:i]...)
	} else if cnt <= 55 {	// 普通压缩
		*newStrSlice = append(*newStrSlice, help(curCh, cnt)...)
	} else {	// 分段压缩
		for cnt >= 55 {
			*newStrSlice = append(*newStrSlice, help(curCh, 55)...)
			cnt -= 55
		}
		if cnt > 0 {
			process(curCh, i-cnt, i, raw_str, newStrSlice)
		}
	}
}


// 注意：a-z 97-122 A-Z 65-90

// 普通压缩 cnt [4,55]
func help(ch byte, cnt int) []byte {
	res := make([]byte, 3)
	res[0] = '0'
	res[2] = ch
	res[1] = cnt2ch(cnt)
	return res
}

func cnt2ch(cnt int) byte {
	if cnt <= 29 {
		return 'z' - byte(29-cnt)
	} else {
		return 'A' + byte(cnt-30)
	}
}