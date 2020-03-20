package lcof40

import "math/rand"

// 最小的k个数

// 思考：
// 1. 先整体排序，再取前k小 O(nlogn)。
// 解法1作了额外的工作，因此可以进一步优化
// 2. 使用基于快排的分治解法，只将前k小的数拍好位置，其他不管。 O(nlogk)，
// 解法2还是作了额外功，题目只要求找前k小，不要求前k小有序
// 3. 快排思路找第k小，并且同时也已经找到了比第k小还小的所有数。 O(n)
// 解法3应该是最优解了
//
// 还有一些其他思路：
// 4. 堆 O(nlogk)
// 5. 对元素上下限进行二分法。O(nlogw)，w为最大值最小值差值

func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if n < k || k == 0 { // 测例有k=0的情况
		return nil
	}
	// 前k小元素
	return quick(arr, 0, n-1, k)
}

func quick(arr []int, l, r, k int) []int {
	// 快排思想
	pivot := rand.Intn(r-l+1) + l
	arr[l], arr[pivot] = arr[pivot], arr[l] // 基准交换至开头
	p := partition(arr, l, r)
	if p == k-1 { // 第k小，意味着排好序后位置是k-1。 并且第k小左边都小于第k小
		return arr[:k]
	} else if p > k-1 {
		return quick(arr, l, p-1, k)
	} else { // p<k
		return quick(arr, p+1, r, k)
	}
}

// 将arr分区，返回分区点下标。（由于是升序处理，分区点下标考虑为Less区间末尾）
func partition(arr []int, l, r int) int {
	p := l                        // 返回的分界点
	for i := l + 1; i <= r; i++ { // 循环体内代码可以简写，现在只是为了让意思清晰而已
		if arr[i] >= arr[l] { // 不动
			continue
		}
		if arr[i] < arr[l] { // 同arr[p+1]交换，并且i后移，p后移
			arr[i], arr[p+1] = arr[p+1], arr[i]
			p++
		}
	}
	// 交换arr[l]和arr[p]
	arr[l], arr[p] = arr[p], arr[l]
	return p
}
