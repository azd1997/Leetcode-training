package lt00000

import "fmt"

// 就模拟比赛咯，如果超时再看看能不能找出规律
// 并且输家是追加到末尾的，而不是与末尾数字交换
// 只意味着arr[0]与arr[1]比较后，不管谁胜谁负，后面的所有元素序号都前移一位
// 为了不移动数组元素，记录两个游标，一个是arr[0]的下标p，一个是arr[1]的下标q
// 一开始 p=0, q=1
// 接下来 若q输了，则q=2
// 若p输了，则将p,q元素对调，再将

// 算了直接暴力模拟
func getWinner(arr []int, k int) int {
	continuousWinCnt := 0	// 连胜计数
	for continuousWinCnt < k {
		fmt.Println(arr, continuousWinCnt)

		tmp := 0
		if arr[0] > arr[1] {	// arr[0]为当前庄家
			continuousWinCnt++	// arr[0]连赢计数加1
			tmp = arr[1]
		} else {
			// 计数重置为1，因为arr[1]赢了这一次
			continuousWinCnt = 1
			tmp = arr[0]
			arr[0] = arr[1]
		}
		arr = append(arr[:1], arr[2:]...)
		arr = append(arr, tmp)
	}
	// 此时arr[0]就是赢家
	return arr[0]
}
