// package de9im ...

package de9im

import "testing"

func TestRelateStringsTransposeByRing(t *testing.T) {
	type args struct {
		rs        string
		inputType int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"FF0FFF102", 2}, "FF0FFF1F2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RelateStringsTransposeByRing(tt.args.rs, tt.args.inputType); got != tt.want {
				t.Errorf("RelateStringsTransposeByRing() = %v, want %v", got, tt.want)
			}
		})
	}
}
