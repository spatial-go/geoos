package space

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestTransGeometry(t *testing.T) {
	type args struct {
		inputGeom matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want Geometry
	}{
		{"test", args{matrix.PolygonMatrix{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}}}, Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransGeometry(tt.args.inputGeom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransGeometry() = %v, want %v", got, tt.want)
			}
		})
	}
}
