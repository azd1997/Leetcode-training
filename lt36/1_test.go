package lt36

import (
	"fmt"
	"testing"
)

// 测试下byte类型小数-大数的后果
func TestSol_1_4(t *testing.T) {
	small := '.'
	large := '0'
	res := small - large
	fmt.Printf("%T, %v\n", res, res)
}

//int32, -2

// 这说明会自动使用int32来保存负的结果