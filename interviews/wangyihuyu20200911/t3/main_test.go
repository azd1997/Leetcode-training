/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/11/20 8:13 PM
* @Description: The file is for
***********************************************************************/

package main

import (
	"reflect"
	"testing"
)

func Test_split_into_list(t *testing.T) {
	tests := []struct {
		N int
		S string
		want []int
	}{
		{2, "011235", []int{0,1,1,2,3,5}},
		{5,"34561245892364388201628", []int{34,56,12,45,89,236,438,820,1628}},
		{2,"11111", []int{}},
	}
	for _, tt := range tests {
		if got := split_into_list(tt.N, tt.S); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("split_into_list() = %v, want %v", got, tt.want)
		}
	}
}