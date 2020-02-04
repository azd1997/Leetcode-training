package lt1323


// 6和9组成的最大数字


// 直接想法就是将每一位都翻转试试，最后保留最大值
// 记得初值就是原数

// 题目给定num<10000，所以可以不用担心除数会溢出

func maximum69Number1 (num int) int {

	cur := num
	bit := 0
	for i:=1000; i>=1; i/=10 {
		bit = cur / 10
		cur = cur % i
		if bit==6 {num = num + 3*i; break}
	}
	return num
}


// 菜一点的写法是这样的：
func maximum69Number2 (num int) int {

	cur, tmp, maxnum := num, num, num
	bit := 0
	for i:=1000; i>=1; i/=10 {
		if i>num {continue}
		bit = cur / 10
		cur = cur % i
		if bit==6 {bit = 9} else {bit = 6}
		tmp = (num/(i*10))*i*10 + bit*i + cur
		if tmp > maxnum {maxnum = tmp}
	}
	return maxnum
}
