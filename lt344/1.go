package lt344

// 反转字符串
// 对以字符数组表示的字符串，不使用额外空间，原地完成左右反转

// 1.我直接提交的解法
func reverseString(s []byte)  {
	// 沿中轴左右交换

	l := len(s)
	for i:=0; i<l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}

// 看了下题解，很奇怪都用双指针去做
// 双指针一个指左，一个指右，同时向中间移动，交换两指针指向的元素。
// 当然是可以，但多了一些指针移动的操作，感觉在这题没有意义，直接交换就行了