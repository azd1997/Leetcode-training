package lt1071

import (
	"fmt"
	"testing"
)

func TestCalcAllCF(t *testing.T) {
	fmt.Println(calcAllCommonFactors(9, 6))
}

func TestCheckSubStr(t *testing.T) {
	yes1 := checkSubStr(3, 6, "ABABAB")
	yes2 := checkSubStr(3, 6, "ABCABC")
	fmt.Println(yes1, yes2)
}

func TestGCDOfTwoString(t *testing.T) {

	gcd1 := gcdOfStrings("ABABAB", "ABAB")
	gcd2 := gcdOfStrings("ABCABC", "ABC")
	gcd3 := gcdOfStrings("LEET", "CODE")

	fmt.Println(gcd1, gcd2, gcd3)
}

func TestGCDOfTwoString2(t *testing.T) {

	gcd1 := gcdOfStrings2("ABABAB", "ABAB")
	gcd2 := gcdOfStrings2("ABCABC", "ABC")
	gcd3 := gcdOfStrings2("LEET", "CODE")

	fmt.Println(gcd1, gcd2, gcd3)
}
