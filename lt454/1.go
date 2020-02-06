package lt454

// 四数相加II

// 题目描述已说明四数之和不会溢出，所以不用担心
// 四数组长度都为N


// 这个解法提交后遇到一个较长的数组，里面的元素都较大，超时了
func fourSumCount1(A []int, B []int, C []int, D []int) int {

	// 这种四数之和，无非就是四指针，定三移一，再利用哈希表实现O(1)查询
	// 总体时间复杂度 O(n3)
	// 要注意的是，题目要求的是所有组合数，所以哈希表记录的应该是元素出现次数

	N := len(A)
	if N==0 {return 0}

	// 将D数组记录到哈希表
	mD := make(map[int]int)
	for i:=0; i<N; i++ {
		mD[D[i]]++
	}

	// 遍历
	count := 0
	for i:=0; i<N; i++ {	// A
		for j:=0; j<N; j++ {	// B
			for k:=0; k<N; k++ {
				count += mD[-A[i]-B[j]-C[k]]
			}
		}
	}
	return count
}


// 既然上面超时了
// 那么试试空间换时间？
// 四个数组分成两组A、B， C、D
// O(N2)/O(N2) 将其组合成两个新数组 AB/CD
// 再对两个新数组去找这个合为0的组合
// 这样的话总体时间复杂度就是O(N2)，比上面的快
func fourSumCount2(A []int, B []int, C []int, D []int) int {

	N := len(A)
	if N==0 {return 0}

	// 两两合并
	AB, CD := make([]int, 0, N*N), make([]int, 0, N*N)
	for i:=0; i<N; i++ {
		for j:=0; j<N; j++ {
			AB = append(AB, A[i]+B[j])
			CD = append(CD, C[i]+D[j])
		}
	}

	// 将CD数组记录到哈希表
	mCD := make(map[int]int)
	for i:=0; i<N*N; i++ {
		mCD[CD[i]]++
	}

	// 遍历AB数组
	count := 0
	for i:=0; i<N*N; i++ {
		count += mCD[-AB[i]]
	}
	return count
}


// 上面的解法通过了，但其实还有很大的优化空间
// 我们知道测例肯定是会有许多和重复的组合的，这样的话上面的解法浪费了太多空间
// 其实可以只用哈希表记录组合之和和组合数


// 实际的提交结果比较慢...哈希表占用更多？不管了
func fourSumCount3(A []int, B []int, C []int, D []int) int {

	N := len(A)
	if N==0 {return 0}

	// 将A/B/C/D数组记录到哈希表
	mA, mB := make(map[int]int), make(map[int]int)
	mC, mD := make(map[int]int), make(map[int]int)
	for i:=0; i<N; i++ {
		mA[A[i]]++; mB[B[i]]++; mC[C[i]]++; mD[D[i]]++
	}

	// 两两合并
	mAB, mCD := make(map[int]int), make(map[int]int)
	for k1, v1 := range mA {
		for k2, v2 := range mB {
			mAB[k1+k2] += v1*v2
		}
	}
	for k1, v1 := range mC {
		for k2, v2 := range mD {
			mCD[k1+k2] += v1*v2
		}
	}

	// 遍历mAB
	count := 0
	for k, v := range mAB {
		count += v * mCD[-k]
	}
	return count
}

// 虽然超越100%的解法很快，但是...拒绝奇技淫巧