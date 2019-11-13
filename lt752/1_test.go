package lt752

import (
	"fmt"
	"testing"
)

// 遇到了将string转为[]byte后修改[]byte疑似影响了原string的情况，这里测试下
func TestString2Bytes(t *testing.T) {
	str := "eiger"
	byt := []byte(str)
	fmt.Println(byt)
	byt[1] = 'z'
	fmt.Println(byt)
	fmt.Println(str)
}
// []byte(string)用的不是同一块内存空间，不用担心