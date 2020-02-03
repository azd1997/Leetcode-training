package lt1344

import (
	"fmt"
	"sort"
	"testing"
)

func TestSol(t *testing.T) {
	a := []int{2,3,1,5}
	b := make([]int, len(a))
	for i:=0; i<len(a); i++ {b[i] = i}
	fmt.Println(b)
	sort.Slice(b, func(i, j int) bool {
		return a[b[i]]<a[b[j]]
	})
	fmt.Println(b)
	for i:=0; i<len(a); i++ {
		idx := b[i]
		fmt.Print(a[idx])
	}
}

func TestMaxJumps3(t *testing.T) {
	arr := []int{40,98,14,22,45,71,20,19,26,9,29,64,76,66,32,79,14,83,62,39,69,25,92,79,70,34,22,19,41,26,5,82,38}
	d := 6
	maxJumps3(arr, d)
}
