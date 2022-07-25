// Package dissovle Slice a geometric polygon.
package dissovle

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestDissovlePolygon(t *testing.T) {
	type args struct {
		poly matrix.PolygonMatrix
		diss matrix.Steric
	}
	tests := []struct {
		name       string
		args       args
		wantResult matrix.Steric
		wantErr    bool
	}{
		{"poly line", args{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}, matrix.LineMatrix{{0.5, 1.5}, {2.5, 1.5}}},
			matrix.Collection{matrix.PolygonMatrix{{{1, 1.5}, {2, 1.5}, {2, 1}, {1, 1}, {1, 1.5}}},
				matrix.PolygonMatrix{{{1, 1.5}, {2, 1.5}, {2, 2}, {1, 2}, {1, 1.5}}},
			}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := DissovlePolygon(tt.args.poly, tt.args.diss)
			if (err != nil) != tt.wantErr {
				t.Errorf("DissovlePolygon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("DissovlePolygon() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
