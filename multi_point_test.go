package geoos

import "testing"

func TestMultiPoint_Nums(t *testing.T) {
	multiPoint, _ := UnmarshalString(`MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`)
	tests := []struct {
		name string
		mp   MultiPoint
		want int
	}{
		{name: "ngeometry multiLineString", mp: multiPoint.(MultiPoint), want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Nums(); got != tt.want {
				t.Errorf("MultiPoint.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}
