package geos

import "testing"

func TestMultiPolygon_Nums(t *testing.T) {
	mp, _ := UnmarshalString(`MULTIPOLYGON (((40 40, 20 45, 45 30, 40 40)),
	((20 35, 10 30, 10 10, 30 5, 45 20, 20 35)),
	((30 20, 20 15, 20 25, 30 20)))`)
	tests := []struct {
		name string
		mp   MultiPolygon
		want int
	}{
		{name: "nums", mp: mp.(MultiPolygon), want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Nums(); got != tt.want {
				t.Errorf("MultiPolygon.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}
