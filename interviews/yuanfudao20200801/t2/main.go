package main

import "fmt"

func main() {
	N := 0	// 学生数
	fmt.Scan(&N)

	students := make([]Student, N)
	for i:=0; i<N; i++ {
		students[i].NO = i + 2
		fmt.Scan(&students[i].HoldBonus, &students[i].BonusFrom)
	}

	fmt.Println(students)

	ans := sol(students)

	fmt.Println(ans)
}

//3
//2 0
//1 2
//-1 2
// 输出
// 3

type Student struct {
	NO int			// 学生编号，从2开始，表示自己是第几行的同学
	HoldBonus int	// 持有的奖励
	BonusFrom int	// 奖励来源自第几行的同学，若为0表示自己是第一个发奖券的
	Max int	// 所获得的最大奖金数
}

func students2graph(students []Student) {}

// 从题面的描述来看，分发是固定的：N张奖券，N个人，每个同学在手头奖券>1时，自己留一张，其余分给其他同学
// 这样来看，最终大奖必定是分发者获得，也就是要计算第一个分发者最多能得多少奖金
// 那么就需要保存每个人分发的列表，或者能判断出某人是不是自己分发过的对象
func sol(students []Student) int {
	// 首先找到初始分发者
	first := 0
	for i := 0; i< len(students); i++ {
		if students[i].BonusFrom == 0 {
			first = students[i].NO
			break
		}
	}

	return dp(students, first)
}

// sender这个人发的奖券，它最多拿多少
func dp(students []Student, sender int) int {

	max := 0
	for i:=0; i<len(students); i++ {
		if students[i].BonusFrom == sender {
			tmp := students[i].Max
			if tmp == 0 {	// 说明没计算过
				tmp = dp(students, students[i].NO)
			}
			if tmp > max {
				max = tmp
			}
		}
	}
	students[sender-2].Max = max + students[sender-2].HoldBonus
	return students[sender-2].Max
}
