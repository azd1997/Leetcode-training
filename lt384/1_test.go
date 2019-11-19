package lt384

import (
	"fmt"
	"testing"
)

// 测试下slice append的一些细节
func TestSolution_Shuffle1_SliceAppend(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("s1=%v, pointer=%p, len=%d, cap=%d\n", s1, s1, len(s1), cap(s1))
	s2 := make([]int, len(s1))		// 测试make切片的结果
	fmt.Printf("s2=%v, pointer=%p, len=%d, cap=%d\n", s2, s2, len(s2), cap(s2))

	s3 := s1[:4]
	fmt.Printf("s3=%v, pointer=%p, len=%d, cap=%d\n", s3, s3, len(s3), cap(s3))
	s4 := append(s1[:4], s1[5:]...)
	fmt.Printf("s4=%v, pointer=%p, len=%d, cap=%d\n", s4, s4, len(s4), cap(s4))
	s5 := append(s1[:4], s1[5:]...)		// append后长度不足原切片
	fmt.Printf("s5=%v, pointer=%p, len=%d, cap=%d\n", s5, s5, len(s5), cap(s5))

	s6 := append(s1[:4], s1[2:]...)		// append后长度超过原切片
	fmt.Printf("s6=%v, pointer=%p, len=%d, cap=%d\n", s6, s6, len(s6), cap(s6))
}
//=== RUN   TestSolution_Shuffle1_SliceAppend
//s1=[1 2 3 4 5 6 7 8], pointer=0xc00001c200, len=8, cap=8
//s2=[0 0 0 0 0 0 0 0], pointer=0xc00001c240, len=8, cap=8
//s3=[1 2 3 4], pointer=0xc00001c200, len=4, cap=8
//s4=[1 2 3 4 6 7 8], pointer=0xc00001c200, len=7, cap=8
//s5=[1 2 3 4 7 8 8], pointer=0xc00001c200, len=7, cap=8
//s6=[1 2 3 4 3 4 7 8 8 8], pointer=0xc000020300, len=10, cap=16		// 只有超过原切片容量，才重分配内存
//--- PASS: TestSolution_Shuffle1_SliceAppend (0.00s)
//PASS



// 蒙特卡罗方法的例子
// 计算圆周率，在正方形内随机打足够多点，统计最后落在正方形四边内切圆内的点数与总点数的比值，当总数越多，比值越接近圆周率3.14

// 在这里，可以对一个较短的数组使用洗牌算法进行打乱，重复足够多次，统计各种情况出现次数，画直方图比较，或直接计算百分比，最后应接近相等。

func TestSolution_Shuffle1(t *testing.T) {
	//test := []int{1, 2, 3}	// 各种排列组合成数字 123， 132， 213，。。。

	m := make(map[int]int)
	N := 1000000
	var number int
	for i:=0; i<N; i++ {
		test := Constructor([]int{1, 2, 3})		// 注意打乱之前要是一样的
		shuffled := test.Shuffle1()
		number = shuffled[0]*100 + shuffled[1]*10 + shuffled[2]
		m[number]++
	}
	var p float64
	for k, v := range m {
		p = float64(v)/float64(N) * 100
		fmt.Printf("probability of %d is %f\n", k, p)
	}
}

func TestSolution_Shuffle2(t *testing.T) {
	//test := []int{1, 2, 3}	// 各种排列组合成数字 123， 132， 213，。。。

	m := make(map[int]int)
	N := 1000000
	var number int
	for i:=0; i<N; i++ {
		test := Constructor([]int{1, 2, 3})		// 注意打乱之前要是一样的
		shuffled := test.Shuffle1()
		number = shuffled[0]*100 + shuffled[1]*10 + shuffled[2]
		m[number]++
	}
	var p float64
	for k, v := range m {
		p = float64(v)/float64(N) * 100
		fmt.Printf("probability of %d is %f\n", k, p)
	}
}