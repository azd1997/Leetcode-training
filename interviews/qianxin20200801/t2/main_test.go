package main

import "testing"

func Test_divideBy7(t *testing.T) {
	tests := []struct {
		digit []int
		want bool
	}{
		{[]int{1,1,2}, true},
		{[]int{1,2,1}, false},
		{[]int{4,9}, true},
	}
	for _, tt := range tests {
		if got := divideBy7(tt.digit); got != tt.want {
			t.Errorf("divideBy7() = %v, want %v", got, tt.want)
		}
	}
}

func Test_reletive_7(t *testing.T) {
	tests := []struct {
		digit []int
		want int
	}{
		{[]int{1,1,2}, 2},
	}
	for _, tt := range tests {
		if got := reletive_7(tt.digit); got != tt.want {
			t.Errorf("divideBy7() = %v, want %v", got, tt.want)
		}
	}
}