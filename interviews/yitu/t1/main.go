package main

import "fmt"

type Op struct {
	OP string
	step int
}

// 每一个测例
type Case struct {
	n int
	m int
	ops []Op
}


func main() {
	T := 0
	fmt.Scan(&T)

	cases := make([]*Case, T)

	for i:=0; i<T; i++ {
		case_ := new(Case)
		fmt.Scan(&case_.n)
		fmt.Scan(&case_.m)
		case_.ops = make([]Op, case_.m)
		for j:=0; j<case_.m; j++ {
			fmt.Scan(&case_.ops[j].OP)
			if case_.ops[j].OP == "G" {
				fmt.Scan(&case_.ops[j].step)
			}
		}
		cases[i] = case_
	}

	// 计算
	for i:=0; i<T; i++ {
		fmt.Printf("Case #%d:\n", i+1)
		sol(cases[i])
	}
}

// 上右下左
// 0123
var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}

// 机器人行走
func sol(c *Case) {
	curx, cury, curD := 0, 0, 0		// 起始点坐标及起始点方向
	n := c.n	// 地图大小
	for _, op := range c.ops {
		switch op.OP {
		case "L":
			curD = (curD  + 4 - 1) % 4	// 左转
		case "R":
			curD = (curD + 1) % 4		// 右转
		case "G":
			curx += dx[curD] * op.step

			if curD == 0 || curD == 2 {	// 向上走或向下走
				cury += dy[curD] * op.step
				if cury >= n  {
					cury = n-1
				}
				if cury < 0 {
					cury = 0
				}

			} else {	// 向右或向左
				curx += dx[curD] * op.step
				if curx >= n {
					curx = n-1
				}
				if curx < 0 {
					curx = 0
				}
			}
		case "P":
			fmt.Printf("%d %d\n", curx, cury)
		}
	}
}

