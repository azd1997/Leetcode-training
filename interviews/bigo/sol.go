package bigo


func t2(str string) string {
	// 转为字符串  [ hello world ]
	slice := []byte(str)
	n := len(slice)
	// 整体反转 [ dlrow olleh ]
	reverse(slice)
	// 单词反转
	start := 0	// 用于去除首部空格 [dlrow olleh ]
	for slice[start] == ' ' {
		start++
	}
	end := n-1  // 用于去除尾部空格 [dlrow olleh]
	for slice[end] == ' ' {
		end--
	}
	// 现在[start:end]是有效范围
	for i:=start; i<end; i++ {
		// 找到完整单词
		wordEnd := i
		for slice[wordEnd] != ' ' {		// 不为空格就算单词一部分
			wordEnd++
		}
		// 对单词作反转
		if wordEnd > i {	// 单词长度>0
			reverse(slice[i:wordEnd])	// 不含wordEnd
			i = wordEnd+1
		} else {
			i++
		}
	}
}

func reverse(slice []byte) {
	n := len(slice)
	for i:=0; i<n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
}
