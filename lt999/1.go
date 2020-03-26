package lt999

// 车的可用捕获量

// 脑残题。先找到车的位置，再尝试四个方向走
// 遇到第一个卒则停住
// 当然遇到己方的象就要停住

var dy = [4]int{1, 0, -1, 0}
var dx = [4]int{0, 1, 0, -1}

func numRookCaptures(board [][]byte) int {
	ry, rx := 0, 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				ry, rx = i, j
				break
			}
		}
	}

	// 四个方向尝试
	res := 0
	for k := 0; k < 4; k++ {
		for step := 1; step < 8; step++ { // 尝试向k方向前进
			newy, newx := ry+step*dy[k], rx+step*dx[k]
			if newy < 0 || newy >= 8 || newx < 0 || newx >= 8 {
				break
			}
			if board[newy][newx] == 'B' {
				break
			}
			if board[newy][newx] == 'p' {
				res++
				break
			}
		}
	}

	return res
}
