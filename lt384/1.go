package lt384

import "math/rand"

// 打乱数组

// 思考，我能想到的就是利用go的rand库生成随机数去随机选取数组数据填过来。但是如果没有这个库随机数怎样生成或者说随机数生成算法怎么设计、有哪些
// 我是不清楚的

// 于是，看题解。
// 官方题解给出了暴力和洗牌算法两种解法。而 labuladong 的题解还补充了洗牌算法的变体及验蒙特卡罗方法进行检验

// 1. 暴力解法。
// 复制原数组，随机选择复制数组的元素按从前向后序填充至原数组。填充过程中复制数组需要删除对应元素，避免重复访问

// 2. 洗牌算法。
// 向后遍历数组，随机生成数组下标，与当前数据交换。

type Solution struct {
	shuffled []int	// 打乱后的数组
	original []int 	// 原始数组
}

func Constructor(nums []int) Solution {
	return Solution{
		//shuffled: nums,		// 注意，这里不能直接这么写， nums引用会被复用，两个都指向同一片内存，这样，就会使得没办法重置
		//original: nums,
		shuffled: nums,
		original: append([]int{}, nums...),
	}
}

// 恢复数组
func (this *Solution) Reset() []int {
	return this.original
}

// 打乱数组: 暴力解法 O(n2)/O(n)
// 10/10 cases passed (56 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 7.69 % of golang submissions (11.3 MB)
func (this *Solution) Shuffle1() []int {
	nums1 := append([]int{}, this.shuffled...)
	var index int
	l := len(this.shuffled)
	for i:=0; i<l; i++ {
		index = rand.Intn(l-i)
		this.shuffled[i] = nums1[index]
		nums1 = append(nums1[:index], nums1[index+1:]...)	// 删除该位置元素，避免重新被选择到
	}
	return this.shuffled
}

// 打乱数组: 洗牌算法 O(n)/O(n)
//10/10 cases passed (56 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 7.69 % of golang submissions (10.7 MB)
func (this *Solution) Shuffle2() []int {
	var index int
	l := len(this.shuffled)
	for i:=0; i<l; i++ {																			// 这两行可以适当调整，准则是保证产生的所有可能结果有n!种！
		index = rand.Intn(l-i) + i		// 生成 [i, l-1] 范围随机数									// 如果产生随机数处写成 index=rand.Intn(l)，ZE产生的可能有n^n种，这导致实际n!的排列会出现概率不均的情况，也就不叫随机乱置（洗牌）
		this.shuffled[i], this.shuffled[index] = this.shuffled[index], this.shuffled[i]
	}
	return this.shuffled
}