package main

import "fmt"

//题目描述：
//小团从某不知名论坛上突然得到了一个测试默契度的游戏，想和小美玩一次来检验两人的默契程度。
// 游戏规则十分简单，首先有给出一个长度为n的序列，最大值不超过m。
//
//小团和小美各自选择一个[1,m]之间的整数，设小美选择的是l，小团选择的是r，我们认为两个人是默契的需要满足以下条件:
//
//1. l 小于等于r。
//
//2. 对于序列中的元素x，如果0<x<l,或r<x<m+1,则x按其顺序保留下来，要求保留下来的子序列单调不下降。
//
//小团为了表现出与小美最大的默契，因此事先做了功课，他想知道能够使得两人默契的二元组<l,r>一共有多少种。
//
//我们称一个序列A为单调不下降的，当且仅当对于任意的i>j,满足A_i>=A_j。
//
//
//
//输入描述
//输入第一行包含两个正整数m和n，表示序列元素的最大值和序列的长度。(1<=n,m<=100000)
//
//输入第二行包含n个正整数，表示该序列。
//
//输出描述
//输出仅包含一个整数，表示能使得两人默契的二元组数量。
//
//
//样例输入
//5 5
//4 1 4 1 2
//样例输出
//10


func main() {
	m, n := 0, 0
	fmt.Scan(&m, &n)


	nums := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&nums[i])
	}

	ans := sol(m, n, nums)

	fmt.Println(ans)
}

// 默契二元组数量 O(n3)暴力
func sol(m, n int, nums []int) int {
	ans := 0
	for l:=1; l<=m; l++ {
		for r:=l; r<=m; r++ {
			if subSeqAndOk(nums, l, r, m) {
				ans++
			}
		}
	}
	return ans
}

// O(n)
func subSeqAndOk(nums []int, l, r, m int) bool {
	//ret := make([]int, 0, len(nums))
	//for i:=0; i<len(nums); i++ {
	//	if (nums[i] > 0 && nums[i] < l) || (nums[i] > r && nums[i] < m) {
	//		if len(ret) == 0 {
	//			ret = append(ret, nums[i])
	//		} else {
	//			if ret[len(ret)-1] <= nums[i] {
	//				ret = append(ret, nums[i])
	//			}
	//		}
	//	}
	//}
	//return true

	prev := -1
	for i:=0; i<len(nums); i++ {	// nums[i]是正整数
		if (nums[i] > 0 && nums[i] < l) || (nums[i] > r && nums[i] < m) {
			if prev <= nums[i] {
				prev = nums[i]
			} else {
				return false
			}
		}
	}
	return true
}
