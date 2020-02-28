package lcci0806

import (
	"fmt"
	"testing"
)

func TestHanota(t *testing.T) {
	A, B, C := []int{3,2,1,0}, []int{}, []int{}
	hanota(&A,&B,&C)
	fmt.Println(C)
}
