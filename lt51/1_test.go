package lt51

import (
	"fmt"
	"testing"
)

func TestGenBlankChessBoard2(t *testing.T) {
	n := 2
	b := genBlankChessBoard2(n)
	fmt.Println(b)
}

func TestBoardToRet2(t *testing.T) {
	board := []uint32{1, 2}
	// 01
	// 10
	ret := boardToRet2(board)
	fmt.Println(ret)
}
