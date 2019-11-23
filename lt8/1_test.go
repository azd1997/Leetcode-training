package lt8

import (
	"fmt"
	"math"
	"testing"
)

// 测试下字符子串是否共用内存
func TestString(t *testing.T) {
	str := "abcdef"
	str1 := str[:3]
	fmt.Printf("%p, %s\n", &str, str)
	fmt.Printf("%p, %s\n", &str1, str1)
	//0xc000086320, abcdef
	//0xc000086330, abc
	// 生成了一个新的字符串，但是从字符串地址的差额来看，两个地址是相邻的，说明字符串本身并不存数据，只是指向了底层的字符数组
}

// 测试下 for循环
func TestForLoop(t *testing.T) {
	var i int
	for i=0; i<10; i++ {
		if i == 7 {break}
	}
	fmt.Println(i)		//7
}

func TestInt32(t *testing.T) {
	fmt.Println(math.MinInt32, math.MaxInt32)
}

func TestRegExpSolution(t *testing.T) {
	test := "9999999999"
	test = "   -42"
	test = "    25"
	fmt.Println(myAtoi3(test))
}
