/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/11/20 7:58 PM
* @Description: The file is for
***********************************************************************/

package main

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		raw_str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{raw_str:"abbbbbbAAAdcdddd"}, "a0cbAAAdc0ad"},
		{"case1", args{raw_str:"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBeeeeeeeeeeFYHHnjHAPQQc"}, "0ZB0tB0geFYHHnjHAPQQc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.raw_str); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}