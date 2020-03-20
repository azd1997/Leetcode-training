package lt680

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {
	s1 := "abca"
	s2 := "aba"
	ok1, ok2 := validPalindrome(s1), validPalindrome(s2)
	fmt.Println(ok1, ok2)
	if !(ok1 && ok2) {
		t.Error("wrong")
	}
}
