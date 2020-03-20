package lt75

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {
	in := []int{2, 0, 2, 1, 1, 0}
	sortColors(in)
	fmt.Println(in)
}
