package lt1395

// 小队数量

// 直接暴力干

func numTeams(rating []int) int {
	n := len(rating)
	if n < 3 {
		return 0
	}

	cnt := 0
	for i := 0; i < n-2; i++ {
		// 先固定i，看看右边j是什么情况
		for j := i + 1; j < n-1; j++ {
			if rating[j] > rating[i] { // 则要去右边找rating[k] > rating[j]
				for k := j + 1; k < n; k++ {
					if rating[k] > rating[j] {
						cnt++
					}
				}
			} else if rating[j] < rating[i] { // 则要去右边找rating[k] > rating[j]
				for k := j + 1; k < n; k++ {
					if rating[k] < rating[j] {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}

// 简洁一些的写法：
func numTeams2(rating []int) int {
	n := len(rating)
	if n < 3 {
		return 0
	}

	cnt := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if (rating[i] > rating[j] && rating[j] > rating[k]) ||
					(rating[i] < rating[j] && rating[j] < rating[k]) {
					cnt++
				}
			}
		}
	}

	return cnt
}
