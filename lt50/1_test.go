package lt50

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {
	res1 := myPow(2.0, 10)
	res2 := myPow(2.10, 3)
	res3 := myPow(0.00001, 2147483647)

	fmt.Println(res1, res2, res3)
}

func TestSol2(t *testing.T) {
	res1 := myPow2(2.0, 10)
	res2 := myPow2(2.10, 3)
	res3 := myPow2(0.00001, 2147483647)

	fmt.Println(res1, res2, res3)
}
