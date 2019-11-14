package lt69

// x 的平方根


// x非负整数，那么x的平方根也一定>=1.
// 其实就是在 [1, x]  找到 y*y <= x 的最大 y
// 怎么做呢？一方面是如果=目标，就返回没错，另一种是平方根为小数，则需要将区间缩小至右上限逼近
func mySqrt1(x int) int {
	if x == 0 {return 0}
	if x <= 3 {return 1}
	l, r := 1, x
	med := 0
	for l < r-1 {       // 保证最后停住时，区间内至少有2个数
		med = l + (r - l) / 2
		if med * med == x {return med}
		if med * med > x {r = med}
		if med * med < x {l = med}
	}
	return l
}

// x非负整数，那么x的平方根也一定>=1.
// 其实就是在 [1, x]  找到 y*y <= x 的最大 y
// 怎么做呢？一方面是如果=目标，就返回没错，另一种是平方根为小数，则需要将区间缩小至右上限逼近
func mySqrt2(x int) int {
	if x == 0 {return 0}
	if x <= 3 {return 1}
	l, r := 1, x
	var med int
	var sqrt int
	for l <= r {
		//fmt.Println("l=",l," r=", r)
		med = l + (r - l) / 2
		sqrt = x / med      // 这样防止med*med乘法溢出
		if sqrt == med {return med}
		if sqrt > med {l = med + 1}     // 比如 6. med = 4, sqrt = 1, sqrt < med  ，
		if sqrt < med {r = med - 1}
	}
	return r
}


// 牛顿法 f(x) = x^2 - a 实际是求 0 = x^2 - a 的解。
// f(x) ~ f(x0) + (x-x0)f'(x0) => 0 = f(x0) + (x-x0)f'(x0) 求x = x0 - f(x0)/f'(x0) = x0 - (x0^2-a)/2x0 = x0/2 + a/2x0
func mySqrt3(x int) int {
	y := x
	for y * y > x {
		y = (y + x/y) / 2
	}
	return y
}



