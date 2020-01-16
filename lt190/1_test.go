package lt190

import (
	"fmt"
	"testing"
)

func TestSol1(t *testing.T) {
	var num uint32 = 4
	ans := reverseBits(num)
	fmt.Printf("num = %d, %b; reversed = %d, %b\n", num, num, ans, ans)
}
