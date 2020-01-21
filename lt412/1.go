package lt412

import "strconv"

// Fizz Buzz

//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果 n 是3的倍数，输出“Fizz”；
//
//2. 如果 n 是5的倍数，输出“Buzz”；
//
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
//
//示例：
//
//n = 15,
//
//返回:
//[
//    "1",
//    "2",
//    "Fizz",
//    "4",
//    "Buzz",
//    "Fizz",
//    "7",
//    "8",
//    "Fizz",
//    "Buzz",
//    "11",
//    "Fizz",
//    "13",
//    "14",
//    "FizzBuzz"
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/fizz-buzz
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 完全想不出这题有什么优化...
// 但是n的范围需要检查一下


func fizzBuzz(n int) []string {
	if n<=0 {return []string{}}
	res := make([]string, n)
	for i:=1; i<=n; i++ {
		if i%15==0 {
			res[i-1] = "FizzBuzz"
		} else if i%3==0 {
			res[i-1] = "Fizz"
		} else if i%5==0 {
			res[i-1] = "Buzz"
		} else {
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}

// 官方题解居然给出了三种解法...嗯...开拓一下思路

// 官方的判断分支式解法与我的一样，但是有个小细节，执行效率更高。因为取模运算比布尔运算复杂得多
//执行用时 : 4 ms , 在所有 Go 提交中击败了 97.50% 的用户
//内存消耗 : 3.4 MB , 在所有 Go 提交中击败了 97.56% 的用户
func fizzBuzz2(n int) []string {
	if n<=0 {return []string{}}
	res := make([]string, n)
	for i:=1; i<=n; i++ {
		divBy3, divBy5 := i%3==0, i%5==0
		if divBy3 && divBy5 {
			res[i-1] = "FizzBuzz"
		} else if divBy3 {
			res[i-1] = "Fizz"
		} else if divBy5 {
			res[i-1] = "Buzz"
		} else {
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}


// 字符串拼接——更优雅的写法
// 本题只有Fizz,Buzz，如果有更多的分支，那么将会写许多的if-else，不太优雅
// 这里用字符串拼接来改写
func fizzBuzz3(n int) []string {
	if n<=0 {return []string{}}
	res := make([]string, n)
	for i:=1; i<=n; i++ {
		if i%3==0 {res[i-1] += "Fizz"}
		if i%5==0 {res[i-1] += "Buzz"}
		if res[i-1]=="" {res[i-1] = strconv.Itoa(i)}
	}
	return res
}

// 使用哈希表——更优雅的做法
// 实际上这就是表驱动的编程范式
// 前面的做法，如果老板要求改"Fizz"为"Fiz"，或者是增加新的特殊数字与特殊字符串的映射，
// 那么上面的写法也还是许多if，而且修改都是直接在逻辑代码里修改
// 优秀的代码应该是将逻辑与数据分离
// 在这里就是讲特殊的映射关系分离出去，不干扰逻辑代码
// 通常使用哈希表来实现这种数据分离，也常叫做 表驱动
func fizzBuzz4(n int) []string {
	if n<=0 {return []string{}}

	// 在本题中要注意，”FizzBuzz“是有顺序的，所以需要用java中的LinkedHashMap存储
	// 还有就是，如果特殊映射的数字都是较小的数字，不妨使用数组来做这个映射表
	// 而在本题，由于go中map是可以做了无序处理的，所以并不适合用在这里
	// 这里只做一个表驱动编程的示例
	table := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	res := make([]string, n)
	for i:=1; i<=n; i++ {
		for k, v :=range table {
			if i%k==0 {res[i-1] += v}
		}
		if res[i-1]=="" {res[i-1] = strconv.Itoa(i)}
	}
	return res
}