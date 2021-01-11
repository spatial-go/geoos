package geoos

import "testing"

func TestMultiLineString_Nums(t *testing.T) {
	multiLineString := MultiLineString{
		{{10, 130}, {50, 190}, {110, 190}, {140, 150}, {150, 80}, {100, 10}, {20, 40}, {10, 130}},
		{{70, 40}, {100, 50}, {120, 80}, {80, 110}, {50, 90}, {70, 40}},
	}
	tests := []struct {
		name string
		mls  MultiLineString
		want int
	}{
		{name: "geometry multiLineString", mls: multiLineString, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mls.Nums(); got != tt.want {
				t.Errorf("MultiLineString.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}
