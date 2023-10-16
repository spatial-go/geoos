package matrix

import (
	"reflect"
	"testing"
)

func TestUniqueArrayFilter_FilterSteric(t *testing.T) {

	type args struct {
		matr Steric
	}
	tests := []struct {
		name string
		want []Matrix
		args args
	}{
		{name: "unique array filter",
			args: args{LineMatrix{{0, 0}, {1, 1}, {2, 1}, {3, 1}, {3, 0}, {0, 0}}},
			want: []Matrix{{0, 0}, {1, 1}, {2, 1}, {3, 1}, {3, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := CreateFilterMatrix()
			u.FilterEntities(TransMatrixes(tt.args.matr))
			if got := u.Entities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique array filter = %v, want %v", got, tt.want)
			}
		})
	}
}
