package filter

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestUniqueArrayFilter_FilterSteric(t *testing.T) {

	type args struct {
		matr matrix.Steric
	}
	tests := []struct {
		name string
		want []matrix.Matrix
		args args
	}{
		{name: "unique array filter",
			args: args{matrix.LineMatrix{{0, 0}, {1, 1}, {2, 1}, {3, 1}, {3, 0}, {0, 0}}},
			want: []matrix.Matrix{{0, 0}, {1, 1}, {2, 1}, {3, 1}, {3, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UniqueArrayFilter{}
			u.FilterSteric(tt.args.matr)
			if got := u.Matrixes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique array filter = %v, want %v", got, tt.want)
			}
		})
	}
}
