package main

import "fmt"

func main() {
	a, b := 0, 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			c := sol(a, b)
			fmt.Println(c)
		}
	}
}

func sol(a, b int) int {
	return a + b
}
