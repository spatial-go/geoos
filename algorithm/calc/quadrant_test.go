package calc

import "testing"

func TestDecimalFloat10(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"dec10", args{100.0123456789876}, 100.012345679},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalFloat10(tt.args.x); got != tt.want {
				t.Errorf("DecimalFloat10() = %v, want %v", got, tt.want)
			}
		})
	}
}
