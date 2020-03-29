package lt89

// 格雷编码

// 数学公式 整数n的格雷码为 n ^ n/2
// 那么，对 0～2^n-1这n个数字都生成格雷码就好了
func grayCode(n int) []int {
	res := make([]int, 1<<uint(n)) // 2^n
	for i := 0; i < len(res); i++ {
		res[i] = int(binary2gray(uint(i)))
	}
	return res
}

func binary2gray(n uint) uint {
	return n ^ (n >> 1) // n ^ n/2
}

// 2. n比特的格雷码可以递归的由n-1比特的格雷码生成
func grayCode2(n int) []int {
	res := make([]int, 0, 1<<uint(n)) // 2^n
	res = append(res, 0)              // 第一个格雷码
	highestBit := 0
	for i := 0; i < n; i++ {
		highestBit = 1 << uint(i)
		for j := len(res) - 1; j >= 0; j-- { // j要和i反着遍历
			res = append(res, highestBit|res[j])
		}
	}
	return res
}
